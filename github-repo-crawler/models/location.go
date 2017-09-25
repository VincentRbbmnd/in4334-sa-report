package models

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Location struct for Locationsitory in DB
type Location struct {
	ID       int    `gorm:"primary_key"` // primary key
	Location string `sql:"type:geometry(Point,4326)"`
	Raw      []byte `sql:"type:jsonb"` // This is the RAW JSONB of the metadata of a Location
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Location) TableName() string {
	return "locations"

}

// Location.
type LocationDB struct {
	Db *gorm.DB
}

// NewLocationDB creates a new storage type.
func NewLocationDB(db *gorm.DB) *LocationDB {
	return &LocationDB{Db: db}
}

// DB returns the underlying database.
func (m *LocationDB) DB() interface{} {
	return m.Db
}

// CRUD Functions

// Tablename returns the table name
func (m *LocationDB) TableName() string {
	return "locations"
}

// Get returns a single Location as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *LocationDB) Get(ctx context.Context, id int) (*Location, error) {

	var native Location
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Location
func (m *LocationDB) List(ctx context.Context) ([]*Location, error) {

	var objs []*Location
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *LocationDB) Add(ctx context.Context, model *Location) error {

	err := m.Db.Create(model).Error
	if err != nil {

		return err
	}

	return nil
}

// Update modifies a single record.
func (m *LocationDB) Update(ctx context.Context, model *Location) error {

	obj, err := m.Get(ctx, model.ID)
	if err != nil {

		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *LocationDB) Delete(ctx context.Context, id int) error {

	var obj Location

	err := m.Db.Delete(&obj, id).Error

	if err != nil {

		return err
	}

	return nil
}
