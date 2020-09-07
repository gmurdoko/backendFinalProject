package providerrepository

import "finalproject/main/master/models"

//ProviderRepository interface for ticket
type ProviderRepository interface {
	DeletePhotoProvider(id string) error
	UpdateDataProvider(id string, Provider *models.Providers) (*models.Providers, error)
	UpdatePhotoProvider(photo, id string) error
	SelectPhotoByID(id string) (*string, error)
}
