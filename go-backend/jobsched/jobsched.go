package main

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"runtime"
)

func main() {
	fmt.Println("Welcome to the luuk180.dev job scheduler!")

	_, err := scheduler.Every(10).Seconds().Run(querySync)
	if err != nil {
		return
	}

	runtime.Goexit()
}
