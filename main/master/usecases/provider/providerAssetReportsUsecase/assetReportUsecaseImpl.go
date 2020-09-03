package providerAssetReportsUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/provider/providerAssetReportsRepo"
	"log"
)

type ProviderAssetReportsUsecaseImpl struct {
	providerReportRepo providerAssetReportsRepo.ProviderAssetReportRepo
}

func InitProviderReportUsecaseImpl(providerReportRepo providerAssetReportsRepo.ProviderAssetReportRepo) ProviderAssetReportsUsecase {
	// return &ListAssetsUsecaseImpl{listAssetsRepo: listAssetsRepo}
	return &ProviderAssetReportsUsecaseImpl{providerReportRepo: providerReportRepo}
}

func (s *ProviderAssetReportsUsecaseImpl) GetReportDaily(provId string) ([]*models.ReportAssetDaily, error) {
	reports, err := s.providerReportRepo.GetReportDaily(provId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return reports, nil
}
func (s *ProviderAssetReportsUsecaseImpl) GetReportMonthly(provId string) ([]*models.ReportAssetMonthly, error) {
	reports, err := s.providerReportRepo.GetReportMonthly(provId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return reports, nil
}
