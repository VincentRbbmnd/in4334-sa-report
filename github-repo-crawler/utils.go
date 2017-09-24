package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"./models"
)

func githubAPICall(url string, method string, payload *interface{}) *http.Response {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "token "+*githubAPIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	rate := resp.Header.Get("x-ratelimit-remaining")
	fmt.Println("X-ratelimit-remaining: ", rate)
	rateInt, _ := strconv.Atoi(rate)
	rateLimit = rateInt
	return resp
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

type Link struct {
	Rel string
	URL string
}

func parseLinkHeader(linkHeader string) []Link {
	var links []Link
	result := strings.Split(linkHeader, ", <")
	if len(result) != 2 {
		fmt.Println("Link header parse error")
	}
	first := result[0]
	firstSplitted := strings.Split(first, ">;")
	if len(firstSplitted) != 2 {
		fmt.Println("Link header parse error")
	}
	firstLinkURL := strings.Replace(firstSplitted[0], "<", "", -1)
	firstRel := firstSplitted[1]

	links = append(links, Link{URL: firstLinkURL, Rel: firstRel})
	second := result[1]
	secondSplitted := strings.Split(second, ">;")
	if len(secondSplitted) != 2 {
		fmt.Println("Link header parse error")
	}
	secondLinkURL :=secondSplitted[0]
	secondRel :=secondSplitted[1]

	links = append(links, Link{URL: secondLinkURL, Rel: secondRel})
	fmt.Println("return: ", links)
	return links
}
