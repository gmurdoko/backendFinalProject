package user

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/user/reviewusecase"
	"finalproject/main/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type ReviewHandler struct {
	review reviewusecase.ReviewUsecase
}

func ReviewController(r *mux.Router, service reviewusecase.ReviewUsecase) {
	reviewHandler := ReviewHandler{review: service}
	r.Use(middleware.ActivityLogMiddleware)

	createReview := r.PathPrefix("/review").Subrouter()
	createReview.HandleFunc("/review", reviewHandler.CreateReview).Methods(http.MethodPost)
}

func (s *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var review *models.Review
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"

	_ = json.NewDecoder(r.Body).Decode(&review)
	data, err := s.review.CreateReview(review)
	if err != nil {
		response.Response = "Cannot Add Data"
	} else {
		response.Response = data
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)

}
