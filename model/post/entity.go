package post

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
	PostID       int64     `json:"post_id"            gorm:"column:post_id;primary_key"`
	Text         string    `json:"text"               gorm:"column:text"             sql:"not null;type:varchar(200)"`
	ActiveStatus int       `json:"active_status"      gorm:"column:active_status"    sql:"not null;type:tinyint(3)"`
	CreatedAt    time.Time `json:"created_at"         gorm:"column:created_at"       sql:"not null;type:datetime"`
	UpdatedAt    time.Time `json:"updated_at"         gorm:"column:updated_at"       sql:"not null;type:datetime"`
}

// New ...
func New(name string) *Entity {
	return &Entity{
		Text:         name,
		ActiveStatus: Active,
		CreatedAt:    libtime.Now(),
		UpdatedAt:    libtime.Now(),
	}
}

// TableName returns the table name
func (e Entity) TableName() string {
	return "posts"
}

// Validate ...
func (e *Entity) Validate() error {
	return nil
}

// Primary returns
func (e *Entity) Primary() (interface{}, []interface{}) {
	return "post_id = ?", []interface{}{e.ID}
}
