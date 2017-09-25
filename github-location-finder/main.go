package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	ghmodels "github.com/VincentRbbmnd/in4334-sa-report/github-repo-crawler/models"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

var db *gorm.DB
var userDB *ghmodels.UserDB
var locationDB *ghmodels.LocationDB

var githubAPIKey *string

type GoogleAddresses struct {
	Results []GoogleAddress `json:"results"`
	Status  string          `json:"status"`
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

type User struct {
	ID  int
	Raw GithubUser
}

type GithubUser struct {
	Location string `json:"location"`
}

func main() {
	initDatabase(true)
	var ctx context.Context

	users, err := userDB.ListNoLocations(ctx)
	if err != nil {
		panic(err)
	}
	for i, user := range users {
		fmt.Println("user login: ", user.Login)
		fmt.Println("user github id: ", user.GithubUserID)
		location := getUserLocation(user.Login)
		if location == "" {
			fmt.Println("TODO set checked and no location")
			user.LocationChecked = true
			err = userDB.Update(ctx, user)
			if err != nil {
				panic(err)
			}
		} else {
			googleLocation := getLocationGoogleForAddress(location)
			locationID, err := locationDB.Add(ctx, googleLocation.Lat, googleLocation.Lng, user.GithubUserID)
			if err != nil {
				panic(err)
			}
			user.LocationID = locationID
			user.LocationChecked = true
			err = userDB.Update(ctx, user)
			if err != nil {
				panic(err)
			}
		}

		//TODO remove
		if i == 2 {
			break
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
	fmt.Println("URL: ", url)
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "token "+*githubAPIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	rate := resp.Header.Get("x-ratelimit-remaining")
	if rate == "0" {
		panic("HIT API LIMIT FUUUUUUUUUUUUUU")
	}
	fmt.Println("X-ratelimit-remaining: ", rate)
	return resp
}

func getLocationGoogleForAddress(address string) LocationGoogle {
	url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + address + "&key=AIzaSyAEn3y2FmCPpqYcc9RfonF8Zw3sbX3PZoM"
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
	var LocationGoogle LocationGoogle
	if len(googleAddress.Results) == 0 {
		return LocationGoogle
	}
	fmt.Println("HIii", googleAddress)
	return googleAddress.Results[0].Geometry.LocationGoogle
}

func googleAPICall(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	return resp
}
