package providerassetreportrepo

import "finalproject/main/master/models"

type ProviderAssetReportRepo interface {
	GetReportDaily(string) ([]*models.ReportAssetDaily, error)
	GetReportMonthly(string) ([]*models.ReportAssetMonthly, error)
}
