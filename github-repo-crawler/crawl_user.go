package main

type RawUserData interface {
}

type User struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}
