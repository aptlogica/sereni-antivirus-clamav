package handlers

import (
	"net/http"
	"path/filepath"
	"sereni-antivirus/internal/services"

	"github.com/gin-gonic/gin"
)

type ScanHandler struct {
	service services.AntivirusService
}

func NewScanHandler(service services.AntivirusService) *ScanHandler {
	return &ScanHandler{service: service}
}

// ScanFile godoc
// @Summary Scan a file for viruses
// @Description Upload a file to scan it for malware using ClamAV
// @Tags antivirus
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to scan"
// @Success 200 {object} interfaces.ScanResult
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /scan [post]
func (h *ScanHandler) ScanFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	result, err := h.service.ScanFile(c.Request.Context(), filepath.Base(header.Filename), file)

	// Check if the result indicates a threat, even if an error is returned (as per provider implementation)
	if !result.Clean && result.Threat != "" {
		// Malware found
		c.JSON(http.StatusOK, result)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// ScanFiles godoc
// @Summary Scan multiple files for viruses
// @Description Upload multiple files to scan them for malware
// @Tags antivirus
// @Accept multipart/form-data
// @Produce json
// @Param files formData []file true "Files to scan" collectionFormat multi
// @Success 200 {array} interfaces.ScanResult
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /scan-files [post]
func (h *ScanHandler) ScanFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	results, err := h.service.ScanFiles(c.Request.Context(), files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
