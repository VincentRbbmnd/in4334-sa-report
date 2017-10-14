// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "GHCrawler": Models
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

// Github commit model in DB
type Commit struct {
	ID           int `gorm:"primary_key"` // primary key
	AuthorID     float64
	CommitterID  float64
	CreatedAt    time.Time
	DeletedAt    *time.Time
	Message      string
	Raw          string `sql:"type:jsonb"`
	RepositoryID float64
	Sha          string
	UpdatedAt    time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Commit) TableName() string {
	return "commits"

}

// CommitDB is the implementation of the storage interface for
// Commit.
type CommitDB struct {
	Db *gorm.DB
}

// NewCommitDB creates a new storage type.
func NewCommitDB(db *gorm.DB) *CommitDB {
	return &CommitDB{Db: db}
}

// DB returns the underlying database.
func (m *CommitDB) DB() interface{} {
	return m.Db
}

// CommitStorage represents the storage interface.
type CommitStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Commit, error)
	Get(ctx context.Context, id int) (*Commit, error)
	Add(ctx context.Context, commit *Commit) error
	Update(ctx context.Context, commit *Commit) error
	Delete(ctx context.Context, id int) error

	ListCommit(ctx context.Context) []*app.Commit
	OneCommit(ctx context.Context, id int) (*app.Commit, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *CommitDB) TableName() string {
	return "commits"

}

// CRUD Functions

// Get returns a single Commit as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *CommitDB) Get(ctx context.Context, id int) (*Commit, error) {
	defer goa.MeasureSince([]string{"goa", "db", "commit", "get"}, time.Now())

	var native Commit
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Commit
func (m *CommitDB) List(ctx context.Context) ([]*Commit, error) {
	defer goa.MeasureSince([]string{"goa", "db", "commit", "list"}, time.Now())

	var objs []*Commit
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *CommitDB) Add(ctx context.Context, model *Commit) error {
	defer goa.MeasureSince([]string{"goa", "db", "commit", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Commit", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *CommitDB) Update(ctx context.Context, model *Commit) error {
	defer goa.MeasureSince([]string{"goa", "db", "commit", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Commit", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *CommitDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "commit", "delete"}, time.Now())

	var obj Commit

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Commit", "error", err.Error())
		return err
	}

	return nil
}
