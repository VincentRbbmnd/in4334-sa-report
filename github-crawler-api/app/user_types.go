// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "GHCrawler": Application User Types
//
// Command:
// $ goagen
// --design=github-crawler-api/design
// --out=$(GOPATH)\src\github-crawler-api
// --version=v1.2.0-dirty

package app

import (
	"time"
)

// listPayload user type.
type listPayload struct {
	From  *time.Time `form:"from,omitempty" json:"from,omitempty" xml:"from,omitempty"`
	Limit *int       `form:"limit,omitempty" json:"limit,omitempty" xml:"limit,omitempty"`
	Till  *time.Time `form:"till,omitempty" json:"till,omitempty" xml:"till,omitempty"`
}

// Publicize creates ListPayload from listPayload
func (ut *listPayload) Publicize() *ListPayload {
	var pub ListPayload
	if ut.From != nil {
		pub.From = ut.From
	}
	if ut.Limit != nil {
		pub.Limit = ut.Limit
	}
	if ut.Till != nil {
		pub.Till = ut.Till
	}
	return &pub
}

// ListPayload user type.
type ListPayload struct {
	From  *time.Time `form:"from,omitempty" json:"from,omitempty" xml:"from,omitempty"`
	Limit *int       `form:"limit,omitempty" json:"limit,omitempty" xml:"limit,omitempty"`
	Till  *time.Time `form:"till,omitempty" json:"till,omitempty" xml:"till,omitempty"`
}
