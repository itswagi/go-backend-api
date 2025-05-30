package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itswagi/go-backend-api/internal/logger"
	"github.com/rs/cors"
)


func StartServer(router http.Handler, log *logger.Logger) {
	// Configure CORS settings
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Set your allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},                         // Specify allowed methods
		AllowedHeaders:   []string{"Authorization", "Content-Type"},                        // Specify allowed headers
		ExposedHeaders:   []string{"Authorization"},                                        // Specify exposed headers
		AllowCredentials: true,                                                             // Allow credentials like cookies
		Debug:            false,                                                            // Set to true to debug CORS requests
	})

	// Wrap the router with the CORS middleware
	handlerWithCORS := corsOptions.Handler(router)

	// Create an HTTP server
	srv := &http.Server{
		Addr:         ":4000",
		Handler:      handlerWithCORS, // Use the handler with CORS
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run the server in a goroutine to enable graceful shutdown
	go func() {
		log.Info("Server running on port 4000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %v", err)
		}
	}()

	// Graceful shutdown handling
	gracefulShutdown(srv, log)
}

func gracefulShutdown(srv *http.Server, log *logger.Logger) {
	// Capture termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-signalChan
	log.Info("Shutdown signal received, shutting down server...")

	// Create a context with a timeout to gracefully shut down the server
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Server forced to shutdown: %v", err)
	}

	log.Info("Server shutdown complete.")
}