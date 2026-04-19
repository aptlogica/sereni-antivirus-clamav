// Copyright 2026-2030 Aptlogica Technologies Pvt Ltd
// Licensed under the Apache License, Version 2.0
// Websites: https://www.aptlogica.com | https://www.serenibase.com
// Support: support@aptlogica.com | support@serenibase.com

package tests

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/aptlogica/sereni-antivirus-clamav/internal/providers/antivirus/interfaces"
)

type mockProviderInterface struct {
	PingFn       func(ctx context.Context) error
	ScanReaderFn func(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error)
}

func (m *mockProviderInterface) Ping(ctx context.Context) error {
	return m.PingFn(ctx)
}

func (m *mockProviderInterface) ScanReader(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
	return m.ScanReaderFn(ctx, fileName, r)
}

func TestInterfaceProvider_Ping(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *mockProviderInterface)
		expectErr bool
	}{
		{
			name: "success",
			mockSetup: func(m *mockProviderInterface) {
				m.PingFn = func(ctx context.Context) error { return nil }
			},
		},
		{
			name: "error",
			mockSetup: func(m *mockProviderInterface) {
				m.PingFn = func(ctx context.Context) error { return errors.New("ping failed") }
			},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockProviderInterface{}
			tt.mockSetup(m)
			err := m.Ping(context.Background())
			if tt.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestInterfaceProvider_ScanReader(t *testing.T) {
	tests := []struct {
		name      string
		fileName  string
		input     io.Reader
		mockSetup func(m *mockProviderInterface)
		expected  interfaces.ScanResult
		expectErr bool
	}{
		{
			name:     "success - clean",
			fileName: "file.txt",
			input:    nil,
			mockSetup: func(m *mockProviderInterface) {
				m.ScanReaderFn = func(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
					return interfaces.ScanResult{FileName: fileName, Clean: true, Threat: ""}, nil
				}
			},
			expected: interfaces.ScanResult{FileName: "file.txt", Clean: true, Threat: ""},
		},
		{
			name:     "error - scan fails",
			fileName: "bad.txt",
			input:    nil,
			mockSetup: func(m *mockProviderInterface) {
				m.ScanReaderFn = func(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
					return interfaces.ScanResult{FileName: fileName, Clean: false, Threat: "virus"}, errors.New("scan error")
				}
			},
			expected:  interfaces.ScanResult{FileName: "bad.txt", Clean: false, Threat: "virus"},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockProviderInterface{}
			tt.mockSetup(m)
			res, err := m.ScanReader(context.Background(), tt.fileName, tt.input)
			if tt.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if res != tt.expected {
				t.Fatalf("expected %+v, got %+v", tt.expected, res)
			}
		})
	}
}
