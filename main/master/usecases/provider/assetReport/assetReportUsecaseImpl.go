package providerassetreportusecase

import (
	"finalproject/main/master/models"
	providerassetreportrepo "finalproject/main/master/repositories/provider/assetReport"
	"log"
)

type ProviderAssetReportsUsecaseImpl struct {
	providerReportRepo providerassetreportrepo.ProviderAssetReportRepo
}

func InitProviderReportUsecaseImpl(providerReportRepo providerassetreportrepo.ProviderAssetReportRepo) ProviderAssetReportsUsecase {
	// return &ListAssetsUsecaseImpl{listAssetsRepo: listAssetsRepo}
	return &ProviderAssetReportsUsecaseImpl{providerReportRepo: providerReportRepo}
}

func (s *ProviderAssetReportsUsecaseImpl) GetReportDaily(provId string) ([]*models.ReportAssetDaily, error) {
	reports, err := s.GetReportDaily(provId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return reports, nil
}
func (s *ProviderAssetReportsUsecaseImpl) GetReportMonthly(provId string) ([]*models.ReportAssetMonthly, error) {
	reports, err := s.GetReportMonthly(provId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return reports, nil
}
