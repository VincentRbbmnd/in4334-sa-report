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
	ID    int    `json:"id"`
	Login string `json:"login"`
	Type  string `json:"type"`
	Test  string `json:"organizations_url"`
}

func getImportantUserData(rawUserData RawUserData) ImportantUserData {
	var user ImportantUserData
	byteData, err := json.Marshal(rawUserData)
	if err != nil {
		fmt.Println("Raw data could not be converted to bytes", err)
	}
	var userData ImportantUserData
	// var dat map[string]interface{}
	err = json.Unmarshal(byteData, &userData)
	if err != nil {
		fmt.Println("Raw user data could not be decoded further into struct", err)
	}
	fmt.Println("user", user.ID)
	return userData
}

func addUserToDB(userData ImportantUserData, rawUserData RawUserData) {
	var ctx context.Context
	fmt.Println("userdata id ", userData.ID)
	user, err := userDB.Get(ctx, userData.ID)
	if user != nil || err == nil {
		return
	}
	byteData, err := json.Marshal(rawUserData)
	if err != nil {
		fmt.Println("Raw data could not be converted to bytes", err)
	}
	var dbUser models.User
	dbUser.Raw = byteData
	dbUser.ID = userData.ID
	dbUser.Login = userData.Login
	dbUser.Type = userData.Type
	err = userDB.Add(ctx, &dbUser)
	if err != nil {
		fmt.Println("user not added to db", err)
	}
}
