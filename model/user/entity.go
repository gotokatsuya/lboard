package user

import (
	"time"

	libtime "github.com/gotokatsuya/lboard/library/time"
)

const (
	// Active ...
	// flag : user is active
	Active = 1

	// InActive ...
	// flag : user is inactive
	InActive = 2
)

// Entity describes user table.
type Entity struct {
	UserID       int64     `json:"user_id"            gorm:"column:user_id;primary_key"`
	TagName      string    `json:"tag_name"           gorm:"column:tag_name"         sql:"not null;type:varchar(190)"`
	ActiveStatus int       `json:"active_status"      gorm:"column:active_status"    sql:"not null;type:tinyint(3)"`
	CreatedAt    time.Time `json:"created_at"         gorm:"column:created_at"       sql:"not null;type:datetime"`
	UpdatedAt    time.Time `json:"updated_at"         gorm:"column:updated_at"       sql:"not null;type:datetime"`
}

// New ...
func New(tagName string) *Entity {
	return &Entity{
		TagName:      tagName,
		ActiveStatus: Active,
		CreatedAt:    libtime.Now(),
		UpdatedAt:    libtime.Now(),
	}
}

// TableName returns the table name
func (e Entity) TableName() string {
	return "users"
}

// Validate ...
func (e *Entity) Validate() error {
	return nil
}

// Primary returns
func (e *Entity) Primary() (interface{}, []interface{}) {
	return "user_id = ?", []interface{}{e.UserID}
}
