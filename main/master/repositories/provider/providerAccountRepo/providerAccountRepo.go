package providerAccountRepo

import "finalproject/main/master/models"

type ProviderAccount interface {
	GetProvider(*models.Providers) (bool, error)
	CreateProvider(*models.Providers) (*models.Providers, error)
}
