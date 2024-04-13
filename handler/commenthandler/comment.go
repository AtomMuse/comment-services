package commenthandler

import (
	"atommuse/backend/comment-services/pkg/service/commentsvc"
)

type CommentHandler struct {
	Service commentsvc.CommentService
}

func NewCommentHandler(service commentsvc.CommentService) *CommentHandler {
	return &CommentHandler{Service: service}
}
