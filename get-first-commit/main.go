package main

import (
	"context"
	"fmt"

	. "github.com/VincentRbbmnd/in4334-sa-report/github-repo-crawler/models"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

var db *gorm.DB
var userDB *UserDB
var locationDB *LocationDB
var repoDB *RepoDB
var commitDB *CommitDB

func main() {
	initDatabase(false)
	var ctx context.Context
	repos, err := repoDB.List(ctx)
	if err != nil {
		panic(err)
	}
	for _, repo := range repos {
		fmt.Println("repo: ", repo.FullName)
		if !repo.FirstCommitDate.IsZero() {
			fmt.Println("Skipping, already has first commit date set!")
			continue
		}
		commit, err := commitDB.GetFirstCommitByRepoID(ctx, repo.ProjectID)
		if err != nil {
			panic(err)
		}
		fmt.Println("first commit sha: ", commit.SHA)
		fmt.Println("first commit date: ", commit.CommitDate)

		repo.FirstCommitDate = commit.CommitDate
	}
}
