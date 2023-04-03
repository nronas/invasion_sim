package repositories

import (
	"context"
	"log"
)

var _ ReportRepository = (*stdoutReportRepositoryImpl)(nil)

type stdoutReportRepositoryImpl struct{}

func NewStdoutReportRepository() *stdoutReportRepositoryImpl {
	return &stdoutReportRepositoryImpl{}
}

func (*stdoutReportRepositoryImpl) Report(_ context.Context, msg string) error {
	log.Println(msg)
	return nil
}
