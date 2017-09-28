package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"./models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var repoDB *models.RepoDB
var commitDB *models.CommitDB
var userDB *models.UserDB
var remainingDB *models.RemainingDB
var rateLimit int
var githubAPIKey *string

type Project struct {
	Name string
}

func main() {
	rateLimit = 10
	initDatabase(false)
	var ctx context.Context
	repos, err := repoDB.List(ctx)
	if err != nil {
		panic(err)
	}
	// repoList := getRepoList()
	for _, repo := range repos {
		//Not necessary anymore, since we have top 1000 projects in db.
		// repoID, err := processRepoData(repo.FullName)
		// if err != nil {
		// 	panic(err)
		// }
		startCommitCrawling(repo.ProjectID, repo.FullName)
	}
}

func getRepoList() []string {
	dat, err := ioutil.ReadFile("top1000.json")
	if err != nil {
		panic(err)
	}
	var repoList []string
	err = json.Unmarshal(dat, &repoList)
	return repoList
}
