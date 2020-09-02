package assetsLocation

import "finalproject/main/master/models"

type AssetsLocationRepo interface {
	ReadAssetsLocation() ([]*models.Asset, error)
}