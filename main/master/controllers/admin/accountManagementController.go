package admin

import (
	"encoding/json"
	"finalproject/main/master/models"
	accountmanagementusecase "finalproject/main/master/usecases/admin/accountManagement"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountManagementControllerHandler struct {
	accountManagementUsecase accountmanagementusecase.AccountManagementUsecase
}

func AccountManagerController(r *mux.Router, service accountmanagementusecase.AccountManagementUsecase) {
	accountManagementHandler := AccountManagementControllerHandler{accountManagementUsecase: service}
	accountManagement := r.PathPrefix("/accountmanagement").Subrouter()
	accountManagement.HandleFunc("/deleteuser/{id}", accountManagementHandler.deleteUser).Methods(http.MethodPut)
	accountManagement.HandleFunc("/deleteasset/{id}", accountManagementHandler.deleteAsset).Methods(http.MethodPut)
	accountManagement.HandleFunc("/deleteprovider/{id}", accountManagementHandler.deleteProvider).Methods(http.MethodPut)
	accountManagement.HandleFunc("/deletecomment/{id}", accountManagementHandler.deleteComment).Methods(http.MethodPut)
	accountManagement.HandleFunc("/approveasset/{id}", accountManagementHandler.approveAsset).Methods(http.MethodPut)
}

func (s *AccountManagementControllerHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	resp, err := s.accountManagementUsecase.DeleteUser(userId)
	if err != nil {
		log.Println(err)
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Response = resp
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *AccountManagementControllerHandler) deleteAsset(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assetId := params["id"]
	resp, err := s.accountManagementUsecase.DeleteUser(assetId)
	if err != nil {
		log.Println(err)
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Response = resp
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *AccountManagementControllerHandler) deleteProvider(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	provId := params["id"]
	resp, err := s.accountManagementUsecase.DeleteProvider(provId)
	if err != nil {
		log.Println(err)
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Response = resp
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *AccountManagementControllerHandler) deleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewId := params["id"]
	resp, err := s.accountManagementUsecase.DeleteComment(reviewId)
	if err != nil {
		log.Println(err)
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Response = resp
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *AccountManagementControllerHandler) approveAsset(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assetId := params["id"]
	resp, err := s.accountManagementUsecase.ApproveAssets(assetId)
	if err != nil {
		log.Println(err)
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Response = resp
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

// approve asset update masih bingung
