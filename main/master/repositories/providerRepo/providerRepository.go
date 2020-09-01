package providerRepo

import "finalproject/main/master/models"

type ProviderRepository interface {
	CreateProvider(*models.ProviderModel) (*models.ProviderModel, error)
	CreateAssetProvider(*models.AssetModel) (*models.AssetModel, error)
}
