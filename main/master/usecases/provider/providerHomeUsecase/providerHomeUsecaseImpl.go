package providerHomeUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/provider/providerHomeRepo"
)

type ProviderHomeUsecaseImpl struct {
	providerRepo providerHomeRepo.ProviderHome
}

func InitProviderHomeUsecase(providerRepo providerHomeRepo.ProviderHome) ProviderHome {
	return &ProviderHomeUsecaseImpl{providerRepo: providerRepo}
}
func (pu *ProviderHomeUsecaseImpl) CreateAssetProvider(provider *models.AssetModel) (*models.AssetModel, error) {
	data, err := pu.providerRepo.CreateAssetProvider(provider)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (pu *ProviderHomeUsecaseImpl) GetProviderSaldo(id string) (int, error) {
	data, err := pu.providerRepo.GetProviderSaldo(id)
	if err != nil {
		return 0, err
	}
	return data, nil
}
