package repositories

import "finalproject/main/master/models"

type AssetCapacityRepo interface {
	ReadCurrentCapacity(ac *models.AssetCapacity, assetId string) (*models.AssetCapacity, error)
}






