package commentsvc

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *CommentService) DeleteCommentByID(ctx context.Context, commentID primitive.ObjectID) error {
	// Check if the comment exists
	_, err := s.Repository.GetCommentByID(ctx, commentID)
	if err != nil {
		return err // Comment not found or other error occurred
	}

	// Proceed with the deletion
	err = s.Repository.DeleteCommentByID(ctx, commentID)
	if err != nil {
		return err // Error occurred while deleting the comment
	}

	return nil
}
