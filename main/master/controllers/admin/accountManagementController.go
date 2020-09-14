package admin

import (
	"encoding/json"
	"finalproject/config"
	"finalproject/main/master/models"
	accountmanagementusecase "finalproject/main/master/usecases/admin/accountManagement"
	"finalproject/main/middleware"
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

	isAuthOn := config.AuthSwitch()
	if isAuthOn {
		accountManagement.Use(middleware.TokenValidationMiddleware)
		detailAccountManagerController(accountManagement, accountManagementHandler)
	} else {
		detailAccountManagerController(accountManagement, accountManagementHandler)
	}

}

func detailAccountManagerController(accountManagement *mux.Router, accountManagementHandler AccountManagementControllerHandler) {
	accountManagement.HandleFunc("/deleteuser/{id}", accountManagementHandler.deleteUser).Methods(http.MethodPut)
	accountManagement.HandleFunc("/deleteasset/{id}", accountManagementHandler.deleteAsset).Methods(http.MethodPut)
	accountManagement.HandleFunc("/deleteprovider/{id}", accountManagementHandler.deleteProvider).Methods(http.MethodPut)
	accountManagement.HandleFunc("/deletecomment/{id}", accountManagementHandler.deleteComment).Methods(http.MethodPut)
	accountManagement.HandleFunc("/approveasset/{id}", accountManagementHandler.approveAsset).Methods(http.MethodPut)

	//	Get all users, providers, assets for admin
	accountManagement.HandleFunc("/allusers", accountManagementHandler.GetAllUsers).Methods(http.MethodGet)
	accountManagement.HandleFunc("/allproviders", accountManagementHandler.GetAllProviders).Methods(http.MethodGet)
	accountManagement.HandleFunc("/allassets", accountManagementHandler.GetAllAssets).Methods(http.MethodGet)
	accountManagement.HandleFunc("/allassetsnotapproved", accountManagementHandler.GetAllAssetsNotApproved).Methods(http.MethodGet)
	accountManagement.HandleFunc("/allreviews", accountManagementHandler.GetAllReviews).Methods(http.MethodGet)
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
	resp, err := s.accountManagementUsecase.DeleteAsset(assetId)
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

//	Get all users, providers, assets

func (s *AccountManagementControllerHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	listUsers, err := s.accountManagementUsecase.GetAllUsers()

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil {
		response.Response = "Data Not Found"
	} else {
		response.Response = listUsers
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *AccountManagementControllerHandler) GetAllProviders(w http.ResponseWriter, r *http.Request) {
	listProviders, err := s.accountManagementUsecase.GetAllProviders()

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil {
		response.Response = "Data Not Found"
	} else {
		response.Response = listProviders
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *AccountManagementControllerHandler) GetAllAssets(w http.ResponseWriter, r *http.Request) {
	listAssets, err := s.accountManagementUsecase.GetAllAssets()

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil {
		response.Response = "Data Not Found"
	} else {
		response.Response = listAssets
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *AccountManagementControllerHandler) GetAllAssetsNotApproved(w http.ResponseWriter, r *http.Request) {
	listAssets, err := s.accountManagementUsecase.GetAllAssetsNotApproved()

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil {
		response.Response = "Data Not Found"
	} else {
		response.Response = listAssets
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *AccountManagementControllerHandler) GetAllReviews(w http.ResponseWriter, r *http.Request) {
	listReviews, err := s.accountManagementUsecase.GetAllReviews()

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	if err != nil {
		response.Response = "Data Not Found"
	} else {
		response.Response = listReviews
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
