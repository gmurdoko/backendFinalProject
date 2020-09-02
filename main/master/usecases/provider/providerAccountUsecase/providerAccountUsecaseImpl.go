package providerAccountUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/provider/providerAccountRepo"
)

type ProviderUsecaseAccImpl struct {
	providerRepo providerAccountRepo.ProviderAccount
}

func InitProviderAccUsecase(providerRepo providerAccountRepo.ProviderAccount) ProviderAccount {
	return &ProviderUsecaseAccImpl{providerRepo: providerRepo}
}
func (pu *ProviderUsecaseAccImpl) GetProvider(provider *models.ProviderModel) (*models.ProviderModel, bool, error) {
	data, isValid, err := pu.providerRepo.GetProvider(provider)
	if err != nil {
		return nil, false, err
	}
	return data, isValid, nil
}
func (pu *ProviderUsecaseAccImpl) CreateProvider(provider *models.ProviderModel) (*models.ProviderModel, error) {
	data, err := pu.providerRepo.CreateProvider(provider)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (pu *ProviderUsecaseAccImpl) GetProviderById(id string) (*models.ProviderModel, error) {
	data, err := pu.providerRepo.GetProviderById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
