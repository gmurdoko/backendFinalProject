package provider

import (
	"encoding/json"
	"finalproject/config"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/provider/providerHomeUsecase"
	"finalproject/main/middleware"
	"finalproject/utils/response"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ProviderHomeHandler struct {
	providerHomeUsecase providerHomeUsecase.ProviderHome
}

func ProviderHomeController(r *mux.Router, service providerHomeUsecase.ProviderHome) {
	providerHandler := ProviderHomeHandler{providerHomeUsecase: service}
	provider := r.PathPrefix("/provider").Subrouter()

	isAuthOn := config.AuthSwitch()
	if isAuthOn {
		provider.Use(middleware.TokenValidationMiddleware)
		detailProviderHomeController(provider, providerHandler)
	} else {
		detailProviderHomeController(provider, providerHandler)
	}

}

func detailProviderHomeController(provider *mux.Router, providerHandler ProviderHomeHandler) {
	provider.HandleFunc("/saldo/{id}", providerHandler.GetProviderSaldo).Methods(http.MethodGet)
	provider.HandleFunc("/asset", providerHandler.CreateProviderAsset).Methods(http.MethodPost)
	provider.HandleFunc("/asset/review/{id}", providerHandler.GetAssetReview).Methods(http.MethodGet)
}

func (ph *ProviderHomeHandler) GetProviderSaldo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, err := ph.providerHomeUsecase.GetProviderSaldo(id)
	if err != nil {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = 0
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
	var providerResponse response.Response
	w.Header().Set("content-type", "application/json")
	r.ParseMultipartForm(1024) // ini untuk batesin file size nya biar maks 1 MB
	photo, handler, err := r.FormFile("photo")
	if err != nil {
		log.Println(`Error while parsing file`, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer photo.Close()
	var assetProviderRequest *models.AssetModel
	result := r.FormValue("result")
	_ = json.Unmarshal([]byte(result), &assetProviderRequest)
	asset, err := ph.providerHomeUsecase.CreateAssetProvider(photo, handler, assetProviderRequest)
	if err != nil {
		providerResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		response.ResponseWrite(&providerResponse, w)
		log.Println(err)
	} else {
		providerResponse = response.Response{Status: http.StatusAccepted, Message: "Success", Data: asset}
		response.ResponseWrite(&providerResponse, w)
	}

	log.Println("Endpoint hit: Update Provider Foto")

}
func (ph *ProviderHomeHandler) GetAssetReview(w http.ResponseWriter, r *http.Request) {
	var assetResponse response.Response
	params := mux.Vars(r)
	id := params["id"]
	data, err := ph.providerHomeUsecase.GetAssetReview(id)
	w.Header().Set("Content-type", "application/json")
	if err != nil {
		assetResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		response.ResponseWrite(&assetResponse, w)
		log.Println(err)
	} else {
		assetResponse = response.Response{Status: http.StatusOK, Message: "Success", Data: data}
		response.ResponseWrite(&assetResponse, w)
		log.Println(err)
	}
}
