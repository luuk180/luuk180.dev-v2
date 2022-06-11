package main

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"runtime"
)

type ApiResult struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	HomepageURL string `json:"homepageUrl"`
	Description string `json:"description"`
	DiskUsage   int    `json:"diskUsage"`
}

func main() {
	fmt.Println("Welcome to the luuk180.dev job scheduler!")

	_, err := scheduler.Every(10).Minutes().Run(querySync)
	if err != nil {
		return
	}

	runtime.Goexit()
}
