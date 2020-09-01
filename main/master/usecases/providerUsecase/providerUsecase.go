package providerUsecase

import "finalproject/main/master/models"

type ProviderUsecase interface {
	CreateProvider(*models.ProviderModel) (*models.ProviderModel, error)
	CreateProviderAsset(*models.AssetModel) (*models.AssetModel, error)
}
