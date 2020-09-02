package usecases

import (
	"finalproject/main/master/models"
)

type AssetCapacityUsecase interface {
	ReadCurrentCapacity(ac *models.AssetCapacity, assetId string) (*models.AssetCapacity, error)
}
