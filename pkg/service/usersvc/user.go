package usersvc

import (
	"atommuse/backend/comment-services/pkg/model"
	"atommuse/backend/comment-services/pkg/repositorty/userrepo"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServices struct {
	Repository userrepo.UserRepository
}
type IUserServices interface {
	GetCommentByID(ctx context.Context, commentID primitive.ObjectID) (*model.Comment, error)
}

func (s *UserServices) GetCommentByID(ctx context.Context, commentID primitive.ObjectID) (*model.Comment, error) {
	return s.Repository.GetCommentByID(ctx, commentID)
}
