package commentrepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CommentRepository) DeleteCommentByID(ctx context.Context, commentID primitive.ObjectID) error {
	// Define the filter to match the comment by its ID
	filter := bson.M{"_id": commentID}

	// Perform the deletion operation
	_, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err // Error occurred while deleting the comment
	}

	return nil
}
