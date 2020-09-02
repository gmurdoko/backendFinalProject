package master

import (
	"database/sql"
	"finalproject/main/master/controllers/provider"
	"finalproject/main/master/controllers/user"
	"finalproject/main/master/repositories/provider/providerAccountRepo"
	"finalproject/main/master/repositories/provider/providerHomeRepo"
	"finalproject/main/master/repositories/user/userAccountRepo"
	"finalproject/main/master/repositories/user/userHomeRepo"
	"finalproject/main/master/usecases/provider/providerAccountUsecase"
	"finalproject/main/master/usecases/provider/providerHomeUsecase"
	"finalproject/main/master/usecases/user/userAccountUsecase"
	"finalproject/main/master/usecases/user/userHomeUsecase"

	"finalproject/main/middleware"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB, activityLog bool) {
	// providerRepo := providerRepo.InitProviderRepoImpl(db)
	// providerUsecase := providerUsecase.InitProviderUsecase(providerRepo)
	// controllers.ProviderController(r, providerUsecase)
	// userRepo := userRepo.InitUserRepoImpl(db)
	// userUsecase := userUsecase.InitUserUsecase(userRepo)
	// controllers.UserController(r, userUsecase)
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

	if activityLog == true {
		r.Use(middleware.ActivityLogMiddleware)
	}
}
