package controllers

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/userUsecase"
	"finalproject/utils/response"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUsecase userUsecase.UserUsecase
}

func UserController(r *mux.Router, service userUsecase.UserUsecase) {
	userHandler := UserHandler{userUsecase: service}
	r.HandleFunc("/user", userHandler.CreateUsers).Methods(http.MethodPost)
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
