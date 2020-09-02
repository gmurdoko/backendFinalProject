package master

import (
	"database/sql"
	"finalproject/main/master/controllers/admin"
	"finalproject/main/master/controllers/provider"
	"finalproject/main/master/controllers/user"
	adminaccountmanagementrepo "finalproject/main/master/repositories/admin/accountManagement"
	adminassetsreportsrepo "finalproject/main/master/repositories/admin/report"
	"finalproject/main/master/repositories/provider/providerAccountRepo"
	"finalproject/main/master/repositories/provider/providerAssetReportsRepo"
	"finalproject/main/master/repositories/provider/providerHomeRepo"
	"finalproject/main/master/repositories/provider/providerListAssetsRepo"
	"finalproject/main/master/repositories/provider/providerrepository"
	"finalproject/main/master/repositories/user/ticketrepository"
	"finalproject/main/master/repositories/user/userAccountRepo"
	"finalproject/main/master/repositories/user/userHomeRepo"
	"finalproject/main/master/repositories/user/walletrepository"
	adminaccountmanagementusecase "finalproject/main/master/usecases/admin/accountManagement"
	"finalproject/main/master/usecases/admin/adminReportUsecase"
	"finalproject/main/master/usecases/provider/providerAccountUsecase"
	"finalproject/main/master/usecases/provider/providerAssetReportsUsecase"
	"finalproject/main/master/usecases/provider/providerHomeUsecase"
	"finalproject/main/master/usecases/provider/providerListAssetUsecase"
	"finalproject/main/master/usecases/provider/providerusecase"
	"finalproject/main/master/usecases/user/ticketusecase"
	"finalproject/main/master/usecases/user/userAccountUsecase"
	"finalproject/main/master/usecases/user/userHomeUsecase"
	"finalproject/main/master/usecases/user/walletusecase"

	"finalproject/main/middleware"

	"github.com/gorilla/mux"
)

// Init app
func Init(r *mux.Router, db *sql.DB, activityLog bool) {
	//Ticket
	ticketRepo := ticketrepository.InitTicketRepositoryImpl(db)
	ticketUsecase := ticketusecase.InitTicketUsecaseImpl(ticketRepo)
	user.TicketController(r, ticketUsecase)

	//Wallet
	walletRepo := walletrepository.InitWalletRepositoryImpl(db)
	walletUsecase := walletusecase.InitWalletUsecaseImpl(walletRepo)
	user.WalletController(r, walletUsecase)

	// Provider
	providerRepo := providerrepository.InitProviderRepositoryImpl(db)
	providerUsecase := providerusecase.InitProviderUsecaseImpl(providerRepo)
	provider.ProviderController(r, providerUsecase)
	//Activity Log Middleware

	providerAccRepo := providerAccountRepo.InitProviderRepoAccImpl(db)
	providerAccUsecase := providerAccountUsecase.InitProviderAccUsecase(providerAccRepo)
	provider.ProviderAccController(r, providerAccUsecase)

	providerHomeRepo := providerHomeRepo.InitProviderHomeRepoImpl(db)
	providerHomeUsecase := providerHomeUsecase.InitProviderHomeUsecase(providerHomeRepo)
	provider.ProviderHomeController(r, providerHomeUsecase)

	userAccRepo := userAccountRepo.InitUserAccRepoImpl(db)
	userAccUsecase := userAccountUsecase.InitUseAccUsecase(userAccRepo)
	user.UserAccController(r, userAccUsecase)

	userHomeRepo := userHomeRepo.InitUserHomeRepoImpl(db)
	userHomeUsecase := userHomeUsecase.InitUserHomeUsecase(userHomeRepo)
	user.UserHomeController(r, userHomeUsecase)

	listAssetsRepo := providerListAssetsRepo.InitListAssetsRepoImpl(db)
	listAssetsUsecase := providerListAssetUsecase.InitListAssetsUsecaseImpl(listAssetsRepo)
	provider.ListAssetsController(r, listAssetsUsecase)

	providerAssetsReportRepo := providerAssetReportsRepo.InitProviderAssetReportRepoImpl(db)
	providerAssetsReportUsecase := providerAssetReportsUsecase.InitProviderReportUsecaseImpl(providerAssetsReportRepo)
	provider.ProviderAssetReportController(r, providerAssetsReportUsecase)

	adminAssetsReportRepo := adminassetsreportsrepo.InitAdminAssetReportRepoImpl(db)
	adminAssetsReportUsecase := adminReportUsecase.InitProviderReportUsecaseImpl(adminAssetsReportRepo)
	admin.AdminAssetReportController(r, adminAssetsReportUsecase)

	adminAccountManagementRepo := adminaccountmanagementrepo.InitAccountManagementImpl(db)
	adminAccountManagementUsecase := adminaccountmanagementusecase.InitAccountManagementUsecaseImpl(adminAccountManagementRepo)
	admin.AccountManagerController(r, adminAccountManagementUsecase)

	if activityLog == true {
		r.Use(middleware.ActivityLogMiddleware)
	}
}
