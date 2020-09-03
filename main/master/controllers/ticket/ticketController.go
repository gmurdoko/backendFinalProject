package ticket_controllers

import (
	"encoding/json"
	"finalproject/main/master/models"
	"finalproject/main/master/usecases/ticket"
	"finalproject/main/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TicketHandler struct {
	ticket ticket_usecases.TicketUsecase
}

func TicketController(r *mux.Router, service ticket_usecases.TicketUsecase) {
	ticketHandler := TicketHandler{ticket: service}
	r.Use(middleware.ActivityLogMiddleware)

	ticket := r.PathPrefix("/ticket").Subrouter()
	ticket.HandleFunc("/new", ticketHandler.CreateTicket).Methods(http.MethodPost)
	ticket.HandleFunc("/active/{id}", ticketHandler.SetTicketActive).Methods(http.MethodPut)
	ticket.HandleFunc("/inactive/{id}", ticketHandler.SetTicketInactive).Methods(http.MethodPut)
}

func (s *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var newTicket *models.Ticket
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"

	_ = json.NewDecoder(r.Body).Decode(&newTicket)
	data, err := s.ticket.CreateNewTicket(newTicket)
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

func (s *TicketHandler) SetTicketActive(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketId := params["id"]
	resp, err := s.ticket.UpdateTicketStatusActive(ticketId)
	if err != nil {
		log.Println(err)
	}

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Response = resp
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)

}

func (s *TicketHandler) SetTicketInactive(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketId := params["id"]
	resp, err := s.ticket.UpdateTicketStatusInactive(ticketId)
	if err != nil {
		log.Println(err)
	}

	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Response = resp
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something went wrong when marshaling data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)

}
