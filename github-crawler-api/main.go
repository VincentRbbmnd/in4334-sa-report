//go:generate goagen bootstrap -d github-crawler-api/design

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/VincentRbbmnd/in4334-sa-report/github-crawler-api/app"
	"github.com/VincentRbbmnd/in4334-sa-report/github-crawler-api/models"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var commitDB *models.CommitDB
var githubAPIKey *string

func main() {
	// initDB
	initDatabase(true)
	// Create service
	service := goa.New("GHCrawler")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "commits" controller
	c := NewCommitsController(service)
	app.MountCommitsController(service, c)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}

func initDatabase(logMode bool) {
	var err error
	db, err = InitDatabase()
	if err != nil {
		panic(err)
	}
	db.LogMode(logMode)

	commitDB = models.NewCommitDB(db)

	db.DB().SetMaxOpenConns(50)
}

func InitDatabase() (*gorm.DB, error) {
	host := os.Getenv("PG_HOST")
	databaseName := os.Getenv("PG_DBNAME")
	user := os.Getenv("PG_USERNAME")
	pass := os.Getenv("PG_PASSWORD")
	if host == "" {
		host = "86.87.235.82"
	}
	if databaseName == "" {
		databaseName = "github_crawler"
	}
	if user == "" {
		user = "rick"
	}
	if pass == "" {
		pass = "postgres"
	}

	// Connect to DB
	dbHost := flag.String("host", host, "Defaults to localhost")
	dbName := flag.String("db", databaseName, "Defaults to hoppa_dev")
	dbUser := flag.String("user", user, "Defaults to postgres")
	dbPass := flag.String("pass", pass, "Defaults to empty")
	dbPort := flag.String("port", 5432, "defaults to 5432")
	githubAPIKey = flag.String("ghkey", "", "Defaults to empty")
	flag.Parse()

	fmt.Println("Configuration:")
	fmt.Println("DB HOST: " + *dbHost)
	fmt.Println("DB Name: " + *dbName)
	fmt.Println("DB User: " + *dbUser)
	fmt.Println("DB port: ", *dbPort)

	url := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=disable port=%d host=%s", *dbName, *dbUser, *dbPass, *dbPort, *dbHost)

	return gorm.Open("postgres", url)
}
