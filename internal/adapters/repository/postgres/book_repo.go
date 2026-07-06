package postgres

import (
	"library-system/internal/core/domain"
	"library-system/internal/core/ports"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// bookRepository is the Postgres adapter implementation of ports.BookRepository.
type bookRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

// NewBookRepository creates a new Postgres adapter for the Book Repository.
func NewBookRepository(db *gorm.DB, logger *zap.Logger) ports.BookRepository {
	return &bookRepository{
		db:     db,
		logger: logger,
	}
}

func (r *bookRepository) CreateBook(book *domain.Book) error {
	r.logger.Info("Attempting to create book in Postgres", zap.String("isbn", book.ISBN))

	dbModel := fromDomain(book)
	err := r.db.Create(dbModel).Error
	if err != nil {
		r.logger.Error("Failed to create book in Postgres", zap.Error(err))
		return err
	}

	// Update domain model with DB generated fields (like ID)
	book.ID = dbModel.ID
	book.CreatedAt = dbModel.CreatedAt
	book.UpdatedAt = dbModel.UpdatedAt

	return nil
}

func (r *bookRepository) GetAllBooks() ([]domain.Book, error) {
	r.logger.Info("Fetching all books from Postgres")
	var dbModels []BookDBModel
	result := r.db.Find(&dbModels)
	if result.Error != nil {
		return nil, result.Error
	}

	var books []domain.Book
	for _, dbm := range dbModels {
		books = append(books, dbm.toDomain())
	}
	return books, nil
}

func (r *bookRepository) GetBookByID(id uint) (*domain.Book, error) {
	r.logger.Info("Fetching book by ID from Postgres", zap.Uint("id", id))
	var dbModel BookDBModel
	result := r.db.First(&dbModel, id)
	if result.Error != nil {
		r.logger.Warn("Book not found in Postgres", zap.Uint("id", id), zap.Error(result.Error))
		return nil, result.Error
	}

	domainBook := dbModel.toDomain()
	return &domainBook, nil
}

func (r *bookRepository) UpdateBook(book *domain.Book) error {
	r.logger.Info("Updating book in Postgres", zap.String("isbn", book.ISBN))
	dbModel := fromDomain(book)
	return r.db.Save(dbModel).Error
}

func (r *bookRepository) DeleteBook(id uint) error {
	r.logger.Info("Deleting book from Postgres", zap.Uint("id", id))
	return r.db.Delete(&BookDBModel{}, id).Error
}
