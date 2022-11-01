package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var uri = os.Getenv("db_url")

func main() {
	lambda.Start(updateEvent)
}

func updateEvent(_ context.Context, _ events.CloudWatchEvent) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	coll := client.Database("pi-s").Collection("nameip")
	_, err = coll.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	return nil
}
