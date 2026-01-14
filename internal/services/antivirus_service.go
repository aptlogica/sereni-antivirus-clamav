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

type AntivirusService interface {
	ScanFile(ctx context.Context, name string, file io.Reader) (interfaces.ScanResult, error)
	ScanFiles(ctx context.Context, files []*multipart.FileHeader) ([]interfaces.ScanResult, error)
}

type antivirusService struct {
	provider interfaces.Provider
}

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
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobCh {
				file, err := job.header.Open()
				if err != nil {
					outcomeCh <- scanOutcome{
						idx: job.idx,
						result: interfaces.ScanResult{
							FileName: job.header.Filename,
							Clean:    false,
							Threat:   "Failed to open file: " + err.Error(),
						},
						err: fmt.Errorf("%s: %w", job.header.Filename, err),
					}
					continue
				}

				res, err := s.provider.ScanReader(ctx, job.header.Filename, file)
				file.Close()

				if err != nil && res.Threat == "" {
					res.Threat = "scan failed: " + err.Error()
				}
				if err != nil {
					err = fmt.Errorf("%s: %w", job.header.Filename, err)
				}

				outcomeCh <- scanOutcome{idx: job.idx, result: res, err: err}
			}
		}()
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
		if outcome.err != nil {
			errs = append(errs, outcome.err)
		}
		results[outcome.idx] = outcome.result
	}

	return results, errors.Join(errs...)
}
