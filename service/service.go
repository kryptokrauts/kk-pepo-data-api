package service

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Service properties
type Service struct {
	MongoClient *mongo.Client
}

// New Service
func New(mongoClient *mongo.Client) *Service {
	return &Service{mongoClient}
}
