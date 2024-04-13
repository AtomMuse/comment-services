package commentrepo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepository struct {
	Collection *mongo.Collection
}

func NewCommentRepository(collection *mongo.Collection) *CommentRepository {
	return &CommentRepository{Collection: collection}
}
