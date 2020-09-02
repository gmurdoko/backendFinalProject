package providerusecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/provider/providerrepository"
)

//providerUsecaseImpl app
type providerUsecaseImpl struct {
	providerRepository providerrepository.ProviderRepository
}

//DeleteProvider app
func (s providerUsecaseImpl) DeleteProviderFoto(id string) error {
	err := s.providerRepository.DeletePhotoProvider(id)
	if err != nil {
		return err
	}
	return nil
}

func (s providerUsecaseImpl) UpdateDataProvider(Provider *models.ProviderModel) error {
	err := s.providerRepository.UpdateDataProvider(Provider)
	if err != nil {
		return err
	}
	return nil
}

//InitProviderUsecaseImpl app
func InitProviderUsecaseImpl(providerRepository providerrepository.ProviderRepository) ProviderUsecase {
	return &providerUsecaseImpl{providerRepository}
}
