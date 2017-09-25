package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"./models"
)

// RawRepoData contains list of repos in raw form
type RawRepoData []interface{}

type ImportantRepoData struct {
	FullName string             `json:"full_name"`
	Owner    ImportantOwnerData `json:"owner"`
	ID       int                `json:"id"`
}

type ImportantOwnerData struct {
	Name string `json:"login"`
	Type string `json:"type"`
}

func createRepoCrawlingURL(current int) string {
	currentString := strconv.Itoa(current)
	baseURL := "https://api.github.com/repositories?since="
	perPage := "&per_page=100"
	return baseURL + currentString + perPage
}

func startRepoCrawling() {
	ProjectIDLastRepoCrawling, err := checkHowFarIWas("repo_crawling")
	if err != nil {
		panic(err)
	}
	crawlRepos(ProjectIDLastRepoCrawling)
}

func crawlRepos(remainingPages int) {
	for rateLimit > 0 {
		remainingPages++
		lastID, err := processRepoData(remainingPages)
		if err != nil {
			panic(err)
		}
		writeToRemaining(lastID)
	}
	duration := time.Duration(1) * time.Hour
	time.Sleep(duration)
	rateLimit = 5000
	crawlRepos(remainingPages)
}

func writeToRemaining(remaining int) {
	var ctx context.Context
	res, _ := remainingDB.GetWhereCrawlType(ctx, "repo_crawling")
	res.LastID = remaining
	err := remainingDB.Update(ctx, res)
	if err != nil {
		fmt.Println("Remaning to DB went wrong", err)
	}
}

func processRepoData(remainingPages int) (int, error) {
	resp := githubAPICall(createRepoCrawlingURL(remainingPages), "GET", nil)
	var res RawRepoData
	err := json.NewDecoder(resp.Body).Decode(&res)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Data could not be decoded into struct", err)
	}
	lastID := 0
	for _, rawData := range res {
		byteData, err := json.Marshal(rawData)
		if err != nil {
			fmt.Println("Raw data could not be converted to bytes", err)
		}
		var repoData ImportantRepoData
		err = json.Unmarshal(byteData, &repoData)
		if err != nil {
			fmt.Println("Raw repo data could not be decoded further into struct", err)
		}
		addRepoToDB(repoData, byteData)
		lastID = repoData.ID
	}
	return lastID, nil
}

func addRepoToDB(repoData ImportantRepoData, byteData []byte) {
	var dbRepo models.Repo
	dbRepo.FullName = repoData.FullName
	dbRepo.Raw = byteData
	dbRepo.Owner = repoData.Owner.Name
	dbRepo.UserType = repoData.Owner.Type
	dbRepo.ProjectID = repoData.ID
	var ctx context.Context
	if dbRepo.UserType == "Organization" {
		dbRepo.Org = true
	}
	err := repoDB.Add(ctx, &dbRepo)
	if err != nil {
		fmt.Println("repo not added to db", err)
	}
}
