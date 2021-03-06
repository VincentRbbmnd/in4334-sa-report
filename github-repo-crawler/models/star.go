package models

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

// Star struct for Star repository in DB
type Star struct {
	ID        int `gorm:"primary_key"` // primary key
	UserID    int64
	RepoID    int64
	StarredAt time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Star) TableName() string {
	return "stars"
}

// Star.
type StarDB struct {
	Db *gorm.DB
}

// NewStarDB creates a new storage type.
func NewStarDB(db *gorm.DB) *StarDB {
	return &StarDB{Db: db}
}

// DB returns the underlying database.
func (m *StarDB) DB() interface{} {
	return m.Db
}

// CRUD Functions

// Tablename returns the table name
func (m *StarDB) TableName() string {
	return "stars"
}

// Get returns a single User as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *StarDB) Get(ctx context.Context, id int) (*Star, error) {

	var native Star
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

func (m *StarDB) ListStars(ctx context.Context, id int64) ([]*Star, error) {

	var objs []*Star
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil {
		return nil, err
	}

	return objs, err
}

func (m *StarDB) ListByGithubUserLogin(ctx context.Context, login string) ([]*Star, error) {

	var objs []*Star
	//TODO this doesn't work yet
	//Join tables user and stars
	err := m.Db.Table(m.TableName()).Where("github_user_login = ?", login).Find(&objs).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, err
}

func (m *StarDB) ListByGithubRepositoryID(ctx context.Context, id int64) ([]*Star, error) {

	var objs []*Star
	//TODO this doesn't work yet
	//Join tables repo and stars
	err := m.Db.Table(m.TableName()).Where("github_repository_id = ?", id).Find(&objs).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, err
}

// Add creates a new record.
func (m *StarDB) Add(ctx context.Context, star *Star) error {

	err := m.Db.Create(star).Error
	if err != nil {
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *StarDB) Update(ctx context.Context, model *Star) error {

	obj, err := m.Get(ctx, model.ID)
	if err != nil {

		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *StarDB) Delete(ctx context.Context, id int) error {

	var obj Star

	err := m.Db.Delete(&obj, id).Error

	if err != nil {

		return err
	}

	return nil
}
