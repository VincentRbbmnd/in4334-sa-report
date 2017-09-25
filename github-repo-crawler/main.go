package main

import (
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

func main() {
	rateLimit = 10
	initDatabase(false)

	startRepoCrawling()

	// startCommitCrawling()
}
