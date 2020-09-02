package providerrepository

import "finalproject/main/master/models"

//ProviderRepository interface for ticket
type ProviderRepository interface {
	DeletePhotoProvider(id string) error
	UpdateDataProvider(Provider *models.ProviderModel) error
}
