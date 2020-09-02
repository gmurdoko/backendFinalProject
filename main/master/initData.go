package master

import (
	"database/sql"
	"finalproject/main/master/controllers"
	"finalproject/main/master/repositories/provider/providerrepository"
	"finalproject/main/master/repositories/user/ticketrepository"
	"finalproject/main/master/repositories/user/walletrepository"
	"finalproject/main/master/usecases/provider/providerusecase"
	"finalproject/main/master/usecases/user/ticketusecase"
	"finalproject/main/master/usecases/user/walletusecase"
	"finalproject/main/middleware"

	"github.com/gorilla/mux"
)

// Init app
func Init(r *mux.Router, db *sql.DB) {
	//Ticket
	ticketRepo := ticketrepository.InitTicketRepositoryImpl(db)
	ticketUsecase := ticketusecase.InitTicketUsecaseImpl(ticketRepo)
	controllers.TicketController(r, ticketUsecase)

	//Wallet
	walletRepo := walletrepository.InitWalletRepositoryImpl(db)
	walletUsecase := walletusecase.InitWalletUsecaseImpl(walletRepo)
	controllers.WalletController(r, walletUsecase)

	// Provider
	providerRepo := providerrepository.InitProviderRepositoryImpl(db)
	providerUsecase := providerusecase.InitProviderUsecaseImpl(providerRepo)
	controllers.ProviderController(r, providerUsecase)
	//Activity Log Middleware
	r.Use(middleware.ActivityLogMiddleware)
}
