package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"./models"
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

func checkHowFarIWas(crawlType string) (int, error) {
	var ctx context.Context
	res, err := remainingDB.GetWhereCrawlType(ctx, crawlType)
	if err != nil {
		createRemainingType(crawlType)
		return 0, nil
	}
	return res.LastID, nil
}

func createRemainingType(crawlType string) {
	var ctx context.Context
	var remaining models.Remaining
	remaining.CrawlType = crawlType
	remainingDB.Add(ctx, &remaining)
}
