package providerUsecase

import "finalproject/main/master/models"

type ProviderUsecase interface {
	GetProvider(*models.ProviderModel) (bool, error)
	GetProviderSaldo(string) (int, error)
	CreateProvider(*models.ProviderModel) (*models.ProviderModel, error)
	CreateProviderAsset(*models.AssetModel) (*models.AssetModel, error)
}
