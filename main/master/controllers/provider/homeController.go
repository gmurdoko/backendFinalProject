package controllers

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/provider/providerHomeUsecase"
	"finalproject/utils/response"
	"net/http"

	"github.com/gorilla/mux"
)

type ProviderHomeHandler struct {
	providerHomeUsecase providerHomeUsecase.ProviderHome
}

func ProviderHomeController(r *mux.Router, service providerHomeUsecase.ProviderHome) {
	providerHandler := ProviderHomeHandler{providerHomeUsecase: service}
	provider := r.PathPrefix("/provider").Subrouter()
	provider.HandleFunc("/saldo", providerHandler.GetProviderSaldo).Methods(http.MethodGet)
	provider.HandleFunc("/asset", providerHandler.CreateProviderAsset).Methods(http.MethodPost)
}
func (ph *ProviderHomeHandler) GetProviderSaldo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, err := ph.providerHomeUsecase.GetProviderSaldo(id)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response response.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = data
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)

}
func (ph *ProviderHomeHandler) CreateProviderAsset(w http.ResponseWriter, r *http.Request) {
	var assetProviderRequest *models.AssetModel
	_ = json.NewDecoder(r.Body).Decode(&assetProviderRequest)
	_, err := ph.providerHomeUsecase.CreateAssetProvider(assetProviderRequest)
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
