package main

import (
	"gorm.io/gorm"
)

func pushDb(apiResult []GithubRemote, dbQuery []GithubRemote, conn *gorm.DB) {
	for i := range dbQuery {
		var statusCode = 0
		for i2 := range apiResult {
			if dbQuery[i].ID == apiResult[i2].ID {
				conn.Updates(dbQuery[i])
				statusCode = 1
			}
			if statusCode == 1 {
				continue
			}
		}
		if statusCode == 0 {
			conn.Delete(dbQuery[i])
		}
	}
	for i := range apiResult {
		var statusCode = 0
		for i2 := range dbQuery {
			if apiResult[i].ID == dbQuery[i2].ID {
				statusCode = 1
				continue
			}
			if statusCode == 1 {
				continue
			}
		}
		if statusCode == 0 {
			conn.Create(apiResult[i])
		}
	}
	return
}
