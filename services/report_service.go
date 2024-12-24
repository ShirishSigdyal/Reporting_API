package services

import (
	"context"
	"reporting-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GenerateSalesReport(ctx context.Context, filters map[string][]string) (interface{}, error) {
	return s.repo.FetchSalesReport(ctx, filters)
}

func (s *ReportService) GenerateCustomerReport(ctx context.Context, filters map[string][]string) (interface{}, error) {
	return s.repo.FetchCustomerReport(ctx, filters)
}
