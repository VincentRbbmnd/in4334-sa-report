package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func githubAPICall(url string, method string, payload *interface{}) io.ReadCloser {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "token "+*githubAPIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	rate := resp.Header.Get("x-ratelimit-remaining")
	fmt.Println(rate)
	rateInt, _ := strconv.Atoi(rate)
	rateLimit = rateInt
	return resp.Body
}
