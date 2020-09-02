package providerHomeUsecase

import "finalproject/main/master/models"

type ProviderHome interface {
	GetProviderSaldo(string) (int, error)
	CreateAssetProvider(*models.AssetModel) (*models.AssetModel, error)
}
