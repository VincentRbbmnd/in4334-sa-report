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
	if rate == "0" {
		panic("HIT API LIMIT FUUUUUUUUUUUUUU")
		//TODO read reset time stamp and sleep time : reset time - now
	}
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
