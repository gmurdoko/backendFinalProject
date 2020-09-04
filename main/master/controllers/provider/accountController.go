package provider

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/provider/providerAccountUsecase"
	"finalproject/utils/jwt"
	"finalproject/utils/response"
	"net/http"

	"github.com/gorilla/mux"
)

type ProviderAccHandler struct {
	providerAccUsecase providerAccountUsecase.ProviderAccount
}

func ProviderAccController(r *mux.Router, service providerAccountUsecase.ProviderAccount) {
	providerHandler := ProviderAccHandler{providerAccUsecase: service}
	provider := r.PathPrefix("/provider").Subrouter()
	provider.HandleFunc("/register", providerHandler.CreateProviders).Methods(http.MethodPost)
	auth := r.PathPrefix("/authProvider").Subrouter()
	auth.HandleFunc("/login", providerHandler.GetProvider).Methods(http.MethodPost)
}
func (ph *ProviderAccHandler) GetProvider(w http.ResponseWriter, r *http.Request) {
	var data models.Providers
	_ = json.NewDecoder(r.Body).Decode(&data)
	provider, isValid, _ := ph.providerAccUsecase.GetProvider(&data)

	if isValid {
		w.Header().Set("Content-type", "application/json")
		token, err := jwt.JwtEncoder(data.Username, "rahasiadong")
		if err != nil {
			http.Error(w, "Failed token generation", http.StatusUnauthorized)
		} else {
			var response response.Response
			response.Status = http.StatusOK
			response.Message = "Success"
			response.Token = token
			response.Data = provider
			byteData, err := json.Marshal(response)
			if err != nil {
				w.Write([]byte("Something Wrong on Marshalling Data"))
			}
			w.Write(byteData)
		}
	} else {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
	}
}
func (ph *ProviderAccHandler) CreateProviders(w http.ResponseWriter, r *http.Request) {
	var providerRequest *models.Providers
	_ = json.NewDecoder(r.Body).Decode(&providerRequest)
	_, err := ph.providerAccUsecase.CreateProvider(providerRequest)
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
