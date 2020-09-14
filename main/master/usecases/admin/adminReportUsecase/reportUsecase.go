package adminReportUsecase

import "finalproject/main/master/models"

type AdminAssetReportsUsecase interface {
	GetReportDaily(string, string, string) ([]*models.ReportAssetDaily, error)
	GetReportMonthly(string) ([]*models.ReportAssetMonthly, error)
}
