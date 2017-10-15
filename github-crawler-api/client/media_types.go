// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "GHCrawler": Application Media Types
//
// Command:
// $ goagen
// --design=github-crawler-api/design
// --out=$(GOPATH)\src\github-crawler-api
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// Commit data (default view)
//
// Identifier: application/vnd.commit+json; view=default
type Commit struct {
	// Owner of the commit
	Author *Ghuser `json:"author"`
	// ID of the commit in the database
	ID int `form:"id" json:"id" xml:"id"`
	// Message of the commit
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Unique identifier of the commit
	Sha string `form:"sha" json:"sha" xml:"sha"`
	// Time the commit happened
	Timestamp time.Time `form:"timestamp" json:"timestamp" xml:"timestamp"`
}

// Validate validates the Commit media type instance.
func (mt *Commit) Validate() (err error) {

	if mt.Sha == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "sha"))
	}
	if mt.Author == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "authorOfCommit"))
	}

	if mt.Author != nil {
		if err2 := mt.Author.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeCommit decodes the Commit instance encoded in resp body.
func (c *Client) DecodeCommit(resp *http.Response) (*Commit, error) {
	var decoded Commit
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// CommitCollection is the media type for an array of Commit (default view)
//
// Identifier: application/vnd.commit+json; type=collection; view=default
type CommitCollection []*Commit

// Validate validates the CommitCollection media type instance.
func (mt CommitCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeCommitCollection decodes the CommitCollection instance encoded in resp body.
func (c *Client) DecodeCommitCollection(resp *http.Response) (CommitCollection, error) {
	var decoded CommitCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// GH user data (default view)
//
// Identifier: application/vnd.ghuser+json; view=default
type Ghuser struct {
	// ID of the user in the database
	ID int `form:"id" json:"id" xml:"id"`
	// Location of the user
	Location *Location `json:"location"`
	// Unique username of the user
	Login string `form:"login" json:"login" xml:"login"`
	// Type of the user
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the Ghuser media type instance.
func (mt *Ghuser) Validate() (err error) {

	if mt.Login == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "login"))
	}
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	return
}

// DecodeGhuser decodes the Ghuser instance encoded in resp body.
func (c *Client) DecodeGhuser(resp *http.Response) (*Ghuser, error) {
	var decoded Ghuser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Location as geocoordinates (default view)
//
// Identifier: application/vnd.location+json; view=default
type Location struct {
	// ID of the location in the database
	ID int `form:"id" json:"id" xml:"id"`
	// coordinates lat
	Lat float64 `form:"lat" json:"lat" xml:"lat"`
	// coordinates lng
	Lng float64 `form:"lng" json:"lng" xml:"lng"`
	// Location as specified by user
	LocationString *string `form:"location_string,omitempty" json:"location_string,omitempty" xml:"location_string,omitempty"`
}

// Validate validates the Location media type instance.
func (mt *Location) Validate() (err error) {

	return
}

// DecodeLocation decodes the Location instance encoded in resp body.
func (c *Client) DecodeLocation(resp *http.Response) (*Location, error) {
	var decoded Location
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Commit data (default view)
//
// Identifier: application/vnd.repository+json; view=default
type Repository struct {
	// First commit of the repository
	FirstCommit *Commit `form:"first_commit,omitempty" json:"first_commit,omitempty" xml:"first_commit,omitempty"`
	// Full name of the repo
	FullName string `form:"full_name" json:"full_name" xml:"full_name"`
	// ID of the commit in the database
	ID int `form:"id" json:"id" xml:"id"`
	// If owner is an organization
	Org *bool `form:"org,omitempty" json:"org,omitempty" xml:"org,omitempty"`
	// Name of the owner of the repository
	Owner string `form:"owner" json:"owner" xml:"owner"`
	// Time the commit happened
	ProjectID float64 `form:"project_id" json:"project_id" xml:"project_id"`
}

// Validate validates the Repository media type instance.
func (mt *Repository) Validate() (err error) {

	if mt.Owner == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "owner"))
	}
	if mt.FullName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "full_name"))
	}

	if mt.FirstCommit != nil {
		if err2 := mt.FirstCommit.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeRepository decodes the Repository instance encoded in resp body.
func (c *Client) DecodeRepository(resp *http.Response) (*Repository, error) {
	var decoded Repository
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// RepositoryCollection is the media type for an array of Repository (default view)
//
// Identifier: application/vnd.repository+json; type=collection; view=default
type RepositoryCollection []*Repository

// Validate validates the RepositoryCollection media type instance.
func (mt RepositoryCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeRepositoryCollection decodes the RepositoryCollection instance encoded in resp body.
func (c *Client) DecodeRepositoryCollection(resp *http.Response) (RepositoryCollection, error) {
	var decoded RepositoryCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
