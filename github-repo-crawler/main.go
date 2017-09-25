package main

import (
	"encoding/json"
	"fmt"
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

	repoList := getRepoList()
	for _, repo := range repoList {
		fmt.Println(repo)
		repoID, err := processRepoData(repo)
		if err != nil {
			panic(err)
		}
		startCommitCrawling(repoID, repo)
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
