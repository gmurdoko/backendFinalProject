package providerAccountUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/provider/providerAccountRepo"
	"fmt"
)

type ProviderUsecaseAccImpl struct {
	providerRepo providerAccountRepo.ProviderAccount
}

func InitProviderAccUsecase(providerRepo providerAccountRepo.ProviderAccount) ProviderAccount {
	return &ProviderUsecaseAccImpl{providerRepo: providerRepo}
}
func (pu *ProviderUsecaseAccImpl) GetProvider(provider *models.Providers) (*models.Providers, bool, error) {
	data, isValid, err := pu.providerRepo.GetProvider(provider)
	fmt.Println(data)
	if err != nil {
		return nil, false, err
	}
	return data, isValid, nil
}
func (pu *ProviderUsecaseAccImpl) CreateProvider(provider *models.Providers) (*models.Providers, error) {
	data, err := pu.providerRepo.CreateProvider(provider)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (pu *ProviderUsecaseAccImpl) GetProviderById(id string) (*models.Providers, error) {
	data, err := pu.providerRepo.GetProviderById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
