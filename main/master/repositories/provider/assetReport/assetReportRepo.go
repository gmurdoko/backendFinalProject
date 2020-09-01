package providerassetreportrepo

import "finalproject/main/master/models"

type ProviderAssetReportRepo interface {
	GetReportDaily(string) []*models.ReportAssetDaily
	GetReportMonthly(string) []*models.ReportAssetMonthly
}
