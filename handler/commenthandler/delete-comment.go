package commenthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	@Summary		DeleteCommentByID
//	@Description	DeleteCommentByID
//	@Tags			Comments
//	@Security		BearerAuth
//	@ID				DeleteCommentByID
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path	string	true	"Comment ID"
//
//	@Success		201
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Failure		500
//	@Router			/api/comments/{id} [delete]
func (h *CommentHandler) DeleteCommentByID(c *gin.Context) {
	// Extract comment ID from the URL parameter
	commentID := c.Param("id")

	// Convert the comment ID string to ObjectID
	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid comment ID"})
		return
	}

	// Call the service method to delete the comment
	err = h.Service.DeleteCommentByID(c.Request.Context(), objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "Failed to delete comment"})
		return
	}

	// Return success response if deletion is successful
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
