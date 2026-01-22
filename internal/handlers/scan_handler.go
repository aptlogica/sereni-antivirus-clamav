// Package handlers provides HTTP request handlers for the antivirus service.
package handlers

import (
	"net/http"
	"path/filepath"
	"sereni-antivirus/internal/services"

	"github.com/gin-gonic/gin"
)

// ScanHandler handles antivirus scan requests.
type ScanHandler struct {
	service        services.AntivirusService
	maxUploadBytes int64
}

// NewScanHandler creates a new scan handler with the given service and upload limit.
func NewScanHandler(service services.AntivirusService, maxUploadBytes int64) *ScanHandler {
	return &ScanHandler{service: service, maxUploadBytes: maxUploadBytes}
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
	h.LimitRequestBody(c)
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		if h.HandleBodyTooLargeError(c, err) {
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	result, err := h.service.ScanFile(c.Request.Context(), filepath.Base(header.Filename), file)
	if err != nil && result.Threat == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !result.Clean || result.Threat != "" {
		c.JSON(http.StatusOK, result)
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
	h.LimitRequestBody(c)
	form, err := c.MultipartForm()
	if err != nil {
		if h.HandleBodyTooLargeError(c, err) {
			return
		}
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

func (h *ScanHandler) LimitRequestBody(c *gin.Context) {
	if h.maxUploadBytes <= 0 {
		return
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, h.maxUploadBytes)
}

func (h *ScanHandler) HandleBodyTooLargeError(c *gin.Context, err error) bool {
	if err != nil && err.Error() == "http: request body too large" {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "Request body too large"})
		return true
	}
	return false
}
