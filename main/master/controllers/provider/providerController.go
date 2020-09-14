package provider

import (
	"encoding/json"
	"finalproject/config"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/provider/providerusecase"
	"finalproject/main/middleware"
	"finalproject/utils/response"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	provider.HandleFunc("/photo/{id}", providerHandler.GetProviderPhoto).Methods(http.MethodGet)
	provider.HandleFunc("/photo/{id}", providerHandler.UpdateProviderPhoto).Methods(http.MethodPut)
	provider.HandleFunc("/photo/{id}", providerHandler.DeleteProviderPhoto).Methods(http.MethodDelete)
	provider.HandleFunc("/data/{id}", providerHandler.PutDataProvider).Methods(http.MethodPut)

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
	data, err := s.providerUsecase.UpdateDataProvider(id, &inProvider)
	if err != nil {
		providerResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		response.ResponseWrite(&providerResponse, w)
		log.Println(err)
	} else {
		providerResponse = response.Response{Status: http.StatusOK, Message: "Update Data Provider Success", Data: data}
		response.ResponseWrite(&providerResponse, w)
	}

	log.Println("Endpoint hit: Update Data Provider")
}

//UpdateProviderPhoto app
func (s *ProviderHandler) UpdateProviderPhoto(w http.ResponseWriter, r *http.Request) {
	// photo := mux.Vars(r)["photo"]
	id := r.FormValue("id")
	var providerResponse response.Response
	w.Header().Set("content-type", "application/json")
	r.ParseMultipartForm(1024) // ini untuk batesin file size nya biar maks 1 MB
	photo, handler, err := r.FormFile("photo")
	if err != nil {
		log.Println(`Error while parsing file`, err)
		w.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(message.Respone("Upload Photos Failed", http.StatusInternalServerError, err.Error()))
		// return
	}
	defer photo.Close()
	err = s.providerUsecase.UpdateProviderFoto(photo, handler, id)
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
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	ex := mux.Vars(r)
	id := ex["id"]
	photo, err := s.providerUsecase.GetPhotoByID(id)
	fileLocation := filepath.Join(dir, "files", *photo)

	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, fileLocation)
	// var providerResponse response.Response
	// w.Header().Set("content-type", "application/json")
	// photo, err := s.providerUsecase.GetPhotoByID(id)
	// if err != nil {
	// 	providerResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
	// 	response.ResponseWrite(&providerResponse, w)
	// 	log.Println(err)
	// } else {
	// 	providerResponse = response.Response{Status: http.StatusAccepted, Message: "Get Provider Foto Success", Data: photo}
	// 	response.ResponseWrite(&providerResponse, w)
	// }

	// log.Println("Endpoint hit: Get Provider Foto")
}
