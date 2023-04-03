package services

import (
	"context"

	"github.com/nronas/invasion_sim/internal/repositories"
)

type reportService struct {
	reportRepository repositories.ReportRepository
}

// NewReportService creates a new service that uses the underlying repository to report messages.
func NewReportService(reportRepository repositories.ReportRepository) *reportService {
	return &reportService{reportRepository: reportRepository}
}

// Report reports a message through the given repository. See repositories.NewStdoutReportRepository for example.
func (rs *reportService) Report(ctx context.Context, msg string) error {
	return rs.reportRepository.Report(ctx, msg)
}
