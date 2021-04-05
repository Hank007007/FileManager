package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/google/uuid"
)

type User struct {
	//ID        string `gorm:"type:uuid;primary_key;default:gen_random_uuid()"` // for gen_random_uuid for postgreS 13
	ID        string `gorm:primary_key;column:user_id;type:VARCHAR;"`
	Name      string `gorm:"column:user_nm;type:VARCHAR(30);"`
	CreatedAt time.Time `gorm:"column:create_dttm;type:TIMESTAMP;`
	UpdatedAt time.Time `gorm:"column:mod_dttm;type:TIMESTAMP;`
	CreateUserNm         string    `gorm:"column:create_user_nm;type:VARCHAR;"`
	ModUserNm            string    `gorm:"column:mod_user_nm;type:VARCHAR;"`
	DeletedAt *time.Time
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *User) BeforeCreate(scope *gorm.Scope) error {
	m.CreatedAt = time.Now().UTC().Truncate(time.Microsecond)
	m.UpdatedAt = time.Now().UTC().Truncate(time.Microsecond)

	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return nil
}

// BeforeUpdate invoked before updating, return an error if field is not populated.
func (m *User) BeforeUpdate(scope *gorm.Scope) error {
	m.UpdatedAt = time.Now().UTC().Truncate(time.Microsecond)
	return nil
}
