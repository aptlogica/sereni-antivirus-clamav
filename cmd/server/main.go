package main

import (
	"log"
	"sereni-antivirus/internal/config"
	"sereni-antivirus/internal/handlers"
	"sereni-antivirus/internal/providers/antivirus"
	"sereni-antivirus/internal/routes"
	"sereni-antivirus/internal/services"
)

// @title Sereni Antivirus API
// @version 1.0
// @description Microservice for antivirus scanning
// @host localhost:8084
// @BasePath /
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Provider
	avProvider, err := antivirus.NewAntivirus(cfg.Antivirus)
	if err != nil {
		log.Fatalf("Failed to initialize antivirus provider: %v", err)
	}

	// Initialize Service
	avService := services.NewAntivirusService(avProvider)

	// Initialize Handler
	scanHandler := handlers.NewScanHandler(avService)

	// Setup Routes
	r := routes.SetupRouter(scanHandler)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
