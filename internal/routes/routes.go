package routes

import (
	_ "sereni-antivirus/docs"
	"sereni-antivirus/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(scanHandler *handlers.ScanHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "antivirus",
		})
	})

	r.POST("/scan", scanHandler.ScanFile)
	r.POST("/scan-files", scanHandler.ScanFiles)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
