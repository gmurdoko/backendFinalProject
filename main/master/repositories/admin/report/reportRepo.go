package adminassetsreportsrepo

import "finalproject/main/master/models"

type AdminAssetReportRepo interface {
	GetReportDaily(string) []*models.ReportAssetDaily
	GetReportMonthly(string) []*models.ReportAssetMonthly
}
