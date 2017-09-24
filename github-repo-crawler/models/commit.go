package models

import (
	"context"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// Commit struct for Commitsitory in DB
type Commit struct {
	CreatedAt    time.Time
	DeletedAt    *time.Time
	UpdatedAt    time.Time
	CommitDate   time.Time
	Message      string
	SHA          string
	AuthorID     int64
	CommitterID  int64
	RepositoryID int
	Raw          []byte `sql:"type:jsonb"` // This is the RAW JSONB of the metadata of a Commit
}

// TableName overrshaes the table name settings in Gorm to force a specific table name
// in the database.
func (m Commit) TableName() string {
	return "Commits"

}

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

// CRUD Functions

// Tablename returns the table name
func (m *CommitDB) TableName() string {
	return "Commits"
}

// Get returns a single Commit as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *CommitDB) Get(ctx context.Context, sha string) (*Commit, error) {
	defer goa.MeasureSince([]string{"goa", "db", "Commit", "get"}, time.Now())

	var native Commit
	err := m.Db.Table(m.TableName()).Where("sha = ?", sha).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

func (m *CommitDB) GetLastCommitByRepoID(ctx context.Context, repoID int) (*Commit, error) {
	defer goa.MeasureSince([]string{"goa", "db", "Commit", "get"}, time.Now())

	var native Commit
	err := m.Db.Table(m.TableName()).Where("repository_id = ?", repoID).Order("created_at DESC LIMIT 1").Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Commit
func (m *CommitDB) List(ctx context.Context) ([]*Commit, error) {
	defer goa.MeasureSince([]string{"goa", "db", "Commit", "list"}, time.Now())

	var objs []*Commit
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *CommitDB) Add(ctx context.Context, model *Commit) error {
	defer goa.MeasureSince([]string{"goa", "db", "Commit", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Commit", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *CommitDB) Update(ctx context.Context, model *Commit) error {
	defer goa.MeasureSince([]string{"goa", "db", "Commit", "update"}, time.Now())

	obj, err := m.Get(ctx, model.SHA)
	if err != nil {
		goa.LogError(ctx, "error updating Commit", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *CommitDB) Delete(ctx context.Context, sha string) error {
	defer goa.MeasureSince([]string{"goa", "db", "Commit", "delete"}, time.Now())

	var obj Commit

	err := m.Db.Delete(&obj, sha).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Commit", "error", err.Error())
		return err
	}

	return nil
}
