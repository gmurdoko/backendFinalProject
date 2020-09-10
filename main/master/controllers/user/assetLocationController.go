package user

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/user/assetusecases"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AssetLocationHandler struct {
	// assetLocation assetusecases.AssetLocationUsecase
	assetLocation assetusecases.AssetLocationUsecase
}

func AssetLocationController(r *mux.Router, service assetusecases.AssetLocationUsecase) {
	assetLocationHandler := AssetLocationHandler{assetLocation: service}
	// r.Use(middleware.ActivityLogMiddleware)

	assetLocation := r.PathPrefix("/providerassets").Subrouter()
	assetLocation.HandleFunc("/locations", assetLocationHandler.getAssetLocation).Methods(http.MethodGet)
	assetLocation.HandleFunc("/{id}", assetLocationHandler.getByID).Methods(http.MethodGet)
}

func (s *AssetLocationHandler) getAssetLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("masuk controller")
	assetLocation, err := s.assetLocation.ReadAssetsLocation()

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil {
		log.Println(err)
		response.Response = "Data Not Found"
	} else {
		log.Println(assetLocation[0].AssetName)
		response.Response = assetLocation
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)

}

func (s *AssetLocationHandler) getByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	assetLocation, err := s.assetLocation.GetAssetByID(id)

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil {
		log.Println(err)
		response.Response = "Data Not Found"
	} else {
		// log.Println(assetLocation[0].AssetName)
		response.Response = assetLocation
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
