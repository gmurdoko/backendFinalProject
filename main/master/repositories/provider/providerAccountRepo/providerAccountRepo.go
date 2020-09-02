package providerAccountRepo

import "finalproject/main/master/models"

type ProviderAccount interface {
	GetProvider(*models.Providers) (*models.Providers, bool, error)
	CreateProvider(*models.Providers) (*models.Providers, error)
	GetProviderById(string) (*models.Providers, error)
}
