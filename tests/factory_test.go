// Copyright (c) 2026 Aptlogica Technologies Private Limited
// SPDX-License-Identifier: MIT
// Websites: https://www.aptlogica.com | https://www.serenibase.com
// Support: support@aptlogica.com | support@serenibase.com

package tests

import (
	"sereni-antivirus/internal/config"
	"sereni-antivirus/internal/providers/antivirus"
	"testing"
)

type mockClamAVProvider struct{}

func (m *mockClamAVProvider) Ping(ctx interface{}) error { return nil }
func (m *mockClamAVProvider) ScanReader(ctx interface{}, fileName string, r interface{}) (interface{}, error) {
	return nil, nil
}

func TestNewAntivirus(t *testing.T) {
	tests := []struct {
		name      string
		cfg       *config.AntivirusConfig
		expectErr bool
	}{
		{
			name: "clamav driver",
			cfg: &config.AntivirusConfig{
				Driver: "clamav",
				ClamAV: config.ClamAVConfig{
					Address:        "127.0.0.1:3310",
					TimeoutSeconds: 10,
				},
			},
		},
		{
			name:      "unsupported driver",
			cfg:       &config.AntivirusConfig{Driver: "unknown"},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := antivirus.NewAntivirus(tt.cfg)
			if tt.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
