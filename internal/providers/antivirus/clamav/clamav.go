// Copyright (c) 2026 Aptlogica Technologies Private Limited
// Licensed under the Apache License, Version 2.0
// Websites: https://www.aptlogica.com | https://www.serenibase.com
// Support: support@aptlogica.com | support@serenibase.com

// Package clamav provides antivirus scanning using ClamAV (clamd).
package clamav

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/aptlogica/sereni-antivirus-clamav/internal/providers/antivirus/interfaces"

	"github.com/dutchcoders/go-clamd"
)

// ClamdClient interface for testing.
type ClamdClient interface {
	Ping() error
	ScanStream(r io.Reader, abort chan bool) (chan *clamd.ScanResult, error)
}

// Config holds the configuration for the ClamAV provider.
type Config struct {
	// Address is the host:port of the clamd service, e.g. 127.0.0.1:3310.
	Address string
	// TimeoutSeconds is the network timeout for clamd operations.
	TimeoutSeconds int
}

// Provider implements interfaces.Provider for ClamAV (clamd).
type Provider struct {
	config Config
	clamd  ClamdClient
}

// New creates a new ClamAV antivirus provider instance.
func New(cfg Config) (*Provider, error) {
	if cfg.Address == "" {
		return nil, errors.New("clamav: address is required")
	}
	if cfg.TimeoutSeconds <= 0 {
		cfg.TimeoutSeconds = 30
	}
	// Use "tcp" for ClamdTCP, or "unix" for ClamdUnix
	c := clamd.NewClamd("tcp://" + cfg.Address)
	return &Provider{config: cfg, clamd: c}, nil
}

// NewWithClient creates a new ClamAV provider with a custom client (for testing).
func NewWithClient(cfg Config, client ClamdClient) *Provider {
	return &Provider{config: cfg, clamd: client}
}

// Ping verifies clamd availability by sending a PING command.
func (p *Provider) Ping(ctx context.Context) error {
	return p.clamd.Ping()
}

// ScanReader scans a stream for malware using clamd INSTREAM.
func (p *Provider) ScanReader(ctx context.Context, fileName string, r io.Reader) (interfaces.ScanResult, error) {
	abort := make(chan bool)
	defer close(abort)

	resultChan, err := p.clamd.ScanStream(r, abort)
	if err != nil {
		return interfaces.ScanResult{
			FileName: fileName,
			Clean:    false,
			Threat:   err.Error(),
		}, errors.New("clamav scan failed for file " + fileName + ": " + err.Error())
	}

	for scanResult := range resultChan {
		switch scanResult.Status {
		case clamd.RES_OK:
			// clean
			return interfaces.ScanResult{
				FileName: fileName,
				Clean:    true,
				Threat:   "",
			}, nil
		case clamd.RES_FOUND:
			return interfaces.ScanResult{
				FileName: fileName,
				Clean:    false,
				Threat:   scanResult.Description,
			}, nil
		case clamd.RES_ERROR, clamd.RES_PARSE_ERROR:
			return interfaces.ScanResult{
				FileName: fileName,
				Clean:    false,
				Threat:   scanResult.Description,
			}, fmt.Errorf("clamav scan error on %s: %s", fileName, scanResult.Description)
		}
	}
	// If no result, something went wrong
	return interfaces.ScanResult{
		FileName: fileName,
		Clean:    false,
		Threat:   "no scan result returned",
	}, errors.New("clamav: no scan result returned ")
}
