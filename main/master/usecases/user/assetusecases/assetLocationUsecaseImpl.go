package assetusecases

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/user/assetrepositories"
	"log"
)

type AssetLocationUsecaseImpl struct {
	assetsLocationRepo assetrepositories.AssetsLocationRepo
}

func InitAssetLocationUsecaseImpl(repo assetrepositories.AssetsLocationRepo) AssetLocationUsecase {
	return &AssetLocationUsecaseImpl{assetsLocationRepo: repo}
}

func (s *AssetLocationUsecaseImpl) ReadAssetsLocation() ([]*models.AssetLocation, error) {
	//panic("implement me")
	log.Println("masuk usecase")
	listLocation, err := s.assetsLocationRepo.ReadAssetsLocation()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listLocation, nil
}
