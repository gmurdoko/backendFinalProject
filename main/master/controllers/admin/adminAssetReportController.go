package admin

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/admin/adminReportUsecase"
	"net/http"

	"github.com/gorilla/mux"
)

type AdminAssetReportHandler struct {
	assetsReport adminReportUsecase.AdminAssetReportsUsecase
}

func AdminAssetReportController(r *mux.Router, service adminReportUsecase.AdminAssetReportsUsecase) {
	assetsReportHandler := AdminAssetReportHandler{assetsReport: service}
	reportAsset := r.PathPrefix("/adminreports").Subrouter()
	reportAsset.HandleFunc("/daily/{id}", assetsReportHandler.getReportDaily).Methods(http.MethodGet)
	reportAsset.HandleFunc("/monthly/{id}", assetsReportHandler.getReportMonthly).Methods(http.MethodGet)
}

func (s *AdminAssetReportHandler) getReportDaily(w http.ResponseWriter, r *http.Request) {
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

func (s *AdminAssetReportHandler) getReportMonthly(w http.ResponseWriter, r *http.Request) {
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
