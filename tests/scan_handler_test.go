// Copyright 2026-2030 Aptlogica Technologies Pvt Ltd
// Licensed under the Apache License, Version 2.0
// Websites: https://www.aptlogica.com | https://www.serenibase.com
// Support: support@aptlogica.com | support@serenibase.com

package tests

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aptlogica/sereni-antivirus-clamav/internal/handlers"
	"github.com/aptlogica/sereni-antivirus-clamav/internal/providers/antivirus/interfaces"

	"github.com/gin-gonic/gin"
)

type mockAntivirusService struct {
	ScanFileFn  func(ctx context.Context, name string, file io.Reader) (interfaces.ScanResult, error)
	ScanFilesFn func(ctx context.Context, files []*multipart.FileHeader) ([]interfaces.ScanResult, error)
}

func (m *mockAntivirusService) ScanFile(ctx context.Context, name string, file io.Reader) (interfaces.ScanResult, error) {
	return m.ScanFileFn(ctx, name, file)
}
func (m *mockAntivirusService) ScanFiles(ctx context.Context, files []*multipart.FileHeader) ([]interfaces.ScanResult, error) {
	return m.ScanFilesFn(ctx, files)
}

func TestScanHandler_ScanFile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAntivirusService{
		ScanFileFn: func(ctx context.Context, name string, file io.Reader) (interfaces.ScanResult, error) {
			if name == "bad.txt" {
				return interfaces.ScanResult{FileName: name, Clean: false, Threat: "virus"}, errors.New("virus found")
			}
			return interfaces.ScanResult{FileName: name, Clean: true, Threat: ""}, nil
		},
	}
	h := handlers.NewScanHandler(mockSvc, 1024*1024)
	r := gin.Default()
	r.POST("/scan", h.ScanFile)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "good.txt")
	part.Write([]byte("clean"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestScanHandler_ScanFile_Virus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAntivirusService{
		ScanFileFn: func(ctx context.Context, name string, file io.Reader) (interfaces.ScanResult, error) {
			return interfaces.ScanResult{FileName: name, Clean: false, Threat: "virus"}, nil
		},
	}
	h := handlers.NewScanHandler(mockSvc, 1024*1024)
	r := gin.Default()
	r.POST("/scan", h.ScanFile)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "bad.txt")
	part.Write([]byte("virus"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestScanHandler_ScanFile_NoLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAntivirusService{
		ScanFileFn: func(ctx context.Context, name string, file io.Reader) (interfaces.ScanResult, error) {
			return interfaces.ScanResult{FileName: name, Clean: true}, nil
		},
	}
	h := handlers.NewScanHandler(mockSvc, 0) // no limit
	r := gin.Default()
	r.POST("/scan", h.ScanFile)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "good.txt")
	part.Write([]byte("clean"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestScanHandler_ScanFile_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAntivirusService{
		ScanFileFn: func(ctx context.Context, name string, file io.Reader) (interfaces.ScanResult, error) {
			return interfaces.ScanResult{}, errors.New("service error")
		},
	}
	h := handlers.NewScanHandler(mockSvc, 1024*1024)
	r := gin.Default()
	r.POST("/scan", h.ScanFile)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.txt")
	part.Write([]byte("data"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

func TestScanHandler_ScanFiles_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAntivirusService{
		ScanFilesFn: func(ctx context.Context, files []*multipart.FileHeader) ([]interfaces.ScanResult, error) {
			return nil, errors.New("service error")
		},
	}
	h := handlers.NewScanHandler(mockSvc, 1024*1024)
	r := gin.Default()
	r.POST("/scan-files", h.ScanFiles)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("files", "test.txt")
	part.Write([]byte("data"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan-files", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

func TestScanHandler_limitRequestBody(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := handlers.NewScanHandler(nil, 10)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest(http.MethodPost, "/", nil)
	h.LimitRequestBody(c)
	// Should not panic or error
}

func TestScanHandler_handleBodyTooLargeError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := handlers.NewScanHandler(nil, 10)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	handled := h.HandleBodyTooLargeError(c, errors.New("http: request body too large"))
	if !handled {
		t.Fatalf("expected true for body too large error")
	}
	handled = h.HandleBodyTooLargeError(c, errors.New("other error"))
	if handled {
		t.Fatalf("expected false for other error")
	}
}

func TestScanHandler_ScanFiles(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAntivirusService{
		ScanFilesFn: func(ctx context.Context, files []*multipart.FileHeader) ([]interfaces.ScanResult, error) {
			if len(files) == 0 {
				return nil, errors.New("no files")
			}
			return []interfaces.ScanResult{{FileName: files[0].Filename, Clean: true, Threat: ""}}, nil
		},
	}
	h := handlers.NewScanHandler(mockSvc, 1024*1024)
	r := gin.Default()
	r.POST("/scan-files", h.ScanFiles)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("files", "good.txt")
	part.Write([]byte("clean"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan-files", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestScanHandler_ScanFiles_Virus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAntivirusService{
		ScanFilesFn: func(ctx context.Context, files []*multipart.FileHeader) ([]interfaces.ScanResult, error) {
			return []interfaces.ScanResult{{FileName: files[0].Filename, Clean: false, Threat: "virus"}}, nil
		},
	}
	h := handlers.NewScanHandler(mockSvc, 1024*1024)
	r := gin.Default()
	r.POST("/scan-files", h.ScanFiles)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("files", "bad.txt")
	part.Write([]byte("virus"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan-files", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestScanHandler_ScanFiles_NoLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAntivirusService{
		ScanFilesFn: func(ctx context.Context, files []*multipart.FileHeader) ([]interfaces.ScanResult, error) {
			return []interfaces.ScanResult{{FileName: files[0].Filename, Clean: true, Threat: ""}}, nil
		},
	}
	h := handlers.NewScanHandler(mockSvc, 0) // no limit
	r := gin.Default()
	r.POST("/scan-files", h.ScanFiles)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("files", "good.txt")
	part.Write([]byte("clean"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan-files", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestScanHandler_ScanFiles_ParseError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockAntivirusService{}
	h := handlers.NewScanHandler(mockSvc, 1024*1024)
	r := gin.Default()
	r.POST("/scan-files", h.ScanFiles)

	// Send invalid multipart data
	body := bytes.NewReader([]byte("invalid multipart"))
	req := httptest.NewRequest(http.MethodPost, "/scan-files", body)
	req.Header.Set("Content-Type", "multipart/form-data; boundary=invalid")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestScanHandler_ScanFile_NoFileUploaded(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := handlers.NewScanHandler(nil, 1024*1024)
	r := gin.Default()
	r.POST("/scan", h.ScanFile)

	// multipart form without the 'file' field
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("other", "value")
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestScanHandler_ScanFile_BodyTooLarge(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// very small limit to trigger body too large
	h := handlers.NewScanHandler(nil, 10)
	r := gin.Default()
	r.POST("/scan", h.ScanFile)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "big.txt")
	// write content larger than limit
	part.Write(bytes.Repeat([]byte("a"), 100))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusRequestEntityTooLarge && w.Code != http.StatusBadRequest {
		t.Fatalf("expected 413 or 400, got %d", w.Code)
	}
}

func TestScanHandler_ScanFiles_NoFilesUploaded(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := handlers.NewScanHandler(nil, 1024*1024)
	r := gin.Default()
	r.POST("/scan-files", h.ScanFiles)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("other", "value")
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan-files", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestScanHandler_ScanFiles_BodyTooLarge(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := handlers.NewScanHandler(nil, 10)
	r := gin.Default()
	r.POST("/scan-files", h.ScanFiles)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("files", "big.txt")
	part.Write(bytes.Repeat([]byte("a"), 100))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/scan-files", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusRequestEntityTooLarge && w.Code != http.StatusBadRequest {
		t.Fatalf("expected 413 or 400, got %d", w.Code)
	}
}
