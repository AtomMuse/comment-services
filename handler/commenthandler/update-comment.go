package commenthandler

import (
	"atommuse/backend/comment-services/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	@Summary		Update a new comment
//	@Description	Update a new comment
//	@Tags			Comments
//	@Security		BearerAuth
//	@ID				UpdateComment
//	@Accept			json
//	@Produce		json
//
//	@Param			id						path	string						true	"comment ID"
//	@Param			requestUpdateExhibition	body	model.RequestUpdateComment	true	"Comment data to update"
//	@Success		201
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Failure		500
//	@Router			/api/comments/{id} [put]
func (h *CommentHandler) UpdateComment(c *gin.Context) {
	commentID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid comment ID"})
		return
	}

	// Check if userID is nil
	userID, exists := c.Get("user_id")
	if !exists || userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errorMessage": "User is not authenticated or user ID is missing"})
		return
	}

	// Assert userID to primitive.ObjectID
	userIDObj, ok := userID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "Failed to convert userID to ObjectID"})
		return
	}

	var requestUpdateComment model.RequestUpdateComment
	if err := c.BindJSON(&requestUpdateComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid request body"})
		return
	}

	err = h.Service.UpdateComment(c.Request.Context(), objID, userIDObj, &requestUpdateComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}
