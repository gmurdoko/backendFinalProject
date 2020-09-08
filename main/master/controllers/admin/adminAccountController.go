package admin

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/admin/adminAccountUsecase"
	"finalproject/utils/jwt"
	"net/http"

	"github.com/gorilla/mux"
)

type AdminAccHandler struct {
	adminUsecase adminAccountUsecase.AdminAccountUsecase
}

func AdminAccController(r *mux.Router, service adminAccountUsecase.AdminAccountUsecase) {
	userHandler := AdminAccHandler{adminUsecase: service}

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", userHandler.AdminLogin).Methods(http.MethodPost)
}
func (ac *AdminAccHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {
	var data models.Admin
	_ = json.NewDecoder(r.Body).Decode(&data)

	isValid, _ := ac.adminUsecase.AdminLogin(&data)
	w.Header().Set("Content-type", "application/json")
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
