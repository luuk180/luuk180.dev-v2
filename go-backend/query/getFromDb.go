package main

import (
	"encoding/json"
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

func getFromDb() string {
	conn, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Error, failed to open database: ", err)
	}

	var gitQuery []GithubRemote

	conn.Order("diskusage desc").Find(&gitQuery)
	returnJson, _ := json.Marshal(gitQuery)

	return string(returnJson)
}
