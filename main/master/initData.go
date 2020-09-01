package master

import (
	"database/sql"
	controllers "finalproject/main/master/controllers/provider"
	providerlistassetsrepo "finalproject/main/master/repositories/provider/listAssets"
	providerlistassetsusecase "finalproject/main/master/usecases/provider/listassets"

	"github.com/gorilla/mux"
)

// Init app
func Init(r *mux.Router, db *sql.DB) {
	listAssetsRepo := providerlistassetsrepo.InitListAssetsRepoImpl(db)
	listAssetsUsecase := providerlistassetsusecase.InitListAssetsUsecaseImpl(listAssetsRepo)
	controllers.ListAssetsController(r, listAssetsUsecase)
}
