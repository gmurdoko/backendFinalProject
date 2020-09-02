package usecases

import (
	"finalproject/main/master/models"
	repositories "finalproject/main/master/repositories/asset/assetsLocation"
	"log"
)

type AssetLocationUsecaseImpl struct{
	assetsLocationRepo repositories.AssetsLocationRepo
}

func InitAssetLocationUsecaseImpl(repo repositories.AssetsLocationRepo) AssetLocationUsecase {
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
