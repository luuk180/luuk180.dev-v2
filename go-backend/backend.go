package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to the luuk180.dev backend!")
	//fs := http.FileServer(http.Dir("./files"))
	http.HandleFunc("/query", handler)
	//http.Handle("/file/", http.StripPrefix("/file", fs))
	http.HandleFunc("/cv", handleCV)
	http.HandleFunc("/", welcomeHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleCV(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./files/cv.pdf")
}

func welcomeHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the backend!")
	if err != nil {
		println(err)
	}
}

func handler(w http.ResponseWriter, _ *http.Request) {
	queryResult := string(ApiQuery())
	_, err := fmt.Fprintf(w, queryResult)
	if err != nil {
		println(err)
	}
}
