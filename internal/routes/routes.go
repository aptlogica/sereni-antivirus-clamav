// Package routes sets up HTTP routes for the application.
package routes

import (
	_ "sereni-antivirus/docs"
	"sereni-antivirus/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter configures and returns the Gin router with all routes.
func SetupRouter(scanHandler *handlers.ScanHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/scan", scanHandler.ScanFile)
	r.POST("/scan-files", scanHandler.ScanFiles)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
