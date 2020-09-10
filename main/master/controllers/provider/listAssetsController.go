package provider

import (
	"encoding/json"
	"finalproject/main/master/usecases/provider/providerListAssetUsecase"
	"finalproject/utils/response"
	"net/http"

	"github.com/gorilla/mux"
)

type ListAssetsHandler struct {
	listAssetsUsecase providerListAssetUsecase.ListAssetsUsecase
}

func ListAssetsController(r *mux.Router, service providerListAssetUsecase.ListAssetsUsecase) {
	listAssetsHandler := ListAssetsHandler{listAssetsUsecase: service}
	listAssets := r.PathPrefix("/providerassets").Subrouter()
	listAssets.HandleFunc("/{id}", listAssetsHandler.getListAssets).Methods(http.MethodGet)
}

func (s *ListAssetsHandler) getListAssets(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	listAssets, err := s.listAssetsUsecase.GetAll(userId)
	var response response.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil || listAssets == nil {
		response.Data = nil
	} else {
		response.Data = listAssets
	}
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
