package providerusecase

import (
	"finalproject/main/master/models"
	"mime/multipart"
)

//ProviderUsecase usecase interface
type ProviderUsecase interface {
	DeleteProviderFoto(id string) error
	UpdateDataProvider(Provider *models.Providers) error
	UpdateProviderFoto(photo multipart.File, handler *multipart.FileHeader, id string) error
	GetPhotoByID(id string) (*string, error)
}