// Copyright (c) 2026 Aptlogica Technologies Private Limited
// SPDX-License-Identifier: MIT
// Websites: https://www.aptlogica.com | https://www.serenibase.com
// Support: support@aptlogica.com | support@serenibase.com

// Package services provides business logic for antivirus operations.
package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"runtime"
	"sync"

	"sereni-antivirus/internal/providers/antivirus/interfaces"
)

// AntivirusService defines methods for scanning files.
type AntivirusService interface {
	ScanFile(ctx context.Context, name string, file io.Reader) (interfaces.ScanResult, error)
	ScanFiles(ctx context.Context, files []*multipart.FileHeader) ([]interfaces.ScanResult, error)
}

type antivirusService struct {
	provider interfaces.Provider
}

// NewAntivirusService creates a new antivirus service with the given provider.
func NewAntivirusService(provider interfaces.Provider) AntivirusService {
	return &antivirusService{
		provider: provider,
	}
}

func (s *antivirusService) ScanFile(ctx context.Context, name string, file io.Reader) (interfaces.ScanResult, error) {
	return s.provider.ScanReader(ctx, name, file)
}

func (s *antivirusService) ScanFiles(ctx context.Context, files []*multipart.FileHeader) ([]interfaces.ScanResult, error) {
	total := len(files)
	if total == 0 {
		return nil, nil
	}

	results := make([]interfaces.ScanResult, total)
	var errs []error

	type scanJob struct {
		idx    int
		header *multipart.FileHeader
	}
	type scanOutcome struct {
		idx    int
		result interfaces.ScanResult
		err    error
	}

	jobCh := make(chan scanJob, total)
	outcomeCh := make(chan scanOutcome, total)

	workerCount := runtime.NumCPU()
	if workerCount <= 0 {
		workerCount = 1
	}
	if workerCount > total {
		workerCount = total
	}

	var wg sync.WaitGroup
	worker := func() {
		defer wg.Done()
		for job := range jobCh {
			outcomeCh <- s.handleScanJob(ctx, job)
		}
	}
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker()
	}

	go func() {
		wg.Wait()
		close(outcomeCh)
	}()

	for idx, header := range files {
		jobCh <- scanJob{idx: idx, header: header}
	}
	close(jobCh)

	for outcome := range outcomeCh {
		s.handleScanOutcome(&results, &errs, outcome)
	}

	return results, errors.Join(errs...)
}

// handleScanJob processes a scanJob and returns a scanOutcome
func (s *antivirusService) handleScanJob(ctx context.Context, job struct {
	idx    int
	header *multipart.FileHeader
}) struct {
	idx    int
	result interfaces.ScanResult
	err    error
} {
	var outcome struct {
		idx    int
		result interfaces.ScanResult
		err    error
	}
	outcome.idx = job.idx
	file, err := job.header.Open()
	if err != nil {
		outcome.result = interfaces.ScanResult{
			FileName: job.header.Filename,
			Clean:    false,
			Threat:   "Failed to open file: " + err.Error(),
		}
		outcome.err = fmt.Errorf("%s: %w", job.header.Filename, err)
		return outcome
	}
	res, err := s.provider.ScanReader(ctx, job.header.Filename, file)
	file.Close()
	if err != nil && res.Threat == "" {
		res.Threat = "scan failed: " + err.Error()
	}
	if err != nil {
		err = fmt.Errorf("%s: %w", job.header.Filename, err)
	}
	outcome.result = res
	outcome.err = err
	return outcome
}

// handleScanOutcome updates results and errs based on scanOutcome
func (s *antivirusService) handleScanOutcome(results *[]interfaces.ScanResult, errs *[]error, outcome struct {
	idx    int
	result interfaces.ScanResult
	err    error
}) {
	if outcome.err != nil {
		*errs = append(*errs, outcome.err)
	}
	(*results)[outcome.idx] = outcome.result
}
