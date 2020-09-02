package master

import (
	"database/sql"
	controllersAdmin "finalproject/main/master/controllers/admin"
	controllersProvider "finalproject/main/master/controllers/provider"

	adminaccountmanagementrepo "finalproject/main/master/repositories/admin/accountManagement"
	adminassetsreportsrepo "finalproject/main/master/repositories/admin/report"
	providerassetsreportsrepo "finalproject/main/master/repositories/provider/assetReport"
	providerlistassetsrepo "finalproject/main/master/repositories/provider/listAssets"

	adminaccountmanagementusecase "finalproject/main/master/usecases/admin/accountManagement"
	adminassetsreportsusecase "finalproject/main/master/usecases/admin/report"
	providerassetreportsusecase "finalproject/main/master/usecases/provider/assetReport"
	providerlistassetsusecase "finalproject/main/master/usecases/provider/listassets"

	"github.com/gorilla/mux"
)

// Init app
func Init(r *mux.Router, db *sql.DB) {
	listAssetsRepo := providerlistassetsrepo.InitListAssetsRepoImpl(db)
	listAssetsUsecase := providerlistassetsusecase.InitListAssetsUsecaseImpl(listAssetsRepo)
	controllersProvider.ListAssetsController(r, listAssetsUsecase)

	providerAssetsReportRepo := providerassetsreportsrepo.InitProviderAssetReportRepoImpl(db)
	providerAssetsReportUsecase := providerassetreportsusecase.InitProviderReportUsecaseImpl(providerAssetsReportRepo)
	controllersProvider.ProviderAssetReportController(r, providerAssetsReportUsecase)

	adminAssetsReportRepo := adminassetsreportsrepo.InitAdminAssetReportRepoImpl(db)
	adminAssetsReportUsecase := adminassetsreportsusecase.InitProviderReportUsecaseImpl(adminAssetsReportRepo)
	controllersAdmin.AdminAssetReportController(r, adminAssetsReportUsecase)

	adminAccountManagementRepo := adminaccountmanagementrepo.InitAccountManagementImpl(db)
	adminAccountManagementUsecase := adminaccountmanagementusecase.InitAccountManagementUsecaseImpl(adminAccountManagementRepo)
	controllersAdmin.AccountManagerController(r, adminAccountManagementUsecase)
}
