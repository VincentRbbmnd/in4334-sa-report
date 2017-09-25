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

func main() {
	initDatabase(true)
	userID := int64(320)
	// userDB.ListNo
	address := "Amstelveen+Nederland"
	fmt.Println(getLocationGoogleForAddress(address))
	googleLocation := getLocationGoogleForAddress(address)
	var ctx context.Context
	fmt.Println("LOCATION: ", googleLocation)
	locationDB.Add(ctx, googleLocation.Lat, googleLocation.Lng, userID)

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
