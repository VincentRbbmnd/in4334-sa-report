package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"./models"
)

// RawRepoData contains list of repos in raw form
type RawRepoData []interface{}

type ImportantRepoData struct {
	FullName string             `json:"full_name"`
	Owner    ImportantOwnerData `json:"owner"`
}

type ImportantOwnerData struct {
	Name string `json:"login"`
	Type string `json:"type"`
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func createRepoCrawlingURL(current int) string {
	currentString := strconv.Itoa(current)
	baseURL := "https://api.github.com/repositories?page="
	perPage := "&per_page=100"
	return baseURL + currentString + perPage
}

func startRepoCrawling(remainingPages int) {
	for rateLimit > 0 {
		remainingPages++
		processRepoData(remainingPages)
		writeToRemaining(remainingPages)
	}
	duration := time.Duration(1) * time.Hour
	time.Sleep(duration)
	rateLimit = 5000
	startRepoCrawling(remainingPages)
}

func writeToRemaining(remaining int) {
	data, _ := checkHowFarIWas()
	data.RemainingRepoPages = remaining
	f, _ := os.Create("remaining.json")
	defer f.Close()
	byteData, _ := json.Marshal(data)
	f.WriteString(string(byteData))
}

func processRepoData(remainingPages int) error {
	body := githubAPICall(createRepoCrawlingURL(remainingPages), "GET", nil)
	var res RawRepoData
	err := json.NewDecoder(body).Decode(&res)
	defer body.Close()
	if err != nil {
		fmt.Println("Data could not be decoded into struct", err)
	}
	for _, rawData := range res {
		byteData, err := json.Marshal(rawData)
		if err != nil {
			fmt.Println("Raw data could not be converted to bytes", err)
		}
		var repoData ImportantRepoData
		err = json.Unmarshal(byteData, &repoData)
		if err != nil {
			fmt.Println("Raw repo data could not be decoded further into struct", err)
		}
		addRepoToDB(repoData, byteData)
	}
	return nil
}

func addRepoToDB(repoData ImportantRepoData, byteData []byte) {
	var dbRepo models.Repo
	dbRepo.FullName = repoData.FullName
	dbRepo.Raw = byteData
	dbRepo.Owner = repoData.Owner.Name
	dbRepo.UserType = repoData.Owner.Type
	var ctx context.Context
	if dbRepo.UserType == "Organization" {
		dbRepo.Org = true
	}
	err := repoDB.Add(ctx, &dbRepo)
	if err != nil {
		fmt.Println("repo not added to db", err)
	}
}

func githubAPICall(url string, method string, payload *interface{}) io.ReadCloser {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "token "+*githubAPIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting github repos", err)
	}
	rate := resp.Header.Get("x-ratelimit-remaining")
	fmt.Println(rate)
	rateInt, _ := strconv.Atoi(rate)
	rateLimit = rateInt
	return resp.Body
}
