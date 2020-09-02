package providerusecase

import "finalproject/main/master/models"

//ProviderUsecase usecase interface
type ProviderUsecase interface {
	DeleteProviderFoto(id string) error
	UpdateDataProvider(Provider *models.Providers) error
	UpdateProviderFoto(photo, id string) error
	GetPhotoByID(id string) (*string, error)
}
