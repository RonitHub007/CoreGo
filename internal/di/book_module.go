package di

import (
	"library-system/internal/adapters/handler/http"
	"library-system/internal/adapters/repository/postgres"
	"library-system/internal/core/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// InitBookModule encapsulates all the Dependency Injection wiring for the Book domain.
// It creates the repository, service, and handler, and registers the routes.
func InitBookModule(db *gorm.DB, logger *zap.Logger, router *gin.Engine) {
	// 1. Auto-Migrate the database schema for Books
	// This ensures the database is ready before the app starts accepting requests.
	if err := db.AutoMigrate(&postgres.BookDBModel{}); err != nil {
		logger.Fatal("Failed to migrate book database schema", zap.Error(err))
	}

	// 2. Wire up Ports and Adapters
	// Database Adapter (Secondary)
	bookRepo := postgres.NewBookRepository(db, logger)
	
	// Core Service (Domain)
	bookService := services.NewBookService(bookRepo, logger)
	
	// HTTP Handler Adapter (Primary)
	bookHandler := http.NewBookHandler(bookService, logger)

	// 3. Register HTTP Routes
	// We group related routes (e.g. /books) inside the module setup.
	router.POST("/books", bookHandler.AddBook)
	
	logger.Info("Book module initialized successfully")
}
