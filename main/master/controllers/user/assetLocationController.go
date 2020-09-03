package user

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/user/assetusecases"
	"finalproject/main/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type AssetLocationHandler struct {
	assetLocation assetusecases.AssetLocationUsecase
}

func AssetLocationController(r *mux.Router, service assetusecases.AssetLocationUsecase) {
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
