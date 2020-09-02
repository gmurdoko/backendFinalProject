package adminaccountmanagementusecase

type AccountManagementUsecase interface {
	DeleteUser(string) (string, error)
	DeleteAsset(string) (string, error)
	DeleteProvider(string) (string, error)
	DeleteComment(string) (string, error)
	ApproveAssets(string) (string, error)
	ApproveAssetsUpdate(string) (string, error)
}
