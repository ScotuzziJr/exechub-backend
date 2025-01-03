package models

import "time"

type Role int

const (
	Headhunter Role = iota // set index to 0 and auto increment by 1
	Mentor
	Professional
	C_Level
)

// represents lead model
type Lead struct {
	ID        string    `gorm:"primaryKey" json:"id,omitempty"` // omitempty - allows to auto generate ID on server side
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
