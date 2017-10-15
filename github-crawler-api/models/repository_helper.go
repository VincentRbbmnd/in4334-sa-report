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
	"github-crawler-api/app"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"time"
)

// MediaType Retrieval Functions

// ListRepository returns an array of view: default.
func (m *RepositoryDB) ListRepository(ctx context.Context) []*app.Repository {
	defer goa.MeasureSince([]string{"goa", "db", "repository", "listrepository"}, time.Now())

	var native []*Repository
	var objs []*app.Repository
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Repository", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.RepositoryToRepository())
	}

	return objs
}

// RepositoryToRepository loads a Repository and builds the default view of media type Repository.
func (m *Repository) RepositoryToRepository() *app.Repository {
	repository := &app.Repository{}
	repository.FullName = m.FullName
	repository.ID = m.ID
	repository.Org = &m.Org
	repository.Owner = m.Owner
	repository.ProjectID = m.ProjectID

	return repository
}

// OneRepository loads a Repository and builds the default view of media type Repository.
func (m *RepositoryDB) OneRepository(ctx context.Context, id int) (*app.Repository, error) {
	defer goa.MeasureSince([]string{"goa", "db", "repository", "onerepository"}, time.Now())

	var native Repository
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Repository", "error", err.Error())
		return nil, err
	}

	view := *native.RepositoryToRepository()
	return &view, err
}