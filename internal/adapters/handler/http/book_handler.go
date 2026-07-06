package http

import (
	"net/http"

	"library-system/internal/core/domain"
	"library-system/internal/core/ports"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// BookHandler is the Primary Adapter for HTTP requests.
// It translates HTTP traffic into domain calls.
type BookHandler struct {
	// It depends purely on the Primary Port (BookService interface)
	bookService ports.BookService
	logger      *zap.Logger
}

// NewBookHandler creates a new BookHandler.
func NewBookHandler(s ports.BookService, logger *zap.Logger) *BookHandler {
	return &BookHandler{
		bookService: s,
		logger:      logger,
	}
}

// AddBook is the Gin Handler for POST /books.
func (h *BookHandler) AddBook(c *gin.Context) {
	h.logger.Info("Received AddBook HTTP request")
	var newBook domain.Book

	// 1. Parse JSON into Domain Model
	if err := c.ShouldBindJSON(&newBook); err != nil {
		h.logger.Warn("Failed to bind JSON in HTTP adapter", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format: " + err.Error()})
		return
	}

	// 2. Call the Domain Service Port
	err := h.bookService.AddBook(&newBook)
	if err != nil {
		h.logger.Error("Core service failed to add book", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. Respond
	h.logger.Info("Successfully fulfilled AddBook HTTP request", zap.String("isbn", newBook.ISBN))
	c.JSON(http.StatusCreated, gin.H{
		"message": "Book successfully created",
		"data":    newBook,
	})
}
