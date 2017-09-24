package main

import (
	"context"
	"encoding/json"
	"fmt"

	"./models"
)

// RawcommitData contains list of commits in raw form
type RawCommitData []interface{}

type ImportantCommitData struct {
	SHA          string      `json:"sha"`
	Author       User        `json:"author"`
	AuthorRaw    RawUserData `json:"author"`
	Committer    User        `json:"committer"`
	CommitterRaw RawUserData `json:"committer"`
}

func startCommitCrawling(repository string, repoID int) {
	commits := githubAPICall("https://api.github.com/repos/"+repository+"/commits", "GET", nil)
	fmt.Println("Commits: ", commits)
}

func getLastCommitOfRepo(repository string, repoID int) string {
	if false {
		//TODO add if commits stored for this project in repo get last sha
		return "TODO add if commits stored for this project in repo get last sha"
	} else {
		return getFirstCommitOfRepo(repository, repoID)
	}
}

func getFirstCommitOfRepo(repository string, repoID int) string {

	resp := githubAPICall("https://api.github.com/repos/"+repository+"/commits", "GET", nil)
	link := resp.Header.Get("link")
	parsedLinkHeader := parseLinkHeader(link)
	resp = githubAPICall(parsedLinkHeader[1].URL, "GET", nil)

	var res RawCommitData
	err := json.NewDecoder(resp.Body).Decode(&res)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Data could not be decoded into struct", err)
	}
	byteData, err := json.Marshal(res[len(res)-1])
	if err != nil {
		fmt.Println("Raw data could not be converted to bytes", err)
	}
	var commitData ImportantCommitData
	err = json.Unmarshal(byteData, &commitData)
	if err != nil {
		fmt.Println("Raw repo data could not be decoded further into struct", err)
	}
	addCommitToDB(commitData, byteData, repoID)
	//TODO add committer and author to db
	return commitData.SHA
}

func addCommitToDB(commitData ImportantCommitData, byteData []byte, repositoryID int) {
	var ctx context.Context
	commit, err := commitDB.Get(ctx, commitData.SHA)
	//if commit already added return
	if commit != nil || err == nil {
		return
	}

	var dbCommit models.Commit
	dbCommit.Raw = byteData
	dbCommit.SHA = commitData.SHA
	dbCommit.AuthorID = commitData.Author.Id
	dbCommit.CommitterID = commitData.Committer.Id
	dbCommit.RepositoryID = repositoryID

	err = commitDB.Add(ctx, &dbCommit)
	if err != nil {
		fmt.Println("commit not added to db", err)
	}
}
