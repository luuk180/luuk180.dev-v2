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
	http.Handle("/", fsCors(fs))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func fsCors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		fs.ServeHTTP(w, r)
	}
}

func handler(w http.ResponseWriter, _ *http.Request) {
	queryResult := string(ApiQuery())
	_, err := fmt.Fprintf(w, queryResult)
	if err != nil {
		println(err)
	}
}
