package providerHomeRepo

import "finalproject/main/master/models"

type ProviderHome interface {
	GetProviderSaldo(string) (string, error)
	CreateAssetProvider(*models.AssetModel) (*models.AssetModel, error)
	GetAssetReview(string) ([]*models.AssetReview, error)
}
