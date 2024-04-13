package commenthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	@Summary		GetCommentByID
//	@Description	GetCommentByID
//	@Tags			Comments
//	@Security		BearerAuth
//	@ID				GetCommentByID
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
//	@Router			/api/comments/{id} [get]
func (h *CommentHandler) GetCommentByID(c *gin.Context) {
	commentID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid comment ID"})
		return
	}

	comment, err := h.Service.GetCommentByID(c.Request.Context(), commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

//	@Summary		GetCommentsByExhibitionID
//	@Description	GetCommentsByExhibitionID
//	@Tags			Comments
//	@Security		BearerAuth
//	@ID				GetCommentsByExhibitionID
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path	string	true	"Exhibition ID"
//
//	@Success		201
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/api/comments/exhibitions/{id} [get]
func (h *CommentHandler) GetCommentsByExhibitionID(c *gin.Context) {
	exhibitionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid exhibition ID"})
		return
	}

	comments, err := h.Service.GetCommentsByExhibitionID(c.Request.Context(), exhibitionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": "Comments not found"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
