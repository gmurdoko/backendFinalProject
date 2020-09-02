package user

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/user/userHomeUsecase"
	"finalproject/utils/response"
	"log"
	"net/http"

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
	user.HandleFunc("/photo/{id}", userHandler.DeleteUserPhoto).Methods(http.MethodPut)
	user.HandleFunc("/photo/{id}", userHandler.GetUserPhoto).Methods(http.MethodGet)
	// user.HandleFunc("/photo/{id}", userHandler.UpdateUserPhoto).Methods(http.MethodPut)
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
		w.Write([]byte("Delete Data Failed!"))
	}
	var response response.Response
	response.Status = http.StatusOK
	response.Message = "Success Deleted Data"
	byteData, err := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
func (uh *UserHomeHandler) UpdateUserData(w http.ResponseWriter, r *http.Request) {
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
func (uh *UserHomeHandler) UpdateUserSaldoTopUp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var data models.Wallets
	_ = json.NewDecoder(r.Body).Decode(&data)

	saldo, err := uh.userUsecase.UpdateUserSaldoTopUp(&data, id)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Update Data Failed!"))
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

// func (uh *UserHomeHandler) UpdateUserPhoto(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	var data models.UserModel
// 	_ = json.NewDecoder(r.Body).Decode(&data)

// 	photo, err := uh.userUsecase.UpdateUserPhoto(&data, id)
// 	if err != nil {
// 		log.Println(err)
// 		w.Write([]byte("Update Data Failed!"))
// 	}
// 	var response response.Response
// 	response.Status = http.StatusOK
// 	response.Message = "Success"
// 	response.Data = photo
// 	byteData, err := json.Marshal(response)
// 	if err != nil {
// 		w.Write([]byte("Something Wrong on Marshalling Data"))
// 	}
// 	w.Header().Set("Content-type", "application/json")
// 	w.Write(byteData)
// }
