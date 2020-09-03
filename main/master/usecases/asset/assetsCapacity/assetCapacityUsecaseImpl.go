package usecases

import (
	"finalproject/main/master/models"
	repositories "finalproject/main/master/repositories/asset/assetsCapacity"
)

type AssetCapacityUsecaseImpl struct {
	assetCapacity repositories.AssetCapacityRepo
}

func InitAssetCapacityUsecaseImpl(repo repositories.AssetCapacityRepo) AssetCapacityUsecase {
	return &AssetCapacityUsecaseImpl{assetCapacity: repo}
}

func (s *AssetCapacityUsecaseImpl) ReadCurrentCapacity(assetId string) (*models.AssetCapacity, error) {
	capacity, err := s.ReadCurrentCapacity(assetId)
	if err != nil {
		return nil, err
	}
	return capacity, nil
}