package models

import "time"

type Experience struct {
	// Id          string     `json:"id" db:"id"`
	Role        string     `json:"role" db:"role"`
	Description string     `json:"description" db:"description"`
	Company     string     `json:"company" db:"company"`
	StartDate   time.Time  `json:"start_date" db:"start_date"`
	EndDate     *time.Time `json:"end_date" db:"end_date"`
	IsActive    bool       `json:"is_active" db:"is_active"`
	TechStack   []string   `json:"tech_stack" db:"tech_stack"`
	Link        string     `json:"link" db:"link"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}
