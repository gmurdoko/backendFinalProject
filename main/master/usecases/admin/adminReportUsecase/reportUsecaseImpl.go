package adminReportUsecase

import (
	"finalproject/main/master/models"
	adminassetreportrepo "finalproject/main/master/repositories/admin/report"
	"log"
)

type AdminAssetReportsUsecaseImpl struct {
	adminReportRepo adminassetreportrepo.AdminAssetReportRepo
}

func InitProviderReportUsecaseImpl(adminReportRepo adminassetreportrepo.AdminAssetReportRepo) AdminAssetReportsUsecase {
	// return &ListAssetsUsecaseImpl{listAssetsRepo: listAssetsRepo}
	return &AdminAssetReportsUsecaseImpl{adminReportRepo: adminReportRepo}
}

func (s *AdminAssetReportsUsecaseImpl) GetReportDaily(provId string) ([]*models.ReportAssetDaily, error) {
	reports, err := s.adminReportRepo.GetReportDaily(provId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return reports, nil
}
func (s *AdminAssetReportsUsecaseImpl) GetReportMonthly(provId string) ([]*models.ReportAssetMonthly, error) {
	reports, err := s.adminReportRepo.GetReportMonthly(provId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return reports, nil
}
