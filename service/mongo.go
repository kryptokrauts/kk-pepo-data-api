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
	var videoUpdateResultList []model.AggregationResultEventPayload
	// TODO return unique video id in case of duplicates
	pipeline := []bson.M{
		bson.M{
			"$group": bson.M{
				"_id": "$data.activity.video.id",
				"data": bson.M{
					"$addToSet": bson.M{
						"_id":        "$_id",
						"topic":      "$topic",
						"created_at": "$created_at",
						"webhook_id": "$webhook_id",
						"version":    "$version",
						"data":       "$data",
					},
				},
			},
		},
		bson.M{
			"$sort": bson.M{"created_at": -1},
		},
		bson.M{
			"$match": bson.M{
				"$and": []bson.M{
					bson.M{
						"data.data.activity.video.status": "ACTIVE",
					},
					bson.M{
						"data.data.activity.video.status": bson.M{
							"$not": bson.M{
								"$eq": "DELETED",
							},
						},
					},
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":    0,
				"result": "$data",
			},
		},
		bson.M{
			"$unwind": "$result",
		},
	}
	cur, err := videoUpdates(s).Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &videoUpdateResultList)
	if err != nil {
		return nil, err
	}
	var pepoVideos = make([]model.PepoVideo, 0)
	for _, videoUpdateResult := range videoUpdateResultList {
		var videoContribution *model.EventPayload
		opts := options.FindOne()
		opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
		if err = videoContributions(s).FindOne(context.Background(), bson.M{"data.activity.video.id": videoUpdateResult.Result.Data.Activity.Video.ID}, opts).Decode(&videoContribution); err != nil {
			fmt.Println(fmt.Sprintf("previous contribution doesn't exist for video with id: %d", videoUpdateResult.Result.Data.Activity.Video.ID))
		}
		pepoVideos = append(pepoVideos, util.Transform(videoUpdateResult.Result, videoContribution))
	}
	return pepoVideos, nil
}

func videoUpdates(s *Service) *mongo.Collection {
	return s.MongoClient.Database("pepo").Collection("video-updates")
}

func videoContributions(s *Service) *mongo.Collection {
	return s.MongoClient.Database("pepo").Collection("video-contributions")
}
