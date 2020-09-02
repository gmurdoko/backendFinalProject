package providerassetreportusecase

import "finalproject/main/master/models"

type ProviderAssetReportsUsecase interface {
	GetReportDaily(string) ([]*models.ReportAssetDaily, error)
	GetReportMonthly(string) ([]*models.ReportAssetMonthly, error)
}
