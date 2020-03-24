package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../model"
	"../util"
)

// GetPepoVideos returns array of model.PepoVideo
func (s *Service) GetPepoVideos(limit int64) ([]model.PepoVideo, error) {
	var videoUpdateList []model.EventPayload
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetLimit(limit)
	// TODO update filter to exclude deleted videos
	cur, err := videoUpdates(s).Find(context.Background(), bson.M{
		"data.activity.video.status": bson.M{
			"$ne": "DELETED",
		}}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &videoUpdateList)
	if err != nil {
		return nil, err
	}
	var pepoVideos = make([]model.PepoVideo, 0)
	for _, videoUpdate := range videoUpdateList {
		var videoContribution *model.EventPayload
		opts := options.FindOne()
		opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
		if err = videoContributions(s).FindOne(context.Background(), bson.M{"data.activity.video.id": videoUpdate.Data.Activity.Video.ID}, opts).Decode(&videoContribution); err != nil {
			fmt.Println(fmt.Sprintf("previous contribution doesn't exist for video with id: %d", videoUpdate.Data.Activity.Video.ID))
		}
		pepoVideos = append(pepoVideos, util.Transform(videoUpdate, videoContribution))
	}
	return pepoVideos, nil
}

func videoUpdates(s *Service) *mongo.Collection {
	return s.MongoClient.Database("pepo").Collection("video-updates")
}

func videoContributions(s *Service) *mongo.Collection {
	return s.MongoClient.Database("pepo").Collection("video-contributions")
}
