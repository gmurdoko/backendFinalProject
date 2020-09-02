package providerUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/providerRepo"
)

type ProviderUsecaseImpl struct {
	providerRepo providerRepo.ProviderRepository
}

func InitProviderUsecase(providerRepo providerRepo.ProviderRepository) ProviderUsecase {
	return &ProviderUsecaseImpl{providerRepo: providerRepo}
}
func (pu *ProviderUsecaseImpl) GetProvider(provider *models.ProviderModel) (bool, error) {
	isValid, err := pu.providerRepo.GetProvider(provider)
	if err != nil {
		return false, err
	}
	return isValid, nil
}
func (pu *ProviderUsecaseImpl) CreateProvider(provider *models.ProviderModel) (*models.ProviderModel, error) {
	data, err := pu.providerRepo.CreateProvider(provider)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (pu *ProviderUsecaseImpl) CreateProviderAsset(provider *models.AssetModel) (*models.AssetModel, error) {
	data, err := pu.providerRepo.CreateAssetProvider(provider)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (pu *ProviderUsecaseImpl) GetProviderSaldo(id string) (int, error) {
	data, err := pu.providerRepo.GetProviderSaldo(id)
	if err != nil {
		return 0, err
	}
	return data, nil
}
