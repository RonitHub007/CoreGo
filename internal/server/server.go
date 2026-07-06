package server

import (
	"context"
	"errors"
	nethttp "net/http" // Alias to avoid collision with our adapter package

	"library-system/internal/config"
	"library-system/internal/database"
	"library-system/internal/di"
	"library-system/internal/logger"
	"library-system/internal/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Server is the core container for the application.
type Server struct {
	router *gin.Engine
	logger *zap.Logger
	db     *gorm.DB
	srv    *nethttp.Server
}

// NewServer initializes all dependencies and wires them together.
func NewServer() (*Server, error) {
	// 1. Initialize Logger
	zapLogger, err := logger.New()
	if err != nil {
		return nil, err
	}

	// 2. Load Configuration
	cfg := config.Load()

	// 3. Connect to Database
	db, err := database.ConnectDB(cfg.DatabaseDSN)
	if err != nil {
		zapLogger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	// 4. Initialize Router & Global Middleware
	router := gin.New()
	router.Use(middleware.LoggerMiddleware(zapLogger))
	router.Use(gin.Recovery())

	// 5. Initialize Modules (Dependency Injection)
	// This keeps our server.go beautifully thin. All the "ugly" wiring
	// is hidden inside specific module initializers.
	di.InitBookModule(db, zapLogger, router)
	
	// If you add a User module later, you just add one line:
	// di.InitUserModule(db, zapLogger, router)

	// 6. Setup HTTP Server configuration
	srv := &nethttp.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &Server{
		router: router,
		logger: zapLogger,
		db:     db,
		srv:    srv,
	}, nil
}

// Start begins the HTTP server.
func (s *Server) Start() error {
	s.logger.Info("Starting server on :8080")
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, nethttp.ErrServerClosed) {
		s.logger.Error("Server crashed", zap.Error(err))
		return err
	}
	return nil
}

// Shutdown gracefully stops the server and cleans up resources.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Initiating graceful shutdown...")

	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.Error("HTTP Server forced to shutdown", zap.Error(err))
	} else {
		s.logger.Info("HTTP Server shutdown cleanly")
	}

	sqlDB, err := s.db.DB()
	if err == nil {
		if err := sqlDB.Close(); err != nil {
			s.logger.Error("Failed to close database connection", zap.Error(err))
		} else {
			s.logger.Info("Database connection closed cleanly")
		}
	}

	_ = s.logger.Sync()
	return nil
}

