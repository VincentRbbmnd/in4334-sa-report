package main

import (
	"context"
	"encoding/json"
	"fmt"

	"./models"
)

// RawRepoData contains list of repos in raw form
type RawRepoData interface{}

type ImportantRepoData struct {
	FullName string             `json:"full_name"`
	Owner    ImportantOwnerData `json:"owner"`
	ID       float64            `json:"id"`
}

type UserDataFromRepoData struct {
	Owner RawUserData `json:"owner"`
}

type ImportantOwnerData struct {
	Name string `json:"login"`
	Type string `json:"type"`
}

func createRepoCrawlingURL(current string) string {
	return "https://api.github.com/repos/" + current
}

func processRepoData(currentProject string) (int64, error) {
	resp := githubAPICall(createRepoCrawlingURL(currentProject), "GET", nil)
	var res RawRepoData
	err := json.NewDecoder(resp.Body).Decode(&res)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Data could not be decoded into struct", err)
	}
	byteData, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Raw data could not be converted to bytes", err)
	}
	var repoData ImportantRepoData
	err = json.Unmarshal(byteData, &repoData)
	if err != nil {
		panic(err)
	}
	addRepoToDB(repoData, byteData)

	var userData UserDataFromRepoData
	err = json.Unmarshal(byteData, &userData)
	if err != nil {
		fmt.Println("Raw repo data could not be decoded further into struct", err)
	}
	//ADD USERS TO DB
	author := getImportantUserData(userData.Owner)
	addUserToDB(author, userData.Owner)

	return int64(repoData.ID), nil
}

func addRepoToDB(repoData ImportantRepoData, byteData []byte) {
	var ctx context.Context

	repo, err := repoDB.Get(ctx, int64(repoData.ID))
	// if commit already added return
	if repo != nil || err == nil {
		return
	}
	var dbRepo models.Repo
	dbRepo.FullName = repoData.FullName
	dbRepo.Raw = byteData
	dbRepo.Owner = repoData.Owner.Name
	dbRepo.UserType = repoData.Owner.Type
	dbRepo.ProjectID = int64(repoData.ID)
	if dbRepo.UserType == "Organization" {
		dbRepo.Org = true
	}
	err = repoDB.Add(ctx, &dbRepo)
	if err != nil {
		fmt.Println("repo not added to db", err)
	}
}
