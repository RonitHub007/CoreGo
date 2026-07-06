package ports

import "library-system/internal/core/domain"

// BookRepository defines the secondary port (outbound adapter interface).
// The core application uses this to talk to the database without knowing it's Postgres or GORM.
type BookRepository interface {
	CreateBook(book *domain.Book) error
	GetAllBooks() ([]domain.Book, error)
	GetBookByID(id uint) (*domain.Book, error)
	UpdateBook(book *domain.Book) error
	DeleteBook(id uint) error
}
