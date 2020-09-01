package providerlistassetsusecase

import "finalproject/main/master/models"

type ListAssetsUsecase interface {
	GetAll(string) ([]*models.Assets, error)
}
