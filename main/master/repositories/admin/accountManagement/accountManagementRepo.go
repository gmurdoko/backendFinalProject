package adminaccountmanagementrepo

type AccountManagementRepo interface {
	DeleteUser(string) error
	DeleteAsset(string) error
	DeleteProvider(string) error
	DeleteComment(string) error
	ApproveAssets(string) error
	ApproveAssetsUpdate(string) error
}
