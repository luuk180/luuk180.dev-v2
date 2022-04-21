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
	type GitApiRes struct {
		Data struct {
			User struct {
				Repositories struct {
					Edges []struct {
						Node struct {
							IsPrivate   bool   `json:"isPrivate"`
							Name        string `json:"name"`
							URL         string `json:"url"`
							HomepageURL string `json:"homepageUrl"`
							Description string `json:"description"`
							DiskUsage   int    `json:"diskUsage"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"repositories"`
			} `json:"user"`
		} `json:"data"`
	}

	//type PublicQuery struct {
	//	Name        string
	//	URL         string
	//	HomepageURL string
	//	Description string
	//	DiskUsage   int
	//}

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

	var QueryResponse = GitApiRes{}
	json.Unmarshal(data, &QueryResponse)

	for _, v := range QueryResponse.Data.User.Repositories.Edges {
		if !v.Node.IsPrivate {
			println(v.Node.Name)
		}
	}

	return string(data)
}
