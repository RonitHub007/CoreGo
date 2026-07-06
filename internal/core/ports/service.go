package ports

import "library-system/internal/core/domain"

// BookService defines the primary port (inbound adapter interface).
// Controllers/Handlers use this to interact with the core business logic.
type BookService interface {
	AddBook(book *domain.Book) error
	FetchAllBooks() ([]domain.Book, error)
	FetchBookByID(id uint) (*domain.Book, error)
	ModifyBook(id uint, updatedData *domain.Book) error
	RemoveBook(id uint) error
}
