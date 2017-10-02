package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	. "github.com/VincentRbbmnd/in4334-sa-report/github-repo-crawler/models"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

var db *gorm.DB
var userDB *UserDB
var locationDB *LocationDB
var repoDB *RepoDB
var commitDB *CommitDB

var githubAPIKey *string

type GoogleAddresses struct {
	Results      []GoogleAddress `json:"results"`
	Status       string          `json:"status"`
	ErrorMessage string          `json:"error_message"`
}

type GoogleAddress struct {
	Geometry         Geometry `json:"geometry"`
	FormattedAddress string   `json:"formatted_address"`
}

type Geometry struct {
	LocationGoogle LocationGoogle `json:"location"`
}

type LocationGoogle struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type GithubUser struct {
	Location string `json:"location"`
}

//TODO query updates:
// UPDATE public.users
//    SET location_checked=false, location_id=null

func main() {
	initDatabase(false)
	var ctx context.Context
	repos, err := repoDB.List(ctx)
	if err != nil {
		panic(err)
	}
	for _, repo := range repos {
		fmt.Println("repo: ", repo)
		users := userDB.ListNoLocationsForRepo(ctx, repo.ProjectID)
		// users, err := userDB.ListNoLocations(ctx)
		// if err != nil {
		// 	panic(err)
		// }
		if len(users) == 0 {
			fmt.Println("No users found sleepy time for half an hour")
			time.Sleep(time.Minute * 30)
		}
		for _, user := range users {
			fmt.Println("user login: ", user.Login)
			location := getUserLocation(user.Login)
			processUserLocation(location, user)
		}
	}
}

func processUserLocation(location string, user *User) {
	var ctx context.Context

	fmt.Println("Location: ", location)
	if location == "" {
		user.LocationChecked = true
		err := userDB.Update(ctx, user)
		if err != nil {
			panic(err)
		}
	} else {
		var ctx context.Context
		locationFromDB, err := locationDB.GetByLocationString(ctx, location)
		if err != nil {
			googleLocation := getLocationGoogleForAddress(location)
			fmt.Println("google loc: ", googleLocation)
			locationID, err := locationDB.Add(ctx, googleLocation.Lat, googleLocation.Lng, user.GithubUserID)
			if err != nil {
				panic(err)
			}
			user.LocationID = locationID
		} else {
			user.LocationID = locationFromDB.ID
		}
		user.LocationChecked = true
		err = userDB.Update(ctx, user)
		if err != nil {
			panic(err)
		}
	}
}

func getUserLocation(login string) string {
	resp := githubAPICall("https://api.github.com/users/"+login, "GET", nil)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var githubUser GithubUser
	err = json.Unmarshal(body, &githubUser)
	if err != nil {
		panic(err)
	}
	return githubUser.Location
}

func githubAPICall(url string, method string, payload *interface{}) *http.Response {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "token "+*githubAPIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	rate := resp.Header.Get("x-ratelimit-remaining")
	if rate == "0" {
		rateLimitReset := resp.Header.Get("x-ratelimit-reset")

		i, err := strconv.ParseInt(rateLimitReset, 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(i, 0)
		duration := tm.Sub(time.Now().Add(-time.Minute * time.Duration(1)))
		fmt.Println("Sleepy time till rate limit reset. Minutes:", duration.Minutes())
		fmt.Println("Going back to work at: ", tm.String())
		time.Sleep(duration)
		return githubAPICall(url, method, payload)
	}
	fmt.Println("X-ratelimit-remaining: ", rate)
	return resp
}

func getLocationGoogleForAddress(address string) LocationGoogle {
	// var LocationGoogle LocationGoogle

	url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + strings.Replace(address, " ", "", -1) + "&key=AIzaSyDpy6APeHM3X1JVqdOyuNkZqOS242e8ij8"
	fmt.Println("API CALL : ", url)
	resp := googleAPICall(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	var googleAddress GoogleAddresses
	err = json.Unmarshal(body, &googleAddress)
	if err != nil {
		panic(err)
	}
	if googleAddress.ErrorMessage != "" {
		panic(googleAddress.ErrorMessage)
	}
	if len(googleAddress.Results) == 0 {
		panic("AAAH")
	}
	fmt.Println("Google address: ", googleAddress)
	return googleAddress.Results[0].Geometry.LocationGoogle
}

func googleAPICall(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	return resp
}
