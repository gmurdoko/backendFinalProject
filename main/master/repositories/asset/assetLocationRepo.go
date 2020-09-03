package asset_repositories

import "finalproject/main/master/models"

type AssetsLocationRepo interface {
	ReadAssetsLocation() ([]*models.AssetLocation, error)
}