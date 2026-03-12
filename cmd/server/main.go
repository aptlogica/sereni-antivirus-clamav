/*
Copyright (c) 2026 Aptlogica Technologies Private Limited
This file is part of software developed by Aptlogica Technologies Private Limited.
Licensed under the MIT License. See the LICENSE file in the project root
for full license information.
Websites:
https://www.aptlogica.com
https://www.serenibase.com
Support:
support@aptlogica.com
support@serenibase.com
*/

package main

import (
	"fmt"
	"log"
	"sereni-antivirus/internal/config"
	"sereni-antivirus/internal/handlers"
	"sereni-antivirus/internal/providers/antivirus"
	"sereni-antivirus/internal/routes"
	"sereni-antivirus/internal/services"

	"sereni-antivirus/docs"
)

// @title Sereni Antivirus API
// @version 1.0
// @description Microservice for antivirus scanning
// @host ${HOST}:${PORT}
// @BasePath /
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Set dynamic host for Swagger
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	// Initialize Provider
	avProvider, err := antivirus.NewAntivirus(cfg.Antivirus)
	if err != nil {
		log.Fatalf("Failed to initialize antivirus provider: %v", err)
	}

	// Initialize Service
	avService := services.NewAntivirusService(avProvider)

	// Initialize Handler
	scanHandler := handlers.NewScanHandler(avService, cfg.MaxUploadSizeBytes)

	// Setup Routes
	r := routes.SetupRouter(scanHandler)

	log.Printf("Starting server on %s:%s", cfg.Host, cfg.Port)
	if err := r.Run(cfg.Host + ":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
