package usecases

import "finalproject/main/master/models"

type AssetLocationUsecase interface {
	ReadAssetsLocation() ([]*models.AssetLocation, error)
}

