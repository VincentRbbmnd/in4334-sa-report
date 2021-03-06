package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"./models"
)

func githubAPICallStarMediaType(url string, method string, payload *interface{}) *http.Response {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "token "+*githubAPIKey)
	req.Header.Add("Accept", "application/vnd.github.v3.star+json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	rate := resp.Header.Get("x-ratelimit-remaining")
	if rate == "0" {
		return handleRateLimitExceeded(resp, url, method, payload)
	}
	fmt.Println("X-ratelimit-remaining: ", rate)
	rateInt, _ := strconv.Atoi(rate)
	rateLimit = rateInt
	return resp
}

func githubAPICall(url string, method string, payload *interface{}) *http.Response {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "token "+*githubAPIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	rate := resp.Header.Get("x-ratelimit-remaining")
	if rate == "0" {
		return handleRateLimitExceeded(resp, url, method, payload)
	}
	fmt.Println("X-ratelimit-remaining: ", rate)
	rateInt, _ := strconv.Atoi(rate)
	rateLimit = rateInt
	return resp
}

func handleRateLimitExceeded(resp *http.Response, url string, method string, payload *interface{}) *http.Response {
	rateLimitReset := resp.Header.Get("x-ratelimit-reset")

	i, err := strconv.ParseInt(rateLimitReset, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	duration := tm.Sub(time.Now().Add(-time.Minute * time.Duration(1)))
	fmt.Println("Sleepy time till rate limit reset. Minutes:", duration.Minutes())
	fmt.Println("Going back to work at: ", tm.String())
	time.Sleep(duration)
	return githubAPICall(url, method, payload)
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

type LinkHeader struct {
	First Link
	Next  Link
	Prev  Link
	Last  Link
}
type Link struct {
	Rel string
	URL string
}

func parseLinkHeader(linkHeader string) LinkHeader {
	var parsedLinkHeader LinkHeader
	if linkHeader == "" {
		return parsedLinkHeader
	}
	results := strings.Split(linkHeader, ", <")
	if len(results) < 2 {
		panic("No valid link length: ")
	}
	for _, res := range results {
		splittedLink := strings.Split(res, ">; ")
		if len(splittedLink) < 2 {
			panic("No valid sublink link first")
		}
		linkURL := strings.Replace(splittedLink[0], "<", "", -1)
		rel := splittedLink[1]
		switch rel {
		case `rel="prev"`:
			parsedLinkHeader.Prev = Link{URL: linkURL, Rel: rel}
			break
		case `rel="first"`:
			parsedLinkHeader.First = Link{URL: linkURL, Rel: rel}
			break
		case `rel="next"`:
			parsedLinkHeader.Next = Link{URL: linkURL, Rel: rel}
			break
		case `rel="last"`:
			parsedLinkHeader.Last = Link{URL: linkURL, Rel: rel}
			break
		}
	}

	return parsedLinkHeader
}
