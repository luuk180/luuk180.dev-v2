package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func pushDb(apiResult []GithubRemote) {

	conn, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	var githubRemote GithubRemote

	for _, v := range apiResult {
		githubRemote.ID = v.ID
		githubRemote.Name = v.Name
		githubRemote.Description = v.Description
		githubRemote.Url = v.Url
		githubRemote.HomepageURL = v.HomepageURL
		githubRemote.Diskusage = v.Diskusage
		conn.Updates(&githubRemote)
	}
	return
}
