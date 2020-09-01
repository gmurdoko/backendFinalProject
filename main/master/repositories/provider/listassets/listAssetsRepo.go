package listassetsrepo

import "finalproject/main/master/models"

type ListAssetsRepo interface {
	GetAllAssets(string) (*[]models.Assets, error)
}
