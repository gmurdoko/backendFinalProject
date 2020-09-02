package providerRepo

import "finalproject/main/master/models"

type ProviderRepository interface {
	GetProvider(*models.ProviderModel) (bool, error)
	GetProviderSaldo(string) (int, error)
	CreateProvider(*models.ProviderModel) (*models.ProviderModel, error)
	CreateAssetProvider(*models.AssetModel) (*models.AssetModel, error)
}
