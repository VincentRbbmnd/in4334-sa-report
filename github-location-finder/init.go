package main

import (
	"flag"
	"fmt"
	"os"

	ghmodels "github.com/VincentRbbmnd/in4334-sa-report/github-repo-crawler/models"

	"github.com/jinzhu/gorm"
)

func initDatabase(logMode bool) {
	var err error
	db, err = InitDatabase()
	if err != nil {
		panic(err)
	}
	db.LogMode(logMode)
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"postgis\";")
	db.AutoMigrate(&ghmodels.User{}, &ghmodels.Location{})

	userDB = ghmodels.NewUserDB(db)
	locationDB = ghmodels.NewLocationDB(db)

	db.DB().SetMaxOpenConns(50)
}

func InitDatabase() (*gorm.DB, error) {
	host := os.Getenv("PG_HOST")
	databaseName := os.Getenv("PG_DBNAME")
	user := os.Getenv("PG_USERNAME")
	pass := os.Getenv("PG_PASSWORD")
	if host == "" {
		host = "127.0.0.1"
	}
	if databaseName == "" {
		databaseName = "github_crawler"
	}
	if user == "" {
		user = "postgres"
	}
	if pass == "" {
		pass = "postgres"
	}

	// Connect to DB
	dbHost := flag.String("host", host, "Defaults to localhost")
	dbName := flag.String("db", databaseName, "Defaults to hoppa_dev")
	dbUser := flag.String("user", user, "Defaults to postgres")
	dbPass := flag.String("pass", pass, "Defaults to empty")
	githubAPIKey = flag.String("ghkey", "", "Defaults to empty")
	flag.Parse()

	fmt.Println("Configuration:")
	fmt.Println("DB HOST: " + *dbHost)
	fmt.Println("DB Name: " + *dbName)
	fmt.Println("DB User: " + *dbUser)
	fmt.Println("GHKey: ", *githubAPIKey)

	url := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=disable port=%d host=%s", *dbName, *dbUser, *dbPass, 8082, *dbHost)

	return gorm.Open("postgres", url)
}
