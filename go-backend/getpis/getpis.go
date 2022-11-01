package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
)

type item struct {
	Hostname string `bson:"hostname"`
	LocalIp  string `bson:"local_ip"`
}

var uri = os.Getenv("db_url")

func main() {
	lambda.Start(handler)
}

func handler(_ events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"},
		Body:       string(getData()),
		StatusCode: http.StatusOK}
	return &resp, nil
}

func getData() []byte {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	coll := client.Database("pi-s").Collection("nameip")
	var returnList []item

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	err = cursor.All(context.TODO(), &returnList)
	if err != nil {
		panic(err)
	}

	returnJson, err := json.Marshal(&returnList)
	if err != nil {
		panic(err)
	}

	return returnJson
}
