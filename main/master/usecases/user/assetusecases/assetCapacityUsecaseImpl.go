package assetusecases

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/user/assetrepositories"
)

type AssetCapacityUsecaseImpl struct {
	assetCapacity assetrepositories.AssetCapacityRepo
}

func InitAssetCapacityUsecaseImpl(repo assetrepositories.AssetCapacityRepo) AssetCapacityUsecase {
	return &AssetCapacityUsecaseImpl{assetCapacity: repo}
}

func (s *AssetCapacityUsecaseImpl) ReadCurrentCapacity(assetId string) (*models.AssetCapacity, error) {
	capacity, err := s.ReadCurrentCapacity(assetId)
	if err != nil {
		return nil, err
	}
	return capacity, nil
}
