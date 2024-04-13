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
//	@Router			/api/comments [post]
func (h *CommentHandler) CreateComment(c *gin.Context) {
	// Get the value of user_id and check if it exists in the context
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	// Convert the value of user_id to primitive.ObjectID
	userID, ok := userIDInterface.(primitive.ObjectID)
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

	// Call CreateComment method with userID
	commentID, err := h.Service.CreateComment(c, userID, &reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"commentID": commentID})
}
