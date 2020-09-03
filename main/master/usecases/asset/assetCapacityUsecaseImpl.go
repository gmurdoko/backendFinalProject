package asset_usecases

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/asset"
)

type AssetCapacityUsecaseImpl struct {
	assetCapacity asset_repositories.AssetCapacityRepo
}

func InitAssetCapacityUsecaseImpl(repo asset_repositories.AssetCapacityRepo) AssetCapacityUsecase {
	return &AssetCapacityUsecaseImpl{assetCapacity: repo}
}

func (s *AssetCapacityUsecaseImpl) ReadCurrentCapacity(assetId string) (*models.AssetCapacity, error) {
	capacity, err := s.ReadCurrentCapacity(assetId)
	if err != nil {
		return nil, err
	}
	return capacity, nil
}