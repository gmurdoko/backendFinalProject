package controllers

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/providerUsecase"
	"finalproject/utils/response"
	"net/http"

	"github.com/gorilla/mux"
)

type ProviderHandler struct {
	providerUsecase providerUsecase.ProviderUsecase
}

func ProviderController(r *mux.Router, service providerUsecase.ProviderUsecase) {
	providerHandler := ProviderHandler{providerUsecase: service}
	provider := r.PathPrefix("/provider").Subrouter()
	provider.HandleFunc("", providerHandler.CreateProviders).Methods(http.MethodPost)
	provider.HandleFunc("/asset", providerHandler.CreateProviderAsset).Methods(http.MethodPost)
}

func (ph *ProviderHandler) CreateProviders(w http.ResponseWriter, r *http.Request) {
	var providerRequest *models.ProviderModel
	_ = json.NewDecoder(r.Body).Decode(&providerRequest)
	_, err := ph.providerUsecase.CreateProvider(providerRequest)
	if err != nil {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = "Fail"
		w.Write([]byte("Cannot Add Data"))
	} else {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = providerRequest
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
	}
}
func (ph *ProviderHandler) CreateProviderAsset(w http.ResponseWriter, r *http.Request) {
	var assetProviderRequest *models.AssetModel
	_ = json.NewDecoder(r.Body).Decode(&assetProviderRequest)
	_, err := ph.providerUsecase.CreateProviderAsset(assetProviderRequest)
	if err != nil {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = "Fail"
		w.Write([]byte("Cannot Add Data"))
	} else {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = assetProviderRequest
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
	}
}
