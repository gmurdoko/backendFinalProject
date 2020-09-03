package repositories

import "finalproject/main/master/models"

type AssetCapacityRepo interface {
	ReadCurrentCapacity(assetId string) (*models.AssetCapacity, error)
}






