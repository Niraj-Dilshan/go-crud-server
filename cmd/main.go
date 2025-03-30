package main

import (
	"crud-server/config"
	"crud-server/startup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig(".env")

	// Initialize module
	module := startup.NewModule(cfg)

	// Initialize database
	if err := module.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize database schema
	if err := startup.InitIndexes(module.GetDB()); err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	// Create server
	server := startup.NewServer(cfg)
	router := server.GetRouter()

	// Initialize API
	api := router.Group("/api")
	module.InitTaskAPI(api)

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-quit
	log.Println("Shutting down server...")
}
