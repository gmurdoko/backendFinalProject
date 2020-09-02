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
func (pu *ProviderUsecaseAccImpl) GetProvider(provider *models.Providers) (bool, error) {
	isValid, err := pu.providerRepo.GetProvider(provider)
	if err != nil {
		return false, err
	}
	return isValid, nil
}
func (pu *ProviderUsecaseAccImpl) CreateProvider(provider *models.Providers) (*models.Providers, error) {
	data, err := pu.providerRepo.CreateProvider(provider)
	if err != nil {
		return nil, err
	}
	return data, nil
}
