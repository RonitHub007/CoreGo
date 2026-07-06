package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"library-system/internal/server"
)

func main() {
	// 1. Initialize the Server Container
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("Failed to initialize server container: %v", err)
	}

	// 2. Start the server in a background goroutine
	// This prevents the main thread from blocking, allowing us to listen for OS signals.
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("Server stopped unexpectedly: %v", err)
		}
	}()

	// 3. Graceful Shutdown Configuration
	// Create a channel to listen for OS signals (like Ctrl+C or Docker shutdown)
	quit := make(chan os.Signal, 1)
	
	// Notify the channel for SIGINT and SIGTERM signals
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// The main thread blocks here, waiting until a signal is sent to the 'quit' channel
	<-quit
	log.Println("Interrupt signal received...")

	// 4. Create a context with a 5-second timeout
	// This gives active HTTP requests 5 seconds to finish before we forcefully kill them.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5. Trigger the Shutdown logic in our Server container
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Application exited cleanly. Goodbye!")
}

