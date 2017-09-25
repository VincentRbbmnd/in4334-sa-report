package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"lab.weave.nl/hoppa/hoppa-backend/models"

	ghmodels "github.com/VincentRbbmnd/in4334-sa-report/github-repo-crawler/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var userDB *models.UserDB

type GoogleAddresses struct {
	Results []GoogleAddress `json:"results"`
	Status  string          `json:"status"`
}

type GoogleAddress struct {
	Geometry         Geometry `json:"geometry"`
	FormattedAddress string   `json:"formatted_address"`
}

type Geometry struct {
	LocationGoogle LocationGoogle `json:"LocationGoogle"`
}

type LocationGoogle struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func main() {
	initDatabase(true)
	address := "Amstelveen+Nederland"
	fmt.Println(getLocationGoogleForAddress(address))
	var location ghmodels.Location
	fmt.Println("je moeder", location)

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
	return googleAddress.Results[0].Geometry.LocationGoogle
}

func googleAPICall(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	return resp
}
