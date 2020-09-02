package providerlistassetsusecase

import (
	"finalproject/main/master/models"
	providerlistassetsrepo "finalproject/main/master/repositories/provider/listAssets"
	"log"
)

type ListAssetsUsecaseImpl struct {
	listAssetsRepo providerlistassetsrepo.ListAssetsRepo
}

func InitListAssetsUsecaseImpl(listAssetsRepo providerlistassetsrepo.ListAssetsRepo) ListAssetsUsecase {
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
