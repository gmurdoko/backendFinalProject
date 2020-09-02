package providerListAssetUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/provider/providerListAssetsRepo"
	"log"
)

type ListAssetsUsecaseImpl struct {
	listAssetsRepo providerListAssetsRepo.ListAssetsRepo
}

func InitListAssetsUsecaseImpl(listAssetsRepo providerListAssetsRepo.ListAssetsRepo) ListAssetsUsecase {
	return &ListAssetsUsecaseImpl{listAssetsRepo: listAssetsRepo}
}

func (s *ListAssetsUsecaseImpl) GetAll(userId string) ([]*models.Assets, error) {
	listAssets, err := s.listAssetsRepo.GetAllAssets(userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listAssets, nil
}
