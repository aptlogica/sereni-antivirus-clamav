// Copyright (c) 2026 Aptlogica Technologies Private Limited
// SPDX-License-Identifier: MIT
// Websites: https://www.aptlogica.com | https://www.serenibase.com
// Support: support@aptlogica.com | support@serenibase.com

package tests

import (
	"os"
	"sereni-antivirus/internal/config"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")

	if got := config.GetEnv("TEST_KEY", "default"); got != "test_value" {
		t.Fatalf("expected test_value, got %s", got)
	}
	if got := config.GetEnv("MISSING_KEY", "default"); got != "default" {
		t.Fatalf("expected default, got %s", got)
	}
}

func TestLoad_Config_Defaults(t *testing.T) {
	os.Clearenv()
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Port != "8084" {
		t.Errorf("expected default port 8084, got %s", cfg.Port)
	}
	if cfg.Antivirus.Driver != "clamav" {
		t.Errorf("expected default antivirus driver clamav, got %s", cfg.Antivirus.Driver)
	}
	if cfg.Antivirus.ClamAV.Address != "127.0.0.1:3310" {
		t.Errorf("expected default clamav address, got %s", cfg.Antivirus.ClamAV.Address)
	}
	if cfg.Antivirus.ClamAV.TimeoutSeconds != 30 {
		t.Errorf("expected default clamav timeout 30, got %d", cfg.Antivirus.ClamAV.TimeoutSeconds)
	}
	if cfg.MaxUploadSizeBytes != 32*1024*1024 {
		t.Errorf("expected default max upload size 32MB, got %d", cfg.MaxUploadSizeBytes)
	}
}

func TestLoad_Config_EnvOverride(t *testing.T) {
	os.Setenv("PORT", "9999")
	os.Setenv("ANTIVIRUS_DRIVER", "custom")
	os.Setenv("CLAMAV_ADDRESS", "10.0.0.1:1234")
	os.Setenv("CLAMAV_TIMEOUT_SECONDS", "99")
	os.Setenv("MAX_UPLOAD_SIZE_MB", "64")
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("ANTIVIRUS_DRIVER")
		os.Unsetenv("CLAMAV_ADDRESS")
		os.Unsetenv("CLAMAV_TIMEOUT_SECONDS")
		os.Unsetenv("MAX_UPLOAD_SIZE_MB")
	}()
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Port != "9999" {
		t.Errorf("expected port 9999, got %s", cfg.Port)
	}
	if cfg.Antivirus.Driver != "custom" {
		t.Errorf("expected antivirus driver custom, got %s", cfg.Antivirus.Driver)
	}
	if cfg.Antivirus.ClamAV.Address != "10.0.0.1:1234" {
		t.Errorf("expected clamav address, got %s", cfg.Antivirus.ClamAV.Address)
	}
	if cfg.Antivirus.ClamAV.TimeoutSeconds != 99 {
		t.Errorf("expected clamav timeout 99, got %d", cfg.Antivirus.ClamAV.TimeoutSeconds)
	}
	if cfg.MaxUploadSizeBytes != 64*1024*1024 {
		t.Errorf("expected max upload size 64MB, got %d", cfg.MaxUploadSizeBytes)
	}
}

func TestGetEnvAsInt(t *testing.T) {
	os.Setenv("TEST_INT", "42")
	defer os.Unsetenv("TEST_INT")

	if got := config.GetEnvAsInt("TEST_INT", 0); got != 42 {
		t.Fatalf("expected 42, got %d", got)
	}
	if got := config.GetEnvAsInt("MISSING_INT", 7); got != 7 {
		t.Fatalf("expected 7, got %d", got)
	}
	os.Setenv("TEST_INT_INVALID", "notanint")
	if got := config.GetEnvAsInt("TEST_INT_INVALID", 9); got != 9 {
		t.Fatalf("expected 9, got %d", got)
	}
	os.Unsetenv("TEST_INT_INVALID")
}
