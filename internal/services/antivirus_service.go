package services

import (
	"context"
	"io"
	"mime/multipart"
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
	var results []interfaces.ScanResult
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			// Append error result for this file
			results = append(results, interfaces.ScanResult{
				FileName: fileHeader.Filename,
				Clean:    false,
				Threat:   "Failed to open file: " + err.Error(),
			})
			continue
		}

		res, err := s.provider.ScanReader(ctx, fileHeader.Filename, file)
		file.Close()

		if err != nil && (res.Threat == "") {
			res.Threat = err.Error()
		}
		results = append(results, res)
	}
	return results, nil
}
