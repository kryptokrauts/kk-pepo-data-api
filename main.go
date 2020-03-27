package main

import (
	"context"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"./controller"
	"./service"

	_ "./docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Pepo: kryptokrauts community
// @version 1.0
// @description This API can be used to receive videos of the kryptokrauts community on Pepo.

// @contact.name kryptokrauts
// @contact.url https://kryptokrauts.com
// @contact.email kryptokrauts@protonmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.kryptokrauts.com
// @BasePath /pepo/v1
func main() {
	r := gin.New()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+os.Getenv("MONGO_HOST")+":"+os.Getenv("MONGO_PORT")))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	service := service.New(client)
	controller := controller.New(service)

	v1 := r.Group("/pepo/v1")
	{
		videos := v1.Group("/videos")
		{
			videos.GET("", controller.GetPepoVideos)
		}
	}

	url := ginSwagger.URL(getSwaggerBaseURL() + "/pepo/api/doc.json")
	r.GET("/pepo/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Use(cors.Default())
	r.Run(":" + getPort())
}

func getSwaggerBaseURL() string {
	if baseURL := os.Getenv("PEPO_SWAGGER_BASE_URL"); baseURL != "" {
		return baseURL
	}
	return "http://localhost:8080"
}

func getPort() string {
	if port := os.Getenv("PEPO_API_PORT"); port != "" {
		return port
	}
	return "8080"
}
