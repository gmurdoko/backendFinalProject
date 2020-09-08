package user

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/user/userHomeUsecase"
	"finalproject/utils/response"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type UserHomeHandler struct {
	userUsecase userHomeUsecase.UserHome
}

func UserHomeController(r *mux.Router, service userHomeUsecase.UserHome) {
	userHandler := UserHomeHandler{userUsecase: service}
	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/{id}", userHandler.UpdateUserData).Methods(http.MethodPut)
	user.HandleFunc("/saldo/{id}", userHandler.GetSaldo).Methods(http.MethodGet)
	user.HandleFunc("/saldo/{id}", userHandler.UpdateUserSaldoTopUp).Methods(http.MethodPut)
	user.HandleFunc("/photo/{id}", userHandler.DeleteUserPhoto).Methods(http.MethodDelete)
	user.HandleFunc("/photo/{id}", userHandler.GetUserPhoto).Methods(http.MethodGet)
	user.HandleFunc("/photo/{id}", userHandler.UpdateUserPhoto).Methods(http.MethodPut)
}
func (uh *UserHomeHandler) GetSaldo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, err := uh.userUsecase.GetSaldo(id)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
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
func (uh *UserHomeHandler) DeleteUserPhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := uh.userUsecase.DeleteUserPhoto(id)
	if err != nil {
		var response response.Response
		response.Status = http.StatusBadRequest
		response.Message = "Success Deleted Data"
		response.Data = nil
		byteData, _ := json.Marshal(response)
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
	}
	var response response.Response
	response.Status = http.StatusOK
	response.Message = "Success Deleted Data"
	response.Data = nil
	byteData, err := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
func (uh *UserHomeHandler) UpdateUserData(w http.ResponseWriter, r *http.Request) {
	var userRequest *models.UserModel
	params := mux.Vars(r)
	id := params["id"]
	_ = json.NewDecoder(r.Body).Decode(&userRequest)
	data, err := uh.userUsecase.UpdateUserData(userRequest, id)
	if err != nil {
		var response response.Response
		response.Status = http.StatusBadGateway
		response.Message = "Fail"
		response.Data = nil
		w.Write([]byte("Cannot Update Data"))
	} else {
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
}
func (uh *UserHomeHandler) UpdateUserSaldoTopUp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var data models.Wallets
	_ = json.NewDecoder(r.Body).Decode(&data)

	saldo, err := uh.userUsecase.UpdateUserSaldoTopUp(&data, id)
	if err != nil {
		var response response.Response
		response.Status = http.StatusBadRequest
		response.Message = "Failed"
		response.Data = nil
		byteData, _ := json.Marshal(response)
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
	}
	var response response.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = saldo
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
func (uh *UserHomeHandler) GetUserPhoto(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	ex := mux.Vars(r)
	id := ex["id"]
	photo, err := uh.userUsecase.GetUserPhoto(id)
	fileLocation := filepath.Join(dir, "files", *photo)
	fmt.Println(fileLocation)
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, fileLocation)
}

func (uh *UserHomeHandler) UpdateUserPhoto(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var providerResponse response.Response
	w.Header().Set("content-type", "application/json")
	r.ParseMultipartForm(1024) // ini untuk batesin file size nya biar maks 1 MB
	photo, handler, err := r.FormFile("photo")
	if err != nil {
		log.Println(`Error while parsing file`, err)
		w.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(message.Respone("Upload Photos Failed", http.StatusInternalServerError, err.Error()))
		// return
	}
	defer photo.Close()
	err = uh.userUsecase.UpdateUserPhoto(photo, handler, id)
	if err != nil {
		providerResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		response.ResponseWrite(&providerResponse, w)
		log.Println(err)
	} else {
		providerResponse = response.Response{Status: http.StatusAccepted, Message: "Update Provider Foto Success", Data: id}
		response.ResponseWrite(&providerResponse, w)
	}

	log.Println("Endpoint hit: Update Provider Foto")
}
