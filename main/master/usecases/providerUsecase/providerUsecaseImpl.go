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
func (pu *ProviderUsecaseImpl) CreateProvider(provider *models.ProviderModel) (*models.ProviderModel, error) {
	data, err := pu.providerRepo.CreateProvider(provider)
	if err != nil {
		return nil, err
	}
	return data, nil
}
