package admin

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/admin/adminReportUsecase"
	"finalproject/main/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AdminAssetReportHandler struct {
	assetsReport adminReportUsecase.AdminAssetReportsUsecase
}

func AdminAssetReportController(r *mux.Router, service adminReportUsecase.AdminAssetReportsUsecase) {
	assetsReportHandler := AdminAssetReportHandler{assetsReport: service}
	reportAsset := r.PathPrefix("/adminreports").Subrouter()
	isAuthOn := false
	if isAuthOn {
		reportAsset.Use(middleware.TokenValidationMiddleware)
		detailAdminAssetReportController(reportAsset, assetsReportHandler)
	} else {
		detailAdminAssetReportController(reportAsset, assetsReportHandler)
	}

}

func detailAdminAssetReportController(reportAsset *mux.Router, assetsReportHandler AdminAssetReportHandler) {
	reportAsset.HandleFunc("/daily", assetsReportHandler.getReportDaily).Queries("start", "{start}", "end", "{end}", "id", "{id}").Methods(http.MethodGet)
	// reportAsset.HandleFunc("/monthly/{id}", assetsReportHandler.getReportMonthly).Methods(http.MethodGet)
}

func (s *AdminAssetReportHandler) getReportDaily(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	start := params["start"]
	end := params["end"]
	assetId := params["id"]
	reports, err := s.assetsReport.GetReportDaily(start, end, assetId)
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil {
		log.Println("error")
		response.Response = "Data Not Found"
	} else {
		if reports != nil {
			response.Response = reports
		} else {
			response.Response = []string{}
		}
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
		if reports != nil {
			response.Response = reports
		} else {
			response.Response = []string{}
		}
	}
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
