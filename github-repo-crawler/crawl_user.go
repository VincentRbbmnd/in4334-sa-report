package main

import (
	"context"
	"encoding/json"
	"fmt"

	"./models"
)

type RawUserData interface {
}

type ImportantUserData struct {
	ID    float64 `json:"id"`
	Login string  `json:"login"`
	Type  string  `json:"type"`
	Test  string  `json:"organizations_url"`
}

func getImportantUserData(rawUserData RawUserData) ImportantUserData {
	byteData, err := json.Marshal(rawUserData)
	if err != nil {
		fmt.Println("Raw data could not be converted to bytes", err)
	}
	var userData ImportantUserData
	err = json.Unmarshal(byteData, &userData)
	if err != nil {
		panic(err)
	}
	return userData
}

func addUserToDB(userData ImportantUserData, rawUserData RawUserData) {
	var ctx context.Context
	user, err := userDB.GetByGithubID(ctx, int64(userData.ID))
	if user != nil || err == nil {
		return
	}
	byteData, err := json.Marshal(rawUserData)
	if err != nil {
		fmt.Println("Raw data could not be converted to bytes", err)
	}
	var dbUser models.User
	dbUser.Raw = byteData
	dbUser.GithubUserID = int64(userData.ID)
	dbUser.Login = userData.Login
	dbUser.Type = userData.Type
	err = userDB.Add(ctx, &dbUser)
	if err != nil {
		fmt.Println("user not added to db", err)
	}
}
