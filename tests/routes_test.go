package tests

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"sereni-antivirus/internal/handlers"
	"sereni-antivirus/internal/providers/antivirus/interfaces"
	"sereni-antivirus/internal/routes"

	"github.com/gin-gonic/gin"
)

// mockService implements the AntivirusService interface for testing.
type mockService struct{}

func (m *mockService) ScanFile(ctx context.Context, filename string, file io.Reader) (interfaces.ScanResult, error) {
	return interfaces.ScanResult{Clean: true}, nil
}

func (m *mockService) ScanFiles(ctx context.Context, files []*multipart.FileHeader) ([]interfaces.ScanResult, error) {
	return []interfaces.ScanResult{{Clean: true}}, nil
}

func TestSetupRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := handlers.NewScanHandler(&mockService{}, 1024*1024)
	router := routes.SetupRouter(handler)

	cases := []struct {
		name   string
		method string
		path   string
		want   int
	}{
		{"scan file POST without file", "POST", "/scan", http.StatusBadRequest},
		{"scan files POST without files", "POST", "/scan-files", http.StatusBadRequest},
		{"scan file GET", "GET", "/scan", http.StatusNotFound},
		{"unknown path", "POST", "/unknown", http.StatusNotFound},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			if w.Code != tc.want {
				t.Errorf("%s: want %d, got %d", tc.name, tc.want, w.Code)
			}
		})
	}
}
