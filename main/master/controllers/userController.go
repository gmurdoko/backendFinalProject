package controllers

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/userUsecase"
	"finalproject/utils/jwt"
	"finalproject/utils/response"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUsecase userUsecase.UserUsecase
}

func UserController(r *mux.Router, service userUsecase.UserUsecase) {
	userHandler := UserHandler{userUsecase: service}
	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("", userHandler.CreateUsers).Methods(http.MethodPost)
	user.HandleFunc("", userHandler.UpdateUserData).Methods(http.MethodPut)
	user.HandleFunc("/saldo/{id}", userHandler.GetSaldo).Methods(http.MethodGet)
	user.HandleFunc("/saldo/{id}", userHandler.UpdateUserSaldoTopUp).Methods(http.MethodPut)
	user.HandleFunc("/photo/{id}", userHandler.DeleteUserPhoto).Methods(http.MethodDelete)
	user.HandleFunc("/photo/{id}", userHandler.GetUserPhoto).Methods(http.MethodGet)
	user.HandleFunc("/photo/{id}", userHandler.DeleteUserPhoto).Methods(http.MethodPut)
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("", userHandler.GetUser).Methods(http.MethodPost)
}
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var data models.UserModel
	_ = json.NewDecoder(r.Body).Decode(&data)
	isValid, _ := uh.userUsecase.GetUser(&data)

	if isValid {
		w.WriteHeader(http.StatusOK)
		token, err := jwt.JwtEncoder(data.Username, "rahasiadong")
		if err != nil {
			http.Error(w, "Failed token generation", http.StatusUnauthorized)
		} else {
			w.Write([]byte(token))
		}
	} else {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
	}
}
func (uh *UserHandler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	var userRequest *models.UserModel
	_ = json.NewDecoder(r.Body).Decode(&userRequest)
	_, err := uh.userUsecase.CreateUser(userRequest)
	if err != nil {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = "Fail"
		w.Write([]byte("Cannot Add Data"))
	} else {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = userRequest
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
	}
}
func (uh *UserHandler) GetSaldo(w http.ResponseWriter, r *http.Request) {
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
func (uh *UserHandler) DeleteUserPhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := uh.userUsecase.DeleteUserPhoto(id)
	if err != nil {
		w.Write([]byte("Delete Data Failed!"))
	}
	var response response.Response
	response.Status = http.StatusOK
	response.Message = "Success Deleted Data"
	byteData, err := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
func (uh *UserHandler) UpdateUserData(w http.ResponseWriter, r *http.Request) {
	var userRequest *models.UserModel
	params := mux.Vars(r)
	id := params["id"]
	_ = json.NewDecoder(r.Body).Decode(&userRequest)
	_, err := uh.userUsecase.UpdateUserData(userRequest, id)
	if err != nil {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = "Fail"
		w.Write([]byte("Cannot Update Data"))
	} else {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = userRequest
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
	}
}
func (uh *UserHandler) UpdateUserSaldoTopUp(w http.ResponseWriter, r *http.Request) {
	var userRequest *models.WalletModel
	params := mux.Vars(r)
	id := params["id"]
	_ = json.NewDecoder(r.Body).Decode(&userRequest)
	saldo, err := uh.userUsecase.UpdateUserSaldoTopUp(userRequest, id)
	if err != nil {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = "Fail"
		w.Write([]byte("Cannot Update Data"))
	} else {
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
}
func (uh *UserHandler) GetUserPhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, err := uh.userUsecase.GetUserPhoto(id)
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
func (uh *UserHandler) UpdateUserPhoto(w http.ResponseWriter, r *http.Request) {
	var userRequest *models.UserModel
	params := mux.Vars(r)
	id := params["id"]
	_ = json.NewDecoder(r.Body).Decode(&userRequest)
	photo, err := uh.userUsecase.UpdateUserPhoto(userRequest, id)
	if err != nil {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = "Fail"
		w.Write([]byte("Cannot Update Data"))
	} else {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = photo
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
	}
}
