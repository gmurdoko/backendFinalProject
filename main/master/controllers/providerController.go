package controllers

import (
	"encoding/json"
	"finalproject/config"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/provider/providerusecase"
	"finalproject/main/middleware"
	"finalproject/utils/response"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//ProviderHandler app
type ProviderHandler struct {
	providerUsecase providerusecase.ProviderUsecase
}

//ProviderController app
func ProviderController(r *mux.Router, s providerusecase.ProviderUsecase) {
	providerHandler := ProviderHandler{s}
	providers := r.PathPrefix("/providers").Subrouter()
	provider := r.PathPrefix("/provider").Subrouter()
	isAuthOn := config.AuthSwitch()
	if isAuthOn {
		providers.Use(middleware.TokenValidationMiddleware)
		provider.Use(middleware.TokenValidationMiddleware)
		detailProviderController(providers, provider, providerHandler)
	} else {
		detailProviderController(providers, provider, providerHandler)
	}
}

func detailProviderController(providers, provider *mux.Router, providerHandler ProviderHandler) {
	//Jamak
	// providers.HandleFunc("", providerHandler.Listproviders).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}", "status", "{status}", "orderBy", "{orderBy}", "sort", "{sort}").Methods(http.MethodGet)
	// providers.HandleFunc("/available", providerHandler.ListAvailableproviders).Methods(http.MethodGet)
	// providers.HandleFunc("/booked", providerHandler.ListBookedproviders).Methods(http.MethodGet)
	//Satuan
	// provider.HandleFunc("/{id}", providerHandler.provider).Methods(http.MethodGet)
	// provider.HandleFunc("", providerHandler.Postprovider).Methods(http.MethodPost)
	// provider.HandleFunc("", providerHandler.Putprovider).Methods(http.MethodPut)
	provider.HandleFunc("/photo/", providerHandler.GetProviderPhoto).Queries("id", "{id}").Methods(http.MethodGet)
	provider.HandleFunc("/photo/", providerHandler.UpdateProviderPhoto).Queries("photo", "{photo}", "id", "{id}").Methods(http.MethodPut)
	provider.HandleFunc("/photo/", providerHandler.DeleteProviderPhoto).Queries("id", "{id}").Methods(http.MethodDelete)
	provider.HandleFunc("/data/", providerHandler.PutDataProvider).Queries("id", "{id}").Methods(http.MethodPut)

}

// DeleteProviderPhoto app
func (s *ProviderHandler) DeleteProviderPhoto(w http.ResponseWriter, r *http.Request) {
	ex := mux.Vars(r)
	id := ex["id"]
	var providerResponse response.Response
	w.Header().Set("content-type", "application/json")
	err := s.providerUsecase.DeleteProviderFoto(id)
	if err != nil {
		providerResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		response.ResponseWrite(&providerResponse, w)
		log.Println(err)
	} else {
		providerResponse = response.Response{Status: http.StatusAccepted, Message: "Delete Provider Foto Success", Data: id}
		response.ResponseWrite(&providerResponse, w)
	}

	log.Println("Endpoint hit: Delete Provider Foto")
}

//PutDataProvider app
func (s *ProviderHandler) PutDataProvider(w http.ResponseWriter, r *http.Request) {
	ex := mux.Vars(r)
	id := ex["id"]
	var inProvider models.Providers
	var providerResponse response.Response
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&inProvider)
	if err != nil {
		log.Println(err)
	}
	inProvider.ID = id
	err = s.providerUsecase.UpdateDataProvider(&inProvider)
	if err != nil {
		providerResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		response.ResponseWrite(&providerResponse, w)
		log.Println(err)
	} else {
		providerResponse = response.Response{Status: http.StatusAccepted, Message: "Update Data Provider Success", Data: inProvider}
		response.ResponseWrite(&providerResponse, w)
	}

	log.Println("Endpoint hit: Update Data Provider")
}

//UpdateProviderPhoto app
func (s *ProviderHandler) UpdateProviderPhoto(w http.ResponseWriter, r *http.Request) {
	photo := mux.Vars(r)["photo"]
	id := mux.Vars(r)["id"]
	var providerResponse response.Response
	w.Header().Set("content-type", "application/json")
	err := s.providerUsecase.UpdateProviderFoto(photo, id)
	if err != nil {
		providerResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		response.ResponseWrite(&providerResponse, w)
		log.Println(err)
	} else {
		providerResponse = response.Response{Status: http.StatusAccepted, Message: "Update Provider Foto Success", Data: id}
		response.ResponseWrite(&providerResponse, w)
	}

	log.Println("Endpoint hit: Update Provider Foto")
}

// GetProviderPhoto app
func (s *ProviderHandler) GetProviderPhoto(w http.ResponseWriter, r *http.Request) {
	ex := mux.Vars(r)
	id := ex["id"]
	var providerResponse response.Response
	w.Header().Set("content-type", "application/json")
	photo, err := s.providerUsecase.GetPhotoByID(id)
	if err != nil {
		providerResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		response.ResponseWrite(&providerResponse, w)
		log.Println(err)
	} else {
		providerResponse = response.Response{Status: http.StatusAccepted, Message: "Get Provider Foto Success", Data: photo}
		response.ResponseWrite(&providerResponse, w)
	}

	log.Println("Endpoint hit: Get Provider Foto")
}
