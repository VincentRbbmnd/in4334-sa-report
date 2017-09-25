package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"./models"
)

// RawcommitData contains list of commits in raw form
type RawCommitData []interface{}

type ImportantCommitData struct {
	SHA          string      `json:"sha"`
	AuthorRaw    RawUserData `json:"author"`
	CommitterRaw RawUserData `json:"committer"`
	Commit       Commit      `json:"commit"`
}

type Commit struct {
	Message string `json:"message"`
	Author  Author `json:"author"`
}

type Author struct {
	Date string `json:"date"`
}

func startCommitCrawling() {
	url := "https://api.github.com/repos/microsoft/vscode/commits?per_page=100"
	getCommitsOfRepo(31, url)
}

func getCommitsOfRepo(repoID int, apiUrl string) {
	resp := githubAPICall(apiUrl, "GET", nil)

	var res RawCommitData
	err := json.NewDecoder(resp.Body).Decode(&res)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Data could not be decoded into struct", err)
	}
	for _, commit := range res {
		byteData, err := json.Marshal(commit)
		if err != nil {
			fmt.Println("Raw data could not be converted to bytes", err)
		}
		var commitData ImportantCommitData
		err = json.Unmarshal(byteData, &commitData)
		if err != nil {
			fmt.Println("Raw repo data could not be decoded further into struct", err)
		}
		//ADD USERS TO DB
		author := getImportantUserData(commitData.AuthorRaw)
		addUserToDB(author, commitData.AuthorRaw)
		committer := getImportantUserData(commitData.CommitterRaw)
		if author.ID != committer.ID {
			addUserToDB(committer, commitData.CommitterRaw)
		}
		//ADD COMMIT TO DB
		addCommitToDB(commitData, byteData, repoID, int64(author.ID), int64(committer.ID))
	}
	parsedLinkHeader := parseLinkHeader(resp.Header.Get("link"))
	if parsedLinkHeader.Next != (Link{}) {
		getCommitsOfRepo(repoID, parsedLinkHeader.Next.URL)
	} else {

	}
}

func addCommitToDB(commitData ImportantCommitData, byteData []byte, repositoryID int, authorID int64, committerID int64) {
	var ctx context.Context
	commit, err := commitDB.Get(ctx, commitData.SHA)
	// if commit already added return
	if commit != nil || err == nil {
		return
	}

	var dbCommit models.Commit
	dbCommit.Raw = byteData
	dbCommit.SHA = commitData.SHA
	dbCommit.AuthorID = authorID
	dbCommit.CommitterID = committerID
	dbCommit.RepositoryID = repositoryID
	time, err := time.Parse(time.RFC3339, commitData.Commit.Author.Date)
	if err != nil {
		fmt.Println("COULD NOT PARSE DATE COMMIT")
		panic(err)
	}
	dbCommit.CommitDate = time
	dbCommit.Message = commitData.Commit.Message

	err = commitDB.Add(ctx, &dbCommit)
	if err != nil {
		fmt.Println("commit not added to db", err)
	}
}
