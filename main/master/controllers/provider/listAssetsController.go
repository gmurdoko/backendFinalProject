package controllers

import (
	"encoding/json"
	"finalproject/main/master/models"
	providerlistassetsusecase "finalproject/main/master/usecases/provider/listassets"
	"net/http"

	"github.com/gorilla/mux"
)

type ListAssetsHandler struct {
	listAssetsUsecase providerlistassetsusecase.ListAssetsUsecase
}

func ListAssetsController(r *mux.Router, service providerlistassetsusecase.ListAssetsUsecase) {
	listAssetsHandler := ListAssetsHandler{listAssetsUsecase: service}
	listAssets := r.PathPrefix("/providerassets").Subrouter()
	listAssets.HandleFunc("/{id}", listAssetsHandler.getListAssets).Methods(http.MethodGet)
}

func (s *ListAssetsHandler) getListAssets(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	listAssets, err := s.listAssetsUsecase.GetAll(userId)
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil || listAssets == nil {
		response.Response = "Data Not Found"
	} else {
		response.Response = listAssets
	}
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
