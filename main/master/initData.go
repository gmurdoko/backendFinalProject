package master

import (
	"database/sql"
	"finalproject/main/master/controller"
	"finalproject/main/master/repository/user/ticketrepository"
	"finalproject/main/master/usecase/user/ticketusecase"
	"finalproject/main/middleware"

	"github.com/gorilla/mux"
)

// Init app
func Init(r *mux.Router, db *sql.DB) {
	//Ticket
	ticketRepo := ticketrepository.InitTicketRepositoryImpl(db)
	ticketUsecase := ticketusecase.InitTicketUsecaseImpl(ticketRepo)
	controller.TicketController(r, ticketUsecase)

	//Activity Log Middleware
	r.Use(middleware.ActivityLogMiddleware)
}
