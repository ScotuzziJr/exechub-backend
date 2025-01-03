package models

type Role int

const (
	Headhunter Role = iota
	Mentor
	Professional
	C_Level
)

// represents lead model
type Lead struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}
