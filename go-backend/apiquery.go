package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func ApiQuery() string {
	type JsonEntry struct {
		name        string
		url         string
		homepageUrl string
		description string
		diskUsage   int
		isPrivate   bool
	}

	type JsonNode struct {
		node JsonEntry
	}

	type Edges struct {
		edges []JsonNode
	}

	type Repositories struct {
		repositories Edges
	}

	type User struct {
		user Repositories
	}

	type totalRequest struct {
		data User
	}

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
	request, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(JsonValue))
	request.Header.Add("AUTHORIZATION", string(os.Getenv("API")))
	request.Header.Add("USER_AGENT", "github-api")
	client := &http.Client{Timeout: time.Second * 10}
	response, _ := client.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err)
		}
	}(response.Body)
	data, _ := ioutil.ReadAll(response.Body)

	//var trimStart = strings.TrimPrefix(string(data), "{\"data\":{\"user\":{\"repositories\":{\"edges\":[")
	//var trimEnd = strings.TrimSuffix(trimStart, "]}}}}")

	var UnmarshData totalRequest

	_ = json.Unmarshal(data, &UnmarshData)

	println(string(data))
	println(UnmarshData.data.user.repositories.edges[1].node.name)

	return string(data)
}
