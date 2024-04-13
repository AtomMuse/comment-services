package userrepo

import (
	"atommuse/backend/comment-services/pkg/model"
	"atommuse/backend/comment-services/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName, collectionName string) *UserRepository {
	return &UserRepository{
		Collection: client.Database(dbName).Collection(collectionName),
	}
}

func (r *UserRepository) GetCommentByID(ctx context.Context, commentID primitive.ObjectID) (*model.Comment, error) {
	mongoURI := os.Getenv("MONGO_URI_USER")
	if mongoURI == "" {
		return nil, errors.New("MONGO_URI environment variable not set")
	}

	client, err := utils.ConnectToMongoDB(mongoURI)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	// Specify the collection names
	userCollection := client.Database("atommuse-user").Collection("users")

	var comment model.Comment
	filter := bson.M{"_id": commentID}

	err = userCollection.FindOne(ctx, filter).Decode(&comment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}

	return &comment, nil
}
