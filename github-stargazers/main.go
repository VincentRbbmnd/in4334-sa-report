package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	. "github.com/vincentrbbmnd/in4334-sa-report/github-repo-crawler/models"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

var db *gorm.DB
var starDB *StarDB
var repoDB *RepoDB
var userDB *UserDB

var githubAPIKey *string

type GithubUser struct {
	Login		string `json:"login"`
	ID			int64 `json:"id"`
}

type GithubStar struct {
	StarredAt	time.Time `json:"starred_at"`
	User 		GithubUser `json:"user"`
}

func main() {
	initDatabase(false)
	var ctx context.Context

	// for repository in repositories
	repos, err := repoDB.List(ctx)
	if err != nil {
		panic(err)
	}
	for _, repo := range repos {
		// get full name
		var fullName = repo.FullName
		// get /repos/:owner/:repo/stargazers
		stars := getStars( fullName )
		// new star from response

		for _, star := range stars {
			user, err := userDB.GetByGithubID( ctx, star.User.ID )
			if err != nil {
				panic(err)
			}
			_, err = starDB.Add( ctx, star.StarredAt, user.ID, repo.ID )
			if err != nil {
				panic(err)
			}
		}
	}

}

func getStars( fullRepoName string) []GithubStar {
	resp := githubAPICall("https://api.github.com/repos/" + fullRepoName + "/stargazers", "GET", nil)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var githubStars []GithubStar
	err = json.Unmarshal(body, &githubStars)
	if err != nil {
		panic(err)
	}
	return githubStars
}

func githubAPICall(url string, method string, payload *interface{}) *http.Response {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add( "Accept", "application/vnd.github.v3.star+json")
	req.Header.Add("Authorization", "token "+*githubAPIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting star", err)
	}
	rate := resp.Header.Get("x-ratelimit-remaining")
	if rate == "0" {
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
	fmt.Println("X-ratelimit-remaining: ", rate)
	return resp
}

