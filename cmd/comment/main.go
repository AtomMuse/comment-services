package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "atommuse/backend/comment-services/cmd/comment/doc"
	"atommuse/backend/comment-services/handler/commenthandler"
	"atommuse/backend/comment-services/pkg/model"
	"atommuse/backend/comment-services/pkg/repositorty/commentrepo"
	"atommuse/backend/comment-services/pkg/service/commentsvc"
	"atommuse/backend/comment-services/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.mongodb.org/mongo-driver/mongo"
)

//	@Title						Comment Service API
//	@Version					v0
//	@Description				Comment Service สำหรับขอจัดการเกี่ยวกับ Workshop Manager ทั้งการสร้าง แก้ไข ลบ Workshop Manager
//	@Schemes					http
//	@SecurityDefinitions.apikey	BearerAuth
//	@In							header
//	@Name						Authorization
func main() {
	initializeEnvironment()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set.")
	}
	log.Println("MongoURI:", mongoURI)

	client, err := utils.ConnectToMongoDB(mongoURI)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	router := setupRouter(client)

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	log.Println("Server started on :8080")
	log.Fatal(router.Run(":8080"))
}

// initializeEnvironment initializes environment variables from .env file
func initializeEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

// authMiddleware is middleware to validate the token and check the role
func authMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		secretKey := os.Getenv("secret_key")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		// Parse the token
		claims := &model.JwtCustomClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			// Check the token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return the secret key for validation

			fmt.Println(secretKey)
			fmt.Println(claims)
			fmt.Println(claims.ID)

			return []byte(secretKey), nil
		})
		// Handle token parsing errors
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			fmt.Println("Token parsing error:", err)
			return
		}

		// Check if the token is valid
		if !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			fmt.Println("Invalid token")
			return
		}

		// Set user ID in context
		c.Set("user_id", claims.ID)
		c.Set("user_first_name", claims.FirstName)
		c.Set("user_last_name", claims.LastName)
		c.Set("user_image", claims.ProfileImage)
		c.Set("user_username", claims.UserName)

		// Check if the role admin
		if claims.Role == "admin" {
			c.Next()
		} else if claims.Role != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			fmt.Println("Insufficient permissions")
			return
		}

		// Continue down the chain to handler etc
		c.Next()
	}
}

// setupRouter initializes the Gin router with routes and middleware
func setupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // Replace "*" with allowed origins
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	// Initialize handlers and services
	commentHandler := initCommentHandler(client)

	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Group routes
	api := router.Group("/api-comments")
	{
		//Comments
		api.POST("/comments", authMiddleware("exhibitor"), commentHandler.CreateComment)
		api.PUT("/comments/:id", authMiddleware("exhibitor"), commentHandler.UpdateComment)
		api.GET("/comments/exhibitions/:id", commentHandler.GetCommentsByExhibitionID)
		api.GET("/comments/:id", authMiddleware("exhibitor"), commentHandler.GetCommentByID)
		api.DELETE("/comments/:id", authMiddleware("exhibitor"), commentHandler.DeleteCommentByID)

	}

	return router
}

// initComplainHandler initializes the Complain handler with required dependencies
func initCommentHandler(client *mongo.Client) *commenthandler.CommentHandler {
	dbCollection := client.Database("atommuse-comment").Collection("comments")
	repo := commentrepo.NewCommentRepository(dbCollection)
	service := commentsvc.NewCommentService(*repo)
	return commenthandler.NewCommentHandler(*service)
}
