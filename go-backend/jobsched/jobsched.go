package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type GithubRemote struct {
	ID          string `gorm:"column:id;primary_key" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Url         string `gorm:"column:url" json:"url"`
	HomepageURL string `gorm:"column:homepageurl" json:"homepageUrl"`
	Description string `gorm:"column:description" json:"description"`
	Diskusage   int    `gorm:"column:diskusage" json:"diskUsage"`
}

func main() {
	lambda.Start(updateEvent)
}

func updateEvent(_ context.Context, _ events.CloudWatchEvent) error {
	conn, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	var getDB []GithubRemote
	conn.Order("id desc").Find(&getDB)

	queryResult := queryGit()
	if queryResult != nil {
		pushDb(queryResult, getDB, conn)
	} else {
		log.Fatal("Could not get data from Github API.")
	}

	return nil
}
