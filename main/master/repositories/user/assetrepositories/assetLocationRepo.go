package assetrepositories

import "finalproject/main/master/models"

type AssetsLocationRepo interface {
	ReadAssetsLocation() ([]*models.AssetLocation, error)
	GetAssetsByID(string) (*models.AssetLocation, error)
}
