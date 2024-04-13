package commentrepo

import (
	"atommuse/backend/comment-services/pkg/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CommentRepository) UpdateComment(ctx context.Context, commentID primitive.ObjectID, userID primitive.ObjectID, update *model.RequestUpdateComment) error {

	// Set the Bangkok timezone offset to UTC+7
	const bangkokOffset = 7 * 60 * 60 // 7 hours in seconds

	// Get the current time in UTC
	currentTimeUTC := time.Now().UTC()

	// Add the Bangkok timezone offset to the current time
	currentTimeBangkok := currentTimeUTC.Add(time.Duration(bangkokOffset) * time.Second)

	filter := bson.M{"_id": commentID, "user._id": userID}

	updateData := bson.M{
		"$set": bson.M{
			"commentMessage": update.CommentMessage,
			"updateDateAt":   primitive.NewDateTimeFromTime(currentTimeBangkok),
		},
	}

	_, err := r.Collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return err
	}

	return nil
}
