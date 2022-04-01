package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Welcome to the luuk180.dev backend!")
	JsonData := map[string]string{
		"query": `
          {
            user(login: "luuk180") {
              repositories(orderBy: {field: PUSHED_AT, direction: ASC}, first: 100) {
                edges {
                  node {
                    isPrivate
                    name
                    url
                    homepageUrl
                    description
                    diskUsage
                  }
                }
              }
            }
          }
          `,
	}
	JsonValue, _ := json.Marshal(JsonData)
	request, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(JsonValue))
	request.Header.Add("AUTHORIZATION", "bearer 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")
	request.Header.Add("USER_AGENT", "github-api")
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		println("Error: ", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	println(string(data))
}
