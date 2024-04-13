package commentrepo

import (
	"atommuse/backend/comment-services/pkg/model"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CommentRepository) CreateComment(ctx context.Context, userID primitive.ObjectID, comment *model.RequestCreateComment) (*primitive.ObjectID, error) {
	// Set UserID
	comment.UserId = userID

	// Set the Bangkok timezone offset to UTC+7
	const bangkokOffset = 7 * 60 * 60 // 7 hours in seconds

	// Get the current time in UTC
	currentTimeUTC := time.Now().UTC()

	// Add the Bangkok timezone offset to the current time
	currentTimeBangkok := currentTimeUTC.Add(time.Duration(bangkokOffset) * time.Second)

	// Set CreateDateAt in Bangkok timezone
	comment.CreateDateAt = primitive.NewDateTimeFromTime(currentTimeBangkok)
	comment.UpdateDateAt = primitive.NewDateTimeFromTime(currentTimeBangkok)

	// Insert the comment into the database
	result, err := r.Collection.InsertOne(ctx, comment)
	if err != nil {
		return nil, fmt.Errorf("failed to insert comment into the database: %w", err)
	}

	// Ensure InsertedID is not nil
	if result.InsertedID == nil {
		return nil, errors.New("failed to get inserted ID")
	}

	// Extract the generated ObjectID from the result
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert inserted ID to ObjectID")
	}

	return &objectID, nil
}
