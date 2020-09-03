package controllers

import (
	"encoding/json"
	"finalproject/main/master/models"
	usecases "finalproject/main/master/usecases/asset/assetsCapacity"
	"finalproject/main/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type AssetCapacityHandler struct {
	assetCapacity usecases.AssetCapacityUsecase
}

func AssetCapacityController(r *mux.Router, service usecases.AssetCapacityUsecase) {
	assetCapacityHandler := AssetCapacityHandler{assetCapacity: service}

	r.Use(middleware.ActivityLogMiddleware)
	assetCapacity := r.PathPrefix("/providerassets").Subrouter()
	assetCapacity.HandleFunc("/currentcap/{id}", assetCapacityHandler.getAssetCapacity).Methods(http.MethodGet)
}

func (s *AssetCapacityHandler) getAssetCapacity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assetId := params["id"]
	assetCap, err := s.assetCapacity.ReadCurrentCapacity(assetId)

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil || assetCap == nil {
		response.Response = "Data Not Found"
	} else {
		response.Response = assetCap
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
