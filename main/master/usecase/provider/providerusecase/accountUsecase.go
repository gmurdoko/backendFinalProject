package providerusecase

import "finalproject/main/master/model"

//ProviderUsecase usecase interface
type ProviderUsecase interface {
	DeleteProviderFoto(id string) error
	UpdateDataProvider(Provider *model.Providers) error
}
