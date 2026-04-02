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
	"github.com/aptlogica/sereni-antivirus-clamav/internal/config"
	"github.com/aptlogica/sereni-antivirus-clamav/internal/handlers"
	"github.com/aptlogica/sereni-antivirus-clamav/internal/providers/antivirus"
	"github.com/aptlogica/sereni-antivirus-clamav/internal/routes"
	"github.com/aptlogica/sereni-antivirus-clamav/internal/services"

	"github.com/aptlogica/sereni-antivirus-clamav/docs"
)

// validateSecrets checks that critical environment variables are set
func validateSecrets() error {
	required := []string{"CLAMAV_HOST"}
	for _, key := range required {
		if val := config.GetEnv(key, ""); val == "" {
			return fmt.Errorf("required environment variable %s is not set", key)
		}
	}
	return nil
}

// @title Sereni Antivirus API
// @version 1.0
// @description Microservice for antivirus scanning
// @host ${HOST}:${PORT}
// @BasePath /
func main() {
	// Validate required secrets before proceeding
	if err := validateSecrets(); err != nil {
		log.Fatalf("Secret validation failed: %v", err)
	}

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
