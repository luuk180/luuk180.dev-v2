package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func querySync() {
	type ApiResult struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		HomepageURL string `json:"homepageUrl"`
		Description string `json:"description"`
		DiskUsage   int    `json:"diskUsage"`
	}

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
	request.Header.Add("AUTHORIZATION", os.Getenv("API"))
	request.Header.Add("USER_AGENT", "github-api")
	client := &http.Client{Timeout: time.Second * 10}
	response, _ := client.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)
	data, _ := ioutil.ReadAll(response.Body)

	validJson := json.Valid(data)
	if validJson {
		var QueryResponse = GitApiRes{}
		err := json.Unmarshal(data, &QueryResponse)
		if err != nil {
			println(err)
		}

		var returnJson []ApiResult

		for _, v := range QueryResponse.Data.User.Repositories.Edges {
			if !v.Node.IsPrivate {
				returnJson = append(returnJson, ApiResult{
					Name:        v.Node.Name,
					URL:         v.Node.URL,
					HomepageURL: v.Node.HomepageURL,
					Description: v.Node.Description,
					DiskUsage:   v.Node.DiskUsage,
				})
			}
		}
		var returnValue []byte
		returnValue, _ = json.Marshal(returnJson)
		fmt.Println(string(returnValue))

		return
	} else {
		fmt.Println("JSON is invalid")
		return
	}
}
