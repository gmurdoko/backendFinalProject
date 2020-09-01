package providerassetreportrepo

import "finalproject/main/master/models"

type AssetReportRepo interface {
	GetReportDaily(string) []*models.ReportAssetDaily
	GetReportMonthly(string) []*models.ReportAssetMonthly
}
