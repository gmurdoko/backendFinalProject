package master

import (
	"database/sql"
	"finalproject/main/master/controller"
	"finalproject/main/master/repository/user/ticketrepository"
	"finalproject/main/master/repository/user/walletrepository"
	"finalproject/main/master/usecase/user/ticketusecase"
	"finalproject/main/master/usecase/user/walletusecase"
	"finalproject/main/middleware"

	"github.com/gorilla/mux"
)

// Init app
func Init(r *mux.Router, db *sql.DB) {
	//Ticket
	ticketRepo := ticketrepository.InitTicketRepositoryImpl(db)
	ticketUsecase := ticketusecase.InitTicketUsecaseImpl(ticketRepo)
	controller.TicketController(r, ticketUsecase)

	//Wallet
	walletRepo := walletrepository.InitWalletRepositoryImpl(db)
	walletUsecase := walletusecase.InitWalletUsecaseImpl(walletRepo)
	controller.WalletController(r, walletUsecase)

	//Activity Log Middleware
	r.Use(middleware.ActivityLogMiddleware)
}
