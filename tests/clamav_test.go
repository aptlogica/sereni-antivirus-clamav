// Copyright 2026-2030 Aptlogica Technologies Pvt Ltd
// Licensed under the Apache License, Version 2.0
// Websites: https://www.aptlogica.com | https://www.serenibase.com
// Support: support@aptlogica.com | support@serenibase.com

package tests

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/aptlogica/sereni-antivirus-clamav/internal/providers/antivirus/clamav"
	"github.com/aptlogica/sereni-antivirus-clamav/internal/providers/antivirus/interfaces"

	"github.com/dutchcoders/go-clamd"
)

type mockClamd struct {
	PingFn       func() error
	ScanStreamFn func(r io.Reader, abort chan bool) (chan *clamd.ScanResult, error)
}

func (m *mockClamd) Ping() error {
	return m.PingFn()
}

func (m *mockClamd) ScanStream(r io.Reader, abort chan bool) (chan *clamd.ScanResult, error) {
	return m.ScanStreamFn(r, abort)
}

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		config    clamav.Config
		expectErr bool
	}{
		{
			name:      "valid config",
			config:    clamav.Config{Address: "127.0.0.1:3310", TimeoutSeconds: 30},
			expectErr: false,
		},
		{
			name:      "missing address",
			config:    clamav.Config{TimeoutSeconds: 30},
			expectErr: true,
		},
		{
			name:      "zero timeout",
			config:    clamav.Config{Address: "127.0.0.1:3310", TimeoutSeconds: 0},
			expectErr: false, // should set default
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := clamav.New(tt.config)
			if tt.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestClamavProvider_Ping(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *mockClamd)
		expectErr bool
	}{
		{
			name: "success",
			mockSetup: func(m *mockClamd) {
				m.PingFn = func() error { return nil }
			},
		},
		{
			name: "error",
			mockSetup: func(m *mockClamd) {
				m.PingFn = func() error { return errors.New("ping failed") }
			},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockClamd{}
			tt.mockSetup(m)
			p := clamav.NewWithClient(clamav.Config{Address: "test"}, m)
			err := p.Ping(nil) // context not used in Ping
			if tt.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestClamavProvider_ScanReader(t *testing.T) {
	tests := []struct {
		name      string
		fileName  string
		input     io.Reader
		mockSetup func(m *mockClamd)
		expected  interfaces.ScanResult
		expectErr bool
	}{
		{
			name:     "success - clean",
			fileName: "file.txt",
			input:    strings.NewReader("clean"),
			mockSetup: func(m *mockClamd) {
				m.ScanStreamFn = func(r io.Reader, abort chan bool) (chan *clamd.ScanResult, error) {
					ch := make(chan *clamd.ScanResult, 1)
					ch <- &clamd.ScanResult{Status: clamd.RES_OK, Description: ""}
					close(ch)
					return ch, nil
				}
			},
			expected: interfaces.ScanResult{FileName: "file.txt", Clean: true, Threat: ""},
		},
		{
			name:     "success - virus found",
			fileName: "bad.txt",
			input:    strings.NewReader("bad"),
			mockSetup: func(m *mockClamd) {
				m.ScanStreamFn = func(r io.Reader, abort chan bool) (chan *clamd.ScanResult, error) {
					ch := make(chan *clamd.ScanResult, 1)
					ch <- &clamd.ScanResult{Status: clamd.RES_FOUND, Description: "EICAR-Test-File"}
					close(ch)
					return ch, nil
				}
			},
			expected: interfaces.ScanResult{FileName: "bad.txt", Clean: false, Threat: "EICAR-Test-File"},
		},
		{
			name:     "error - scan error",
			fileName: "err.txt",
			input:    strings.NewReader("err"),
			mockSetup: func(m *mockClamd) {
				m.ScanStreamFn = func(r io.Reader, abort chan bool) (chan *clamd.ScanResult, error) {
					return nil, errors.New("scan error")
				}
			},
			expected:  interfaces.ScanResult{FileName: "err.txt", Clean: false, Threat: "scan error"},
			expectErr: true,
		},
		{
			name:     "error - parse error",
			fileName: "parse.txt",
			input:    strings.NewReader("parse"),
			mockSetup: func(m *mockClamd) {
				m.ScanStreamFn = func(r io.Reader, abort chan bool) (chan *clamd.ScanResult, error) {
					ch := make(chan *clamd.ScanResult, 1)
					ch <- &clamd.ScanResult{Status: clamd.RES_PARSE_ERROR, Description: "parse error"}
					close(ch)
					return ch, nil
				}
			},
			expected:  interfaces.ScanResult{FileName: "parse.txt", Clean: false, Threat: "parse error"},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockClamd{}
			tt.mockSetup(m)
			p := clamav.NewWithClient(clamav.Config{Address: "test"}, m)
			got, err := p.ScanReader(nil, tt.fileName, tt.input) // context not used
			if tt.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Fatalf("expected %+v, got %+v", tt.expected, got)
			}
		})
	}
}
