package controllers

import (
	"finalproject/config"
	"finalproject/main/master/usecase/provider/providerusecase"
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
	provider.HandleFunc("/{id}", providerHandler.DeleteProviderPhoto).Methods(http.MethodDelete)
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
		providerResponse = response.Response{Status: http.StatusAccepted, Message: "Delete Ticket Success", Data: id}
		response.ResponseWrite(&providerResponse, w)
	}

	log.Println("Endpoint hit: Delete Ticket")
}
