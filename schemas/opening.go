package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   float64
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"localtion"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    float64   `json:"salary"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}
