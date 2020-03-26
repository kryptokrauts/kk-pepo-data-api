package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../model"
	"../util"
)

// GetPepoVideos returns array of model.PepoVideo
func (s *Service) GetPepoVideos(limit int64) ([]model.PepoVideo, error) {
	var videoUpdateResultList []model.AggregationResultEventPayload
	pipeline := []bson.M{
		bson.M{
			"$sort": bson.M{"created_at": -1},
		},
		bson.M{
			"$group": bson.M{
				"_id": "$data.activity.video.id",
				"status": bson.M{
					"$addToSet": "$data.activity.video.status",
				},
				"result": bson.M{
					"$first": "$$ROOT",
				},
			},
		},
		bson.M{
			"$match": bson.M{
				"$and": []bson.M{
					bson.M{
						"status": "ACTIVE",
					},
					bson.M{
						"status": bson.M{
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
				"result": 1,
			},
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
		videoContributions(s).FindOne(context.Background(), bson.M{"data.activity.video.id": videoUpdateResult.Result.Data.Activity.Video.ID}, opts).Decode(&videoContribution)
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
