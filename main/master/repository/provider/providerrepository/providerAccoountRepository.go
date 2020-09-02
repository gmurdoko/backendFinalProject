package providerrepository

import "finalproject/main/master/model"

//ProviderRepository interface for ticket
type ProviderRepository interface {
	DeletePhotoProvider(id string) error
	UpdateDataProvider(Provider *model.Providers) error
}
