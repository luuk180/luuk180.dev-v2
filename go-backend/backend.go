package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to the luuk180.dev backend!")
	fs := http.FileServer(http.Dir("./files"))
	http.HandleFunc("/query", handler)
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	queryResult := string(ApiQuery())
	_, err := fmt.Fprintf(w, queryResult)
	if err != nil {
		println(err)
	}
}
