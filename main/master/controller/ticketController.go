package controller

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

	tickets.HandleFunc("/history", ticketHandler.HistoryTickets).Queries("page", "{page}", "limit", "{limit}", "user_id", "{user_id}").Methods(http.MethodGet)

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

//HistoryTickets app
func (s *TicketHandler) HistoryTickets(w http.ResponseWriter, r *http.Request) {
	offset := mux.Vars(r)["page"]
	limit := mux.Vars(r)["limit"]
	id := mux.Vars(r)["user_id"]

	println("INI CONTROLLER", offset, limit, id)
	tickets, totalField, err := s.ticketUsecase.GetHistoryTicketByID(offset, limit, id)
	var ticketsResponse response.Response
	w.Header().Set("content-type", "application/json")
	if err != nil {
		ticketsResponse = response.Response{Status: http.StatusNotFound, Message: "Not Found", TotalField: *totalField, Data: err.Error()}
		response.ResponseWrite(&ticketsResponse, w)
		log.Println(err)
	} else {
		ticketsResponse = response.Response{Status: http.StatusOK, Message: "Get History Tickets Success", TotalField: *totalField, Data: tickets}
		response.ResponseWrite(&ticketsResponse, w)
	}
	log.Println("Endpoint hit: Get History Tickets")

}
