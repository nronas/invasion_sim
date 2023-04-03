package services

import (
	"context"

	"github.com/nronas/invasion_sim/internal/repositories"
)

type reportService struct {
	reportRepository repositories.ReportRepository
}

func NewReportService(reportRepository repositories.ReportRepository) *reportService {
	return &reportService{reportRepository: reportRepository}
}

func (rs *reportService) Report(ctx context.Context, msg string) error {
	return rs.reportRepository.Report(ctx, msg)
}
