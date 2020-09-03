package usecases

import (
	"finalproject/main/master/models"
)

type AssetCapacityUsecase interface {
	ReadCurrentCapacity(assetId string) (*models.AssetCapacity, error)
}
