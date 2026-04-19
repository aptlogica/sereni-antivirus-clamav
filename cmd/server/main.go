/*
Copyright 2026-2030 Aptlogica Technologies Pvt Ltd
This file is part of software developed by Aptlogica Technologies Private Limited.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
