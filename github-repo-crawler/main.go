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
var starDB *models.StarDB
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
	for _, repo := range repos {
		if repo.ID < 400 {
			continue
		}
		startCommitCrawling(repo.ProjectID, repo.FullName)
		startStarCrawling(repo.ProjectID, repo.FullName)
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
