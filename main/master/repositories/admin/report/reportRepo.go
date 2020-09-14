package adminassetsreportsrepo

import "finalproject/main/master/models"

type AdminAssetReportRepo interface {
	GetReportDaily(string, string, string) ([]*models.ReportAssetDaily, error)
	GetReportMonthly(string) ([]*models.ReportAssetMonthly, error)
}
