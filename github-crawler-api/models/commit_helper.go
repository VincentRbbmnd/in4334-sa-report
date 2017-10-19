// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "GHCrawler": Model Helpers
//
// Command:
// $ goagen
// --design=github-crawler-api/design
// --out=$(GOPATH)\src\github-crawler-api
// --version=v1.2.0-dirty

package models

import (
	"context"
	"github.com/VincentRbbmnd/in4334-sa-report/github-crawler-api/app"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"time"
)

// MediaType Retrieval Functions

// ListCommit returns an array of view: default.
func (m *CommitDB) ListCommit(ctx context.Context) []*app.Commit {
	defer goa.MeasureSince([]string{"goa", "db", "commit", "listcommit"}, time.Now())

	var native []*Commit
	var objs []*app.Commit
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Commit", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.CommitToCommit())
	}

	return objs
}

// CommitToCommit loads a Commit and builds the default view of media type Commit.
func (m *Commit) CommitToCommit() *app.Commit {
	commit := &app.Commit{}
	commit.ID = m.ID
	commit.Message = &m.Message
	commit.Sha = m.Sha

	return commit
}

// OneCommit loads a Commit and builds the default view of media type Commit.
func (m *CommitDB) OneCommit(ctx context.Context, id int) (*app.Commit, error) {
	defer goa.MeasureSince([]string{"goa", "db", "commit", "onecommit"}, time.Now())

	var native Commit
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Commit", "error", err.Error())
		return nil, err
	}

	view := *native.CommitToCommit()
	return &view, err
}
