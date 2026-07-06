package postgres

import (
	"library-system/internal/core/domain"

	"gorm.io/gorm"
)

// BookDBModel represents the database schema for a book.
// It exclusively uses GORM tags. This keeps our domain layer pure.
type BookDBModel struct {
	gorm.Model
	Title         string `gorm:"not null"`
	Author        string `gorm:"not null"`
	ISBN          string `gorm:"uniqueIndex;not null"`
	PublishedYear int
}

// TableName overrides the default GORM table name to ensure consistency.
func (BookDBModel) TableName() string {
	return "books"
}

// toDomain converts a Database model into a Pure Domain model.
func (dbm *BookDBModel) toDomain() domain.Book {
	return domain.Book{
		ID:            dbm.ID,
		Title:         dbm.Title,
		Author:        dbm.Author,
		ISBN:          dbm.ISBN,
		PublishedYear: dbm.PublishedYear,
		CreatedAt:     dbm.CreatedAt,
		UpdatedAt:     dbm.UpdatedAt,
	}
}

// fromDomain converts a Pure Domain model into a Database model.
func fromDomain(d *domain.Book) *BookDBModel {
	return &BookDBModel{
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Title:         d.Title,
		Author:        d.Author,
		ISBN:          d.ISBN,
		PublishedYear: d.PublishedYear,
	}
}
