package providerAccountRepo

import "finalproject/main/master/models"

type ProviderAccount interface {
	GetProvider(*models.ProviderModel) (bool, error)
	CreateProvider(*models.ProviderModel) (*models.ProviderModel, error)
}
