package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func pushDb(apiResult []ApiResult) {
	type GithubRemote struct {
		ID          string `gorm:"column:id;primary_key"`
		Name        string `gorm:"column:name"`
		Url         string `gorm:"column:url"`
		HomepageURL string `gorm:"column:homepageurl"`
		Description string `gorm:"column:description"`
		Diskusage   int    `gorm:"column:diskusage"`
	}

	conn, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	var githubRemote GithubRemote
	conn.First(&githubRemote)

	for _, v := range apiResult {
		githubRemote.ID = v.Id
		githubRemote.Name = v.Name
		githubRemote.Description = v.Description
		githubRemote.Url = v.URL
		githubRemote.HomepageURL = v.HomepageURL
		githubRemote.Diskusage = v.DiskUsage
		conn.Updates(&githubRemote)
	}
	return
}
