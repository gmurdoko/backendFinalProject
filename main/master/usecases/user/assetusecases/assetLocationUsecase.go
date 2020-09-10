package assetusecases

import "finalproject/main/master/models"

type AssetLocationUsecase interface {
	ReadAssetsLocation() ([]*models.AssetLocation, error)
	GetAssetByID(string) (*models.AssetLocation, error)
}
