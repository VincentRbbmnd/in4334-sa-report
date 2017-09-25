package models

import (
	"context"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// Repo struct for repository in DB
type Repo struct {
	ID        int    `gorm:"primary_key"` // primary key
	Owner     string // has many Repo
	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
	Org       bool
	ProjectID int64
	UserType  string
	FullName  string // timestamp
	Raw       []byte `sql:"type:jsonb"` // This is the RAW JSONB of the metadata of a Repo
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Repo) TableName() string {
	return "repositories"

}

// Repo.
type RepoDB struct {
	Db *gorm.DB
}

// NewRepoDB creates a new storage type.
func NewRepoDB(db *gorm.DB) *RepoDB {
	return &RepoDB{Db: db}
}

// DB returns the underlying database.
func (m *RepoDB) DB() interface{} {
	return m.Db
}

// CRUD Functions

// Tablename returns the table name
func (m *RepoDB) TableName() string {
	return "repositories"

}

// Get returns a single Repo as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *RepoDB) Get(ctx context.Context, id int64) (*Repo, error) {
	defer goa.MeasureSince([]string{"goa", "db", "Repo", "get"}, time.Now())

	var native Repo
	err := m.Db.Table(m.TableName()).Where("project_id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Repo
func (m *RepoDB) List(ctx context.Context) ([]*Repo, error) {
	defer goa.MeasureSince([]string{"goa", "db", "Repo", "list"}, time.Now())

	var objs []*Repo
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *RepoDB) Add(ctx context.Context, model *Repo) error {
	defer goa.MeasureSince([]string{"goa", "db", "Repo", "add"}, time.Now())

	m.Db.Create(model)
	// if err != nil {
	// 	goa.LogError(ctx, "error adding Repo", "error", err.Error())
	// 	return err
	// }

	return nil
}

// Update modifies a single record.
func (m *RepoDB) Update(ctx context.Context, model *Repo) error {
	defer goa.MeasureSince([]string{"goa", "db", "Repo", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ProjectID)
	if err != nil {
		goa.LogError(ctx, "error updating Repo", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *RepoDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "Repo", "delete"}, time.Now())

	var obj Repo

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Repo", "error", err.Error())
		return err
	}

	return nil
}
