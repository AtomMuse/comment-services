package commenthandler

import (
	"atommuse/backend/comment-services/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	@Summary		Create a new comment
//	@Description	Create a new comment
//	@Tags			Comments
//	@Security		BearerAuth
//	@ID				CreateComment
//	@Accept			json
//	@Produce		json
//	@Param			requestExhibition	body	model.RequestCreateComment	true	"Comment data to create"
//	@Success		201
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Failure		500
//	@Router			/api-comments/comments [post]
func (h *CommentHandler) CreateComment(c *gin.Context) {
	// Get user information from request context
	userID, _ := c.Get("user_id")
	firstName, _ := c.Get("user_first_name")
	lastName, _ := c.Get("user_last_name")
	profileImage, _ := c.Get("user_image")

	// Convert user ID to primitive.ObjectID
	userIDObj, ok := userID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id is not a valid ObjectID"})
		return
	}

	// Parse request body to get comment data
	var reqBody model.RequestCreateComment
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Populate User field with user information
	reqBody.User = model.User{
		ID:           userIDObj,
		FirstName:    firstName.(string),
		LastName:     lastName.(string),
		ProfileImage: profileImage.(string),
	}

	// Call CreateComment method with userID
	commentID, err := h.Service.CreateComment(c, userIDObj, &reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"commentID": commentID})
}
