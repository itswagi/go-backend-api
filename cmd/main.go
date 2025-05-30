package main

import (
	"github.com/gorilla/mux"
	"github.com/itswagi/go-backend-api/internal/user"
	"github.com/itswagi/go-backend-api/internal/logger"
	"github.com/itswagi/go-backend-api/internal/server"
)

func main() {
	// Set up logger
	log := logger.NewLogger()

	router := mux.NewRouter()

	router.Use(log.LoggingMiddleware)

	// Like NestJS module registration
	user.RegisterUserModule(router)

	// Start the HTTP server with graceful shutdown
	server.StartServer(router, log)
}
