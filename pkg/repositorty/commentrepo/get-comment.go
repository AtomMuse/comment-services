package commentrepo

import (
	"atommuse/backend/comment-services/pkg/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *CommentRepository) GetCommentByID(ctx context.Context, commentID primitive.ObjectID) (*model.Comment, error) {
	var comment model.Comment
	filter := bson.M{"_id": commentID}

	err := r.Collection.FindOne(ctx, filter).Decode(&comment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}

	return &comment, nil
}

func (r *CommentRepository) GetCommentsByExhibitionID(ctx context.Context, exhibitionID primitive.ObjectID) ([]*model.Comment, error) {
	var comments []*model.Comment

	filter := bson.M{"exhibitionID": exhibitionID}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment model.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
