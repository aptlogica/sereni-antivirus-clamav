// Copyright (c) 2026 Aptlogica Technologies Private Limited
// SPDX-License-Identifier: MIT
// Websites: https://www.aptlogica.com | https://www.serenibase.com
// Support: support@aptlogica.com | support@serenibase.com

package tests

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"testing"

	"sereni-antivirus/internal/providers/antivirus/interfaces"
	"sereni-antivirus/internal/services"
)

type mockProvider struct {
	ScanReaderFn func(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error)
}

func (m *mockProvider) ScanReader(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
	return m.ScanReaderFn(ctx, fileName, r)
}

func (m *mockProvider) Ping(ctx context.Context) error {
	return nil
}

// inMemoryMultipartFile implements multipart.File for testing
type inMemoryMultipartFile struct {
	*bytes.Reader
}

func (f *inMemoryMultipartFile) Close() error               { return nil }
func (f *inMemoryMultipartFile) Read(p []byte) (int, error) { return f.Reader.Read(p) }
func (f *inMemoryMultipartFile) Seek(offset int64, whence int) (int64, error) {
	return f.Reader.Seek(offset, whence)
}
func (f *inMemoryMultipartFile) ReadAt(p []byte, off int64) (int, error) {
	cur, _ := f.Reader.Seek(0, io.SeekCurrent)
	f.Reader.Seek(off, io.SeekStart)
	n, err := f.Reader.Read(p)
	f.Reader.Seek(cur, io.SeekStart)
	return n, err
}

func newMultipartFileHeader(t *testing.T, name string, content []byte) *multipart.FileHeader {
	buf := &bytes.Buffer{}
	mw := multipart.NewWriter(buf)
	part, err := mw.CreateFormFile("file", name)
	if err != nil {
		t.Fatalf("failed to create form file: %v", err)
	}
	part.Write(content)
	mw.Close()

	req := &http.Request{Header: make(http.Header)}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))

	mr, err := req.MultipartReader()
	if err != nil {
		t.Fatalf("failed to get multipart reader: %v", err)
	}
	form, err := mr.ReadForm(10 << 20) // 10MB
	if err != nil {
		t.Fatalf("failed to read multipart form: %v", err)
	}
	files := form.File["file"]
	if len(files) == 0 {
		t.Fatalf("no file found in multipart form")
	}
	return files[0]
}

func TestAntivirusService_ScanFile(t *testing.T) {
	tests := []struct {
		name        string
		fileName    string
		fileData    []byte
		mockSetup   func(m *mockProvider)
		expected    interfaces.ScanResult
		expectError bool
	}{
		{
			name:     "success - clean file",
			fileName: "clean.txt",
			fileData: []byte("clean content"),
			mockSetup: func(m *mockProvider) {
				m.ScanReaderFn = func(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
					return interfaces.ScanResult{FileName: fileName, Clean: true, Threat: ""}, nil
				}
			},
			expected: interfaces.ScanResult{FileName: "clean.txt", Clean: true, Threat: ""},
		},
		{
			name:     "error - scan fails",
			fileName: "fail.txt",
			fileData: []byte("bad content"),
			mockSetup: func(m *mockProvider) {
				m.ScanReaderFn = func(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
					return interfaces.ScanResult{FileName: fileName, Clean: false, Threat: "scan failed"}, errors.New("scan error")
				}
			},
			expected:    interfaces.ScanResult{FileName: "fail.txt", Clean: false, Threat: "scan failed"},
			expectError: true,
		},
		{
			name:     "nil file reader",
			fileName: "nil.txt",
			fileData: nil,
			mockSetup: func(m *mockProvider) {
				m.ScanReaderFn = func(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
					if r == nil {
						return interfaces.ScanResult{}, errors.New("nil reader")
					}
					return interfaces.ScanResult{FileName: fileName, Clean: true, Threat: ""}, nil
				}
			},
			expected:    interfaces.ScanResult{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockProvider{}
			tt.mockSetup(mock)
			svc := services.NewAntivirusService(mock)
			var reader io.Reader
			if tt.fileData != nil {
				reader = bytes.NewReader(tt.fileData)
			}
			result, err := svc.ScanFile(context.Background(), tt.fileName, reader)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Fatalf("expected %+v, got %+v", tt.expected, result)
			}
		})
	}
}

func TestAntivirusService_ScanFiles(t *testing.T) {
	tests := []struct {
		name        string
		files       []*multipart.FileHeader
		mockSetup   func(m *mockProvider)
		expected    []interfaces.ScanResult
		expectError bool
	}{
		{
			name: "success - multiple clean files",
			files: []*multipart.FileHeader{
				newMultipartFileHeader(t, "a.txt", []byte("a")),
				newMultipartFileHeader(t, "b.txt", []byte("b")),
			},
			mockSetup: func(m *mockProvider) {
				m.ScanReaderFn = func(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
					return interfaces.ScanResult{FileName: fileName, Clean: true, Threat: ""}, nil
				}
			},
			expected: []interfaces.ScanResult{
				{FileName: "a.txt", Clean: true, Threat: ""},
				{FileName: "b.txt", Clean: true, Threat: ""},
			},
		},
		{
			name: "error - one file fails to open",
			files: []*multipart.FileHeader{
				newMultipartFileHeader(t, "fail.txt", []byte("")),
			},
			mockSetup: func(m *mockProvider) {
				m.ScanReaderFn = func(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
					if fileName == "fail.txt" {
						return interfaces.ScanResult{FileName: fileName, Clean: false, Threat: "Failed to open file: open error"}, fmt.Errorf("fail.txt: open error")
					}
					return interfaces.ScanResult{}, nil
				}
			},
			expected: []interfaces.ScanResult{
				{FileName: "fail.txt", Clean: false, Threat: "Failed to open file: open error"},
			},
			expectError: true,
		},
		{
			name:      "empty input",
			files:     []*multipart.FileHeader{},
			mockSetup: func(m *mockProvider) {},
			expected:  nil,
		},
		{
			name: "scan error propagates",
			files: []*multipart.FileHeader{
				newMultipartFileHeader(t, "bad.txt", []byte("bad")),
			},
			mockSetup: func(m *mockProvider) {
				m.ScanReaderFn = func(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
					return interfaces.ScanResult{FileName: fileName, Clean: false, Threat: "virus"}, errors.New("virus found")
				}
			},
			expected: []interfaces.ScanResult{
				{FileName: "bad.txt", Clean: false, Threat: "virus"},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockProvider{}
			tt.mockSetup(mock)
			svc := services.NewAntivirusService(mock)
			results, err := svc.ScanFiles(context.Background(), tt.files)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if fmt.Sprintf("%+v", results) != fmt.Sprintf("%+v", tt.expected) {
				t.Fatalf("expected %+v, got %+v", tt.expected, results)
			}
		})
	}
}
