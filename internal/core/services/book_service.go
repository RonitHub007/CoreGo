package services

import (
	"errors"
	"library-system/internal/core/domain"
	"library-system/internal/core/ports"

	"go.uber.org/zap"
)

// bookService is the concrete implementation of the BookService port.
type bookService struct {
	repo   ports.BookRepository
	logger *zap.Logger
}

// NewBookService creates a new instance of BookService, injecting dependencies.
func NewBookService(repo ports.BookRepository, logger *zap.Logger) ports.BookService {
	return &bookService{
		repo:   repo,
		logger: logger,
	}
}

func (s *bookService) AddBook(book *domain.Book) error {
	s.logger.Debug("Validating new book", zap.String("title", book.Title))
	if book.Title == "" {
		return errors.New("book title cannot be empty")
	}
	if book.Author == "" {
		return errors.New("book author cannot be empty")
	}
	if book.ISBN == "" {
		return errors.New("book ISBN cannot be empty")
	}

	s.logger.Debug("Book validation passed, sending to repository")
	return s.repo.CreateBook(book)
}

func (s *bookService) FetchAllBooks() ([]domain.Book, error) {
	return s.repo.GetAllBooks()
}

func (s *bookService) FetchBookByID(id uint) (*domain.Book, error) {
	return s.repo.GetBookByID(id)
}

func (s *bookService) ModifyBook(id uint, updatedData *domain.Book) error {
	s.logger.Debug("Fetching existing book for modification", zap.Uint("id", id))
	existingBook, err := s.repo.GetBookByID(id)
	if err != nil {
		return errors.New("book not found")
	}

	if updatedData.Title != "" {
		existingBook.Title = updatedData.Title
	}
	if updatedData.Author != "" {
		existingBook.Author = updatedData.Author
	}
	if updatedData.PublishedYear != 0 {
		existingBook.PublishedYear = updatedData.PublishedYear
	}
	if updatedData.ISBN != "" {
		existingBook.ISBN = updatedData.ISBN
	}

	s.logger.Debug("Sending updated book to repository", zap.Uint("id", id))
	return s.repo.UpdateBook(existingBook)
}

func (s *bookService) RemoveBook(id uint) error {
	return s.repo.DeleteBook(id)
}
