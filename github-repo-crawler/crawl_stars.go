package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"./models"
)

// RawstarData contains list of commits in raw form
type RawStarData []interface{}

type ImportantStarData struct {
	StarredAt time.Time   `json:"starred_at"`
	User      interface{} `json:"user"`
}

func startStarCrawling(repoID int64, repoName string) {
	//TODO
	url := "https://api.github.com/repos/" + repoName + "/stargazers?per_page=100"
	getStargazersOfRepo(repoID, url)
}

func getStargazersOfRepo(repoID int64, apiUrl string) {
	resp := githubAPICallStarMediaType(apiUrl, "GET", nil)
	fmt.Println("URL: ", apiUrl)
	var res RawStarData
	err := json.NewDecoder(resp.Body).Decode(&res)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Data could not be decoded into struct", err)
		getStargazersOfRepo(repoID, apiUrl)
	}
	for _, star := range res {
		byteData, err := json.Marshal(star)
		if err != nil {
			fmt.Println("Raw data could not be converted to bytes", err)
			getStargazersOfRepo(repoID, apiUrl)
		}
		var starData ImportantStarData
		err = json.Unmarshal(byteData, &starData)
		if err != nil {
			fmt.Println("Raw repo data could not be decoded further into struct", err)
			getStargazersOfRepo(repoID, apiUrl)
		}
		//ADD USERS TO DB
		user := getImportantUserData(starData.User)
		if user.ID != 0 {
			addUserToDB(user, starData.User)
		}
		//ADD STAR TO DB
		addStarToDB(starData, byteData, repoID, int64(user.ID))
	}
	parsedLinkHeader := parseLinkHeader(resp.Header.Get("link"))
	if parsedLinkHeader.Next != (Link{}) {
		getStargazersOfRepo(repoID, parsedLinkHeader.Next.URL)
	}
}

func addStarToDB(starData ImportantStarData, byteData []byte, repositoryID int64, userID int64) bool {
	var ctx context.Context
	var dbStar models.Star
	dbStar.UserID = userID
	dbStar.RepoID = repositoryID
	dbStar.StarredAt = starData.StarredAt

	err := starDB.Add(ctx, &dbStar)
	if err != nil {
		return true
	}
	return false
}
