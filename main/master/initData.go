package master

import (
	"database/sql"
	controllersAdmin "finalproject/main/master/controllers/admin"
	controllers "finalproject/main/master/controllers/provider"

	adminaccountmanagementrepo "finalproject/main/master/repositories/admin/accountManagement"
	adminassetsreportsrepo "finalproject/main/master/repositories/admin/report"
	providerassetsreportsrepo "finalproject/main/master/repositories/provider/assetReport"
	providerlistassetsrepo "finalproject/main/master/repositories/provider/listAssets"

	"finalproject/main/master/controllers/user"
	"finalproject/main/master/repositories/provider/providerAccountRepo"
	"finalproject/main/master/repositories/provider/providerHomeRepo"
	"finalproject/main/master/repositories/user/userAccountRepo"
	"finalproject/main/master/repositories/user/userHomeRepo"
	adminaccountmanagementusecase "finalproject/main/master/usecases/admin/accountManagement"
	adminassetsreportsusecase "finalproject/main/master/usecases/admin/report"
	providerassetreportsusecase "finalproject/main/master/usecases/provider/assetReport"
	providerlistassetsusecase "finalproject/main/master/usecases/provider/listassets"
	"finalproject/main/master/usecases/provider/providerAccountUsecase"
	"finalproject/main/master/usecases/provider/providerHomeUsecase"
	"finalproject/main/master/usecases/user/userAccountUsecase"
	"finalproject/main/master/usecases/user/userHomeUsecase"

	"finalproject/main/master/controller"
	"finalproject/main/master/repository/user/ticketrepository"
	"finalproject/main/master/usecase/user/ticketusecase"
	"finalproject/main/middleware"

	"github.com/gorilla/mux"
)

// Init app
func Init(r *mux.Router, db *sql.DB, activityLog bool) {
	listAssetsRepo := providerlistassetsrepo.InitListAssetsRepoImpl(db)
	listAssetsUsecase := providerlistassetsusecase.InitListAssetsUsecaseImpl(listAssetsRepo)
	controllers.ListAssetsController(r, listAssetsUsecase)

	providerAssetsReportRepo := providerassetsreportsrepo.InitProviderAssetReportRepoImpl(db)
	providerAssetsReportUsecase := providerassetreportsusecase.InitProviderReportUsecaseImpl(providerAssetsReportRepo)
	controllers.ProviderAssetReportController(r, providerAssetsReportUsecase)

	adminAssetsReportRepo := adminassetsreportsrepo.InitAdminAssetReportRepoImpl(db)
	adminAssetsReportUsecase := adminassetsreportsusecase.InitProviderReportUsecaseImpl(adminAssetsReportRepo)
	controllersAdmin.AdminAssetReportController(r, adminAssetsReportUsecase)

	adminAccountManagementRepo := adminaccountmanagementrepo.InitAccountManagementImpl(db)
	adminAccountManagementUsecase := adminaccountmanagementusecase.InitAccountManagementUsecaseImpl(adminAccountManagementRepo)
	controllersAdmin.AccountManagerController(r, adminAccountManagementUsecase)

	providerAccRepo := providerAccountRepo.InitProviderRepoAccImpl(db)
	providerAccUsecase := providerAccountUsecase.InitProviderAccUsecase(providerAccRepo)
	controllers.ProviderAccController(r, providerAccUsecase)
	providerHomeRepo := providerHomeRepo.InitProviderHomeRepoImpl(db)
	providerHomeUsecase := providerHomeUsecase.InitProviderHomeUsecase(providerHomeRepo)
	controllers.ProviderHomeController(r, providerHomeUsecase)

	userAccRepo := userAccountRepo.InitUserAccRepoImpl(db)
	userAccUsecase := userAccountUsecase.InitUseAccUsecase(userAccRepo)
	user.UserAccController(r, userAccUsecase)
	userHomeRepo := userHomeRepo.InitUserHomeRepoImpl(db)
	userHomeUsecase := userHomeUsecase.InitUserHomeUsecase(userHomeRepo)
	user.UserHomeController(r, userHomeUsecase)

	ticketRepo := ticketrepository.InitTicketRepositoryImpl(db)
	ticketUsecase := ticketusecase.InitTicketUsecaseImpl(ticketRepo)
	controller.TicketController(r, ticketUsecase)

	if activityLog == true {
		r.Use(middleware.ActivityLogMiddleware)
	}
}
