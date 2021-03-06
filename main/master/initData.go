package master

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// Init app
func Init(r *mux.Router, db *sql.DB) {
	//Rooms
	roomRepo := repositories.InitRoomRepoImpl(db)
	roomUsecase := usecases.InitRoomUsecaseImpl(roomRepo)
	controllers.RoomController(r, roomUsecase)

	//Transaction
	reserveRepo := repositories.InitReserveRepoImpl(db)
	reserveUsecase := usecases.InitReserveUsecaseImpl(reserveRepo)
	controllers.ReserveController(r, reserveUsecase)

	//Auth
	userRepo := repositories.InitUserRepoImpl(db)
	userUsecase := usecases.InitUserUsecaseImpl(userRepo)
	controllers.UserController(r, userUsecase)
	r.Use(logger.ActivityLogMiddleware)
}
