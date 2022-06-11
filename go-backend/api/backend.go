package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to the luuk180.dev backend!")
	http.HandleFunc("/query", handler)
	http.HandleFunc("/", welcomeHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func welcomeHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the backend!")
	if err != nil {
		println(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	return
}

func handler(w http.ResponseWriter, _ *http.Request) {
	enableCors(&w)
	queryResult := string(ApiQuery())
	_, err := fmt.Fprintf(w, queryResult)
	if err != nil {
		println(err)
	}
}
