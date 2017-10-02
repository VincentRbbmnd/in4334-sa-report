package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// User struct for Usersitory in DB
type User struct {
	ID              int    `gorm:"primary_key"` // primary key
	Login           string //user name
	CreatedAt       time.Time
	DeletedAt       *time.Time
	UpdatedAt       time.Time
	GithubUserID    int64
	Type            string //type of user
	Raw             []byte `sql:"type:jsonb"` // This is the RAW JSONB of the metadata of a User
	LocationChecked bool
	LocationID      int
	Stars           []Star
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m User) TableName() string {
	return "users"

}

// User.
type UserDB struct {
	Db *gorm.DB
}

// NewUserDB creates a new storage type.
func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{Db: db}
}

// DB returns the underlying database.
func (m *UserDB) DB() interface{} {
	return m.Db
}

// CRUD Functions

// Tablename returns the table name
func (m *UserDB) TableName() string {
	return "users"
}

// Get returns a single User as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserDB) Get(ctx context.Context, id int) (*User, error) {

	var native User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

func (m *UserDB) GetByGithubID(ctx context.Context, id int64) (*User, error) {

	var native User
	err := m.Db.Table(m.TableName()).Where("github_user_id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of User
func (m *UserDB) List(ctx context.Context) ([]*User, error) {

	var objs []*User
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// List returns an array of User
func (m *UserDB) ListNoLocations(ctx context.Context) ([]*User, error) {

	var objs []*User
	err := m.Db.Table(m.TableName()).Select("id,github_user_id,login").Where("location_checked is null OR location_checked = false").Limit(1000).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// ListNoLocationsForRepo lists users for certain repo
func (m *UserDB) ListNoLocationsForRepo(ctx context.Context, PID int64) []*User {
	var res []*User
	err := m.Db.Table("repositories").Select("id, github_user_id, login").
		Joins(`LEFT JOIN "Commits" ON repository_id = project_id`).
		Joins(`LEFT JOIN users ON github_user_id = author_id`).
		Where("(location_checked is null OR location_checked = false) AND project_id = ?", PID).Limit(100).
		Find(&res).Error
	if err != nil {
		fmt.Println("Wat is aan die handje", err)
		return res
	}
	return res
	//TODO use repo to get users
	//commitDB get commits where repo = current repo
	//Join with users op author id where locationchecked = false
}

// Add creates a new record.
func (m *UserDB) Add(ctx context.Context, model *User) error {

	m.Db.Create(model)
	// if err != nil {
	//
	// 	return err
	// }

	return nil
}

// Update modifies a single record.
func (m *UserDB) Update(ctx context.Context, model *User) error {

	obj, err := m.Get(ctx, model.ID)
	if err != nil {

		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserDB) Delete(ctx context.Context, id int) error {

	var obj User

	err := m.Db.Delete(&obj, id).Error

	if err != nil {

		return err
	}

	return nil
}
