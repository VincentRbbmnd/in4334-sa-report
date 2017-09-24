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
var userDB *models.UserDB
var rateLimit int
var githubAPIKey *string

func main() {
	rateLimit = 10
	initDatabase(true)
	remaining, err := checkHowFarIWas()
	if err != nil {
		panic(err)
	}
	startRepoCrawling(remaining.RemainingRepoPages)
}

func checkHowFarIWas() (remainingJSON, error) {
	var res remainingJSON
	dat, err := ioutil.ReadFile("./remaining.json")
	if err != nil {
		fmt.Println("remaining json couldn not be found")
		return res, err
	}
	err = json.Unmarshal(dat, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

type remainingJSON struct {
	RemainingRepoPages int `json:"repo_crawling_page"`
}
