package commentsvc

import (
	"atommuse/backend/comment-services/pkg/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *CommentService) UpdateComment(ctx context.Context, commentID primitive.ObjectID, userID primitive.ObjectID, update *model.RequestUpdateComment) error {
	// Check if the user is authorized to update the comment
	comment, err := s.Repository.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	if comment.User.ID != userID {
		return errors.New("unauthorized: user is not allowed to update this comment")
	}

	// Proceed with the update
	err = s.Repository.UpdateComment(ctx, commentID, userID, update)
	if err != nil {
		return err
	}

	return nil
}
