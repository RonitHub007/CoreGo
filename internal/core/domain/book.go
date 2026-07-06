package domain

import "time"

// Book represents the core business entity.
// In Hexagonal Architecture, this struct MUST NOT import any external frameworks 
// like GORM, Gin, or database drivers. It must be pure Go.
type Book struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	ISBN          string    `json:"isbn"`
	PublishedYear int       `json:"published_year"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
