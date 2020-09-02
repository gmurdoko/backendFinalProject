package providerHomeUsecase

import (
	"finalproject/main/master/models"
	"mime/multipart"
)

type ProviderHome interface {
	GetProviderSaldo(string) (int, error)
	CreateAssetProvider(photo multipart.File, handler *multipart.FileHeader, provider *models.AssetModel) (*models.AssetModel, error)
}
