package main

import (
	"context"
	"encoding/json"
	"errors"
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
	counter := 0
	for 1 < 2 {
		fmt.Println("repo: ", repos[counter].FullName)
		users := userDB.ListNoLocationsForRepo(ctx, repos[counter].ProjectID)
		if len(users) == 0 {
			counter++
		}
		for _, user := range users {
			location := getUserLocation(user.Login)
			processUserLocation(location, user)
		}
	}
}

func processUserLocation(location string, user *UserContainer) {
	var ctx context.Context

	fmt.Println("Location: ", location)
	if location == "" {
		fmt.Println("No location found: ", location)
		user.LocationChecked = true
		err := userDB.Update(ctx, &user.User)
		if err != nil {
			panic(err)
		}
	} else {
		var ctx context.Context
		locationFromDB, err := locationDB.GetByLocationString(ctx, location)
		if err != nil {
			googleLocation, err := getLocationGoogleForAddress(location)
			if err == nil {
				fmt.Println("google loc: ", googleLocation)
				locationID, err := locationDB.Add(ctx, googleLocation.Lat, googleLocation.Lng, user.User.GithubUserID, location)
				if err != nil {
					panic(err)
				}
				user.User.LocationID = locationID
			}
		} else {
			fmt.Println("Location already found in db: ", locationFromDB.ID)
			user.User.LocationID = locationFromDB.ID
		}

		user.User.LocationChecked = true
		err = userDB.Update(ctx, &user.User)
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

func getLocationGoogleForAddress(address string) (LocationGoogle, error) {
	var locationGoogle LocationGoogle

	url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + strings.Replace(address, " ", "", -1) + "&key=AIzaSyDpy6APeHM3X1JVqdOyuNkZqOS242e8ij8"
	fmt.Println("GOOGLE PLACES API CALL : ", url)
	resp := googleAPICall(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return locationGoogle, err
	}

	defer resp.Body.Close()
	var googleAddress GoogleAddresses
	err = json.Unmarshal(body, &googleAddress)
	if err != nil {
		return locationGoogle, err
	}
	if googleAddress.ErrorMessage != "" {
		panic("PLACES API RATE LIMIT THIS SHOULD NOT HAPPEN" + googleAddress.ErrorMessage)
	}
	if len(googleAddress.Results) == 0 {
		return locationGoogle, errors.New("AAAH")
	}
	fmt.Println("Google address: ", googleAddress)
	return googleAddress.Results[0].Geometry.LocationGoogle, nil
}

func googleAPICall(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	return resp
}
