package models

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

// Remaining struct for Remainingsitory in DB
type Remaining struct {
	ID        int `gorm:"primary_key"` // primary key
	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
	CrawlType string
	LastID    int
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Remaining) TableName() string {
	return "remaining"

}

// Remaining.
type RemainingDB struct {
	Db *gorm.DB
}

// NewRemainingDB creates a new storage type.
func NewRemainingDB(db *gorm.DB) *RemainingDB {
	return &RemainingDB{Db: db}
}

// DB returns the underlying database.
func (m *RemainingDB) DB() interface{} {
	return m.Db
}

// CRUD Functions

// Tablename returns the table name
func (m *RemainingDB) TableName() string {
	return "remaining"

}

// Get returns a single Remaining as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *RemainingDB) Get(ctx context.Context, id int) (*Remaining, error) {

	var native Remaining
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// Get returns a single Remaining as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *RemainingDB) GetWhereCrawlType(ctx context.Context, crawlType string) (*Remaining, error) {

	var native Remaining
	err := m.Db.Table(m.TableName()).Where("crawl_type = ?", crawlType).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Remaining
func (m *RemainingDB) List(ctx context.Context) ([]*Remaining, error) {

	var objs []*Remaining
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *RemainingDB) Add(ctx context.Context, model *Remaining) error {

	err := m.Db.Create(model).Error
	if err != nil {

		return err
	}

	return nil
}

// Update modifies a single record.
func (m *RemainingDB) Update(ctx context.Context, model *Remaining) error {

	obj, err := m.Get(ctx, model.ID)
	if err != nil {

		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *RemainingDB) Delete(ctx context.Context, id int) error {

	var obj Remaining

	err := m.Db.Delete(&obj, id).Error

	if err != nil {

		return err
	}

	return nil
}
