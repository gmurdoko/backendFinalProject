package asset_controllers

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/asset"
	"finalproject/main/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type AssetLocationHandler struct {
	assetLocation asset_usecases.AssetLocationUsecase
}

func AssetLocationController(r *mux.Router, service asset_usecases.AssetLocationUsecase) {
	assetLocationHandler := AssetLocationHandler{assetLocation: service}
	r.Use(middleware.ActivityLogMiddleware)

	assetLocation := r.PathPrefix("/providerassets").Subrouter()
	assetLocation.HandleFunc("/locations", assetLocationHandler.getAssetLocation).Methods(http.MethodGet)
}

func (s *AssetLocationHandler) getAssetLocation(w http.ResponseWriter, r *http.Request) {
	assetLocation, err := s.assetLocation.ReadAssetsLocation()

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil || assetLocation == nil {
		response.Response = "Data Not Found"
	} else {
		response.Response = assetLocation
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)

}
