package commentsvc

import (
	"atommuse/backend/comment-services/pkg/model"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *CommentService) GetCommentByID(ctx context.Context, commentID primitive.ObjectID) (*model.Comment, error) {
	return s.Repository.GetCommentByID(ctx, commentID)
}

func (s *CommentService) GetCommentsByExhibitionID(ctx context.Context, exhibitionID primitive.ObjectID) ([]*model.Comment, error) {
	return s.Repository.GetCommentsByExhibitionID(ctx, exhibitionID)
}
