package controllers

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/providerUsecase"
	"finalproject/utils/jwt"
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
	provider.HandleFunc("/saldo", providerHandler.GetProviderSaldo).Methods(http.MethodGet)
	provider.HandleFunc("/asset", providerHandler.CreateProviderAsset).Methods(http.MethodPost)
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("", providerHandler.GetProvider).Methods(http.MethodPost)
}
func (ph *ProviderHandler) GetProvider(w http.ResponseWriter, r *http.Request) {
	var data models.ProviderModel
	_ = json.NewDecoder(r.Body).Decode(&data)
	isValid, _ := ph.providerUsecase.GetProvider(&data)

	if isValid {
		w.WriteHeader(http.StatusOK)
		token, err := jwt.JwtEncoder(data.Username, "rahasiadong")
		if err != nil {
			http.Error(w, "Failed token generation", http.StatusUnauthorized)
		} else {
			w.Write([]byte(token))
		}
	} else {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
	}
}
func (ph *ProviderHandler) GetProviderSaldo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, err := ph.providerUsecase.GetProviderSaldo(id)
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
