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

func (s providerUsecaseImpl) UpdateDataProvider(Provider *models.Providers) error {
	err := s.providerRepository.UpdateDataProvider(Provider)
	if err != nil {
		return err
	}
	return nil
}

func (s providerUsecaseImpl) UpdateProviderFoto(photo, id string) error {
	err := s.providerRepository.UpdatePhotoProvider(photo, id)
	if err != nil {
		return err
	}
	return nil
}

func (s providerUsecaseImpl) GetPhotoByID(id string) (*string, error) {
	photo, err := s.providerRepository.SelectPhotoByID(id)
	if err != nil {
		return nil, err
	}
	return photo, nil
}

//InitProviderUsecaseImpl app
func InitProviderUsecaseImpl(providerRepository providerrepository.ProviderRepository) ProviderUsecase {
	return &providerUsecaseImpl{providerRepository}
}
