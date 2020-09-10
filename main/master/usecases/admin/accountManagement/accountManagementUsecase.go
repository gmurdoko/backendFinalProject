package adminaccountmanagementusecase

import "finalproject/main/master/models"

type AccountManagementUsecase interface {
	DeleteUser(string) (string, error)
	DeleteAsset(string) (string, error)
	DeleteProvider(string) (string, error)
	DeleteComment(string) (string, error)
	ApproveAssets(string) (string, error)
	ApproveAssetsUpdate(string) (string, error)

	// Get all
	GetAllUsers() ([]*models.UserManagement, error)
	GetAllProviders() ([]*models.ProvidersManagement, error)
	GetAllAssets() ([]*models.AssetManagement, error)
	GetAllReviews() ([]*models.ReviewManagement, error)
}
