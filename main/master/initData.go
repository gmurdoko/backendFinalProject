package master

import (
	"database/sql"
	"finalproject/main/master/controllers/asset"
	"finalproject/main/master/controllers/review"
	"finalproject/main/master/controllers/ticket"
	"finalproject/main/master/repositories/asset"
	"finalproject/main/master/repositories/review"
	"finalproject/main/master/repositories/ticket"
	"finalproject/main/master/usecases/asset"
	"finalproject/main/master/usecases/review"
	"finalproject/main/master/usecases/ticket"

	"github.com/gorilla/mux"
)

// Init app
func Init(r *mux.Router, db *sql.DB) {
	assetCapRepo := asset_repositories.InitAssetCapacityRepoImpl(db)
	assetCapUsecase := asset_usecases.InitAssetCapacityUsecaseImpl(assetCapRepo)
	asset_controllers.AssetCapacityController(r, assetCapUsecase)

	assetLocRepo := asset_repositories.InitAssetsLocationRepoImpl(db)
	assetLocUsecase := asset_usecases.InitAssetLocationUsecaseImpl(assetLocRepo)
	asset_controllers.AssetLocationController(r, assetLocUsecase)

	reviewRepo := review_repositories.InitReviewRepoImpl(db)
	reviewUsecase := review_usecases.InitReviewUsecaseImpl(reviewRepo)
	review_controllers.ReviewController(r, reviewUsecase)

	ticketRepo := ticket_repositories.InitTicketRepoImpl(db)
	ticketUsecase := ticket_usecases.InitTicketUsecase(ticketRepo)
	ticket_controllers.TicketController(r, ticketUsecase)

}
