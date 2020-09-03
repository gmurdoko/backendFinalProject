package controllers

import (
	"encoding/json"
	"finalproject/main/master/models"
	usecases "finalproject/main/master/usecases/review"
	"github.com/gorilla/mux"
	"net/http"
)

type ReviewHandler struct {
	review usecases.ReviewUsecase
}

func ReviewController(r *mux.Router, service usecases.ReviewUsecase) {
	reviewHandler := ReviewHandler{review: service}

	createReview := r.PathPrefix("/providerassets").Subrouter()
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
