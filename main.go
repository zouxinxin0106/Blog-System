package main

import (
	"blog-system/config"
	"blog-system/database"
	"blog-system/router"
	"fmt"
	"log"
)

func main() {
	// Initialize database
	_, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup router
	r := router.SetupRouter()

	// Start server
	addr := fmt.Sprintf(":%s", config.Port)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
