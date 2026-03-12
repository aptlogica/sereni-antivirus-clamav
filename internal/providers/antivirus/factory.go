// Copyright (c) 2026 Aptlogica Technologies Private Limited
// SPDX-License-Identifier: MIT
// Websites: https://www.aptlogica.com | https://www.serenibase.com
// Support: support@aptlogica.com | support@serenibase.com

// Package antivirus provides factory functions for creating antivirus providers.
package antivirus

import (
	"fmt"
	"strings"

	"sereni-antivirus/internal/config"
	"sereni-antivirus/internal/providers/antivirus/clamav"
	"sereni-antivirus/internal/providers/antivirus/interfaces"
)

// NewAntivirus constructs an antivirus provider based on configuration.
func NewAntivirus(cfg *config.AntivirusConfig) (interfaces.Provider, error) {
	switch strings.ToLower(cfg.Driver) {
	case "clamav":
		return clamav.New(clamav.Config{
			Address:        cfg.ClamAV.Address,
			TimeoutSeconds: cfg.ClamAV.TimeoutSeconds,
		})
	default:
		return nil, fmt.Errorf("unsupported antivirus driver: %s", cfg.Driver)
	}
}
