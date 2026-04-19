// Copyright 2026-2030 Aptlogica Technologies Pvt Ltd
// Licensed under the Apache License, Version 2.0
// Websites: https://www.aptlogica.com | https://www.serenibase.com
// Support: support@aptlogica.com | support@serenibase.com

// Package routes sets up HTTP routes for the application.
package routes

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/aptlogica/sereni-antivirus-clamav/docs"
	"github.com/aptlogica/sereni-antivirus-clamav/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter configures and returns the Gin router with all routes.
func SetupRouter(scanHandler *handlers.ScanHandler) *gin.Engine {
	r := gin.Default()

	// Configure CORS
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}

	corsConfig := cors.DefaultConfig()
	if allowedOrigins == "*" {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = strings.Split(allowedOrigins, ",")
	}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(corsConfig))

	r.POST("/scan", scanHandler.ScanFile)
	r.POST("/scan-files", scanHandler.ScanFiles)

	url := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", os.Getenv("HOST"), os.Getenv("PORT")))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
