package user

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/user/userAccountUsecase"
	"finalproject/utils/jwt"
	"finalproject/utils/response"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserAccHandler struct {
	userUsecase userAccountUsecase.UserAccount
}

func UserAccController(r *mux.Router, service userAccountUsecase.UserAccount) {
	userHandler := UserAccHandler{userUsecase: service}

	auth := r.PathPrefix("/authUser").Subrouter()
	auth.HandleFunc("/login", userHandler.GetUser).Methods(http.MethodPost)
	auth.HandleFunc("/register", userHandler.CreateUsers).Methods(http.MethodPost)
	auth.HandleFunc("/{id}", userHandler.GetUserById).Methods(http.MethodGet)
}
func (uh *UserAccHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var data models.UserModel
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Println("username " + data.Username + ", password " + data.Password + ", email : " + data.Email)

	dataUser, isValid, _ := uh.userUsecase.GetUser(&data)
	w.Header().Set("Content-type", "application/json")
	if isValid {
		// w.WriteHeader(http.StatusOK)
		token, err := jwt.JwtEncoder(data.Username, "rahasiadong")
		if err != nil {
			http.Error(w, "Failed token generation", http.StatusUnauthorized)
		} else {
			var response response.Response
			response.Status = http.StatusOK
			response.Message = "Success"
			response.Token = token
			response.Data = dataUser
			byteData, err := json.Marshal(response)
			if err != nil {
				w.Write([]byte("Something Wrong on Marshalling Data"))
			}
			w.Write(byteData)
		}
	} else {
		var response response.Response
		response.Status = http.StatusBadRequest
		response.Message = "Login User Failed"
		response.Token = ""
		response.Data = nil
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Write(byteData)
	}
}
func (uh *UserAccHandler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	var userRequest *models.UserModel
	_ = json.NewDecoder(r.Body).Decode(&userRequest)
	data, err := uh.userUsecase.CreateUser(userRequest)
	if err != nil {
		var response response.Response
		response.Status = http.StatusBadRequest
		response.Message = "Success"
		response.Data = nil
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
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
func (uh *UserAccHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	user, _ := uh.userUsecase.GetUserById(id)
	// var response models.Response
	token, err := jwt.JwtEncoder(user.Username, "rahasiadong")
	if err != nil {
		http.Error(w, "Failed token generation", http.StatusUnauthorized)
	} else {
		var response response.Response
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Token = token
		response.Data = user
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
	}
}
