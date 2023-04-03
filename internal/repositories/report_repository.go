package repositories

import "context"

type ReportRepository interface {
	Report(ctx context.Context, msg string) error
}
