package adminaccountmanagementrepo

import "finalproject/main/master/models"

type AccountManagementRepo interface {
	DeleteUser(string) error
	DeleteAsset(string) error
	DeleteProvider(string) error
	DeleteComment(string) error
	ApproveAssets(string) error
	ApproveAssetsUpdate(string) error

	// Get all
	GetAllUsers() ([]*models.UserManagement, error)
	GetAllProviders() ([]*models.ProvidersManagement, error)
	GetAllAssets() ([]*models.AssetManagement, error)
	GetAllReviews() ([]*models.ReviewManagement, error)
}
