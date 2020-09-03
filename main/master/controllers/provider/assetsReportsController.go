package provider

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/provider/providerAssetReportsUsecase"
	"net/http"

	"github.com/gorilla/mux"
)

type ProviderAssetReportHandler struct {
	assetsReport providerAssetReportsUsecase.ProviderAssetReportsUsecase
}

func ProviderAssetReportController(r *mux.Router, service providerAssetReportsUsecase.ProviderAssetReportsUsecase) {
	assetsReportHandler := ProviderAssetReportHandler{assetsReport: service}
	reportAsset := r.PathPrefix("/providerreports").Subrouter()
	reportAsset.HandleFunc("/daily", assetsReportHandler.getReportDaily).Queries("id", "{id}").Methods(http.MethodGet)
	reportAsset.HandleFunc("/monthly", assetsReportHandler.getReportMonthly).Queries("id", "{id}").Methods(http.MethodGet)
}

func (s *ProviderAssetReportHandler) getReportDaily(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assetId := params["id"]
	reports, err := s.assetsReport.GetReportDaily(assetId)
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil || reports == nil {
		response.Response = "Data Not Found"
	} else {
		response.Response = reports
	}
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *ProviderAssetReportHandler) getReportMonthly(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assetId := params["id"]
	reports, err := s.assetsReport.GetReportMonthly(assetId)
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil || reports == nil {
		response.Response = "Data Not Found"
	} else {
		response.Response = reports
	}
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
