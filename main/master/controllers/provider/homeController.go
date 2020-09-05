package provider

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/provider/providerHomeUsecase"
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
	provider.HandleFunc("/saldo/{id}", providerHandler.GetProviderSaldo).Methods(http.MethodGet)
	provider.HandleFunc("/asset", providerHandler.CreateProviderAsset).Methods(http.MethodPost)
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
	// var assetProviderRequest *models.AssetModel
	// _ = json.NewDecoder(r.Body).Decode(&assetProviderRequest)
	// _, err := ph.providerHomeUsecase.CreateAssetProvider(assetProviderRequest)
	// if err != nil {
	// 	var response response.Response
	// 	response.Status = http.StatusOK
	// 	response.Message = "Success"
	// 	response.Data = "Fail"
	// 	w.Write([]byte("Cannot Add Data"))
	// } else {
	// 	var response response.Response
	// 	response.Status = http.StatusOK
	// 	response.Message = "Success"
	// 	response.Data = assetProviderRequest
	// 	byteData, err := json.Marshal(response)
	// 	if err != nil {
	// 		w.Write([]byte("Something Wrong on Marshalling Data"))
	// 	}
	// 	w.Header().Set("Content-type", "application/json")
	// 	w.Write(byteData)
	// }
}
