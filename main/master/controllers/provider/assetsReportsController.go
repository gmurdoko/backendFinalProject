package provider

import (
	"encoding/json"
	"finalproject/config"
	"finalproject/main/master/usecases/provider/providerAssetReportsUsecase"
	"finalproject/main/middleware"
	"finalproject/utils/response"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ProviderAssetReportHandler struct {
	assetsReport providerAssetReportsUsecase.ProviderAssetReportsUsecase
}

func ProviderAssetReportController(r *mux.Router, service providerAssetReportsUsecase.ProviderAssetReportsUsecase) {
	assetsReportHandler := ProviderAssetReportHandler{assetsReport: service}
	reportAsset := r.PathPrefix("/providerreports").Subrouter()
	isAuthOn := config.AuthSwitch()
	if isAuthOn {
		reportAsset.Use(middleware.TokenValidationMiddleware)
		detailProviderAssetReportController(reportAsset, assetsReportHandler)
	} else {
		detailProviderAssetReportController(reportAsset, assetsReportHandler)
	}

}

func detailProviderAssetReportController(reportAsset *mux.Router, assetsReportHandler ProviderAssetReportHandler) {
	reportAsset.HandleFunc("/daily", assetsReportHandler.getReportDaily).Queries("id", "{id}").Methods(http.MethodGet)
	reportAsset.HandleFunc("/monthly", assetsReportHandler.getReportMonthly).Queries("id", "{id}").Methods(http.MethodGet)
}

func (s *ProviderAssetReportHandler) getReportDaily(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assetId := params["id"]
	reports, err := s.assetsReport.GetReportDaily(assetId)
	var response response.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil || reports == nil {
		response.Data = nil
	} else {
		response.Data = reports
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
	var response response.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	fmt.Println(reports)
	if err != nil || reports == nil {
		response.Data = nil
	} else {
		response.Data = reports
	}
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
