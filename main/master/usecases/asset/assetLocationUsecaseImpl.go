package asset_usecases

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/asset"
	"log"
)

type AssetLocationUsecaseImpl struct{
	assetsLocationRepo asset_repositories.AssetsLocationRepo
}

func InitAssetLocationUsecaseImpl(repo asset_repositories.AssetsLocationRepo) AssetLocationUsecase {
	return &AssetLocationUsecaseImpl{assetsLocationRepo: repo}
}

func (s *AssetLocationUsecaseImpl) ReadAssetsLocation() ([]*models.AssetLocation, error) {
	//panic("implement me")
	listLocation, err := s.assetsLocationRepo.ReadAssetsLocation()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listLocation, nil
}
