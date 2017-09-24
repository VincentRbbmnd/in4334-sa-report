package main

import (
	"./models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var repoDB *models.RepoDB
var userDB *models.UserDB
var remainingDB *models.RemainingDB
var rateLimit int
var githubAPIKey *string

func main() {
	rateLimit = 10
	initDatabase(true)
	ProjectIDLastRepoCrawling, err := checkHowFarIWas("repo_crawling")
	if err != nil {
		panic(err)
	}
	startRepoCrawling(ProjectIDLastRepoCrawling)
}
