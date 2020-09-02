package controllers

import (
	"finalproject/config"
	"finalproject/main/master/usecase/user/ticketusecase"
	"finalproject/main/middleware"
	"finalproject/utils/response"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//TicketHandler app
type TicketHandler struct {
	ticketUsecase ticketusecase.TicketUsecase
}

//TicketController app
func TicketController(r *mux.Router, s ticketusecase.TicketUsecase) {
	ticketHandler := TicketHandler{s}
	tickets := r.PathPrefix("/tickets").Subrouter()
	ticket := r.PathPrefix("/ticket").Subrouter()
	isAuthOn := config.AuthSwitch()
	if isAuthOn {
		tickets.Use(middleware.TokenValidationMiddleware)
		ticket.Use(middleware.TokenValidationMiddleware)
		detailTicketController(tickets, ticket, ticketHandler)
	} else {
		detailTicketController(tickets, ticket, ticketHandler)
	}
}

func detailTicketController(tickets, ticket *mux.Router, ticketHandler TicketHandler) {
	//Jamak
	// tickets.HandleFunc("", ticketHandler.Listtickets).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}", "status", "{status}", "orderBy", "{orderBy}", "sort", "{sort}").Methods(http.MethodGet)
	// tickets.HandleFunc("/available", ticketHandler.ListAvailabletickets).Methods(http.MethodGet)
	// tickets.HandleFunc("/booked", ticketHandler.ListBookedtickets).Methods(http.MethodGet)
	//Satuan
	// ticket.HandleFunc("/{id}", ticketHandler.ticket).Methods(http.MethodGet)
	// ticket.HandleFunc("", ticketHandler.Postticket).Methods(http.MethodPost)
	// ticket.HandleFunc("", ticketHandler.Putticket).Methods(http.MethodPut)
	ticket.HandleFunc("/{id}", ticketHandler.DeleteTicket).Methods(http.MethodDelete)
}

// DeleteTicket app
func (s *TicketHandler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	ex := mux.Vars(r)
	id := ex["id"]
	var ticketResponse response.Response
	w.Header().Set("content-type", "application/json")
	err := s.ticketUsecase.DeleteTicket(id)
	if err != nil {
		ticketResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		response.ResponseWrite(&ticketResponse, w)
		log.Println(err)
	} else {
		ticketResponse = response.Response{Status: http.StatusAccepted, Message: "Delete Ticket Success", Data: id}
		response.ResponseWrite(&ticketResponse, w)
	}

	log.Println("Endpoint hit: Delete Ticket")
}
