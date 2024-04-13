package commentsvc

import (
	"atommuse/backend/comment-services/pkg/model"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *CommentService) CreateComment(ctx context.Context, userID primitive.ObjectID, comment *model.RequestCreateComment) (*primitive.ObjectID, error) {
	return s.Repository.CreateComment(ctx, userID, comment)
}
