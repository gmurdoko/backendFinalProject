package master

import (
	"database/sql"
	"finalproject/main/master/controllers"
	"finalproject/main/master/repositories/providerRepo"
	"finalproject/main/master/repositories/userRepo"
	"finalproject/main/master/usecases/providerUsecase"
	"finalproject/main/master/usecases/userUsecase"
	"finalproject/main/middleware"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB, activityLog bool) {
	providerRepo := providerRepo.InitProviderRepoImpl(db)
	providerUsecase := providerUsecase.InitProviderUsecase(providerRepo)
	controllers.ProviderController(r, providerUsecase)
	userRepo := userRepo.InitUserRepoImpl(db)
	userUsecase := userUsecase.InitUserUsecase(userRepo)
	controllers.UserController(r, userUsecase)
	if activityLog == true {
		r.Use(middleware.ActivityLogMiddleware)
	}
}
