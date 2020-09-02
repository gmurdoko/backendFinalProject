package controller

import (
	"encoding/json"
	"finalproject/config"
	"finalproject/main/master/model"
	"finalproject/main/master/usecase/user/walletusecase"
	"finalproject/main/middleware"
	"finalproject/utils/response"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//WalletHandler app
type WalletHandler struct {
	walletUsecase walletusecase.WalletUsecase
}

//WalletController app
func WalletController(r *mux.Router, s walletusecase.WalletUsecase) {
	walletHandler := WalletHandler{s}
	wallets := r.PathPrefix("/wallets").Subrouter()
	wallet := r.PathPrefix("/wallet").Subrouter()
	isAuthOn := config.AuthSwitch()
	if isAuthOn {
		wallets.Use(middleware.TokenValidationMiddleware)
		wallet.Use(middleware.TokenValidationMiddleware)
		detailWalletController(wallets, wallet, walletHandler)
	} else {
		detailWalletController(wallets, wallet, walletHandler)
	}

}

func detailWalletController(wallets, wallet *mux.Router, walletHandler WalletHandler) {
	//Jamak
	// wallets.HandleFunc("", walletHandler.Listwallets).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}", "status", "{status}", "orderBy", "{orderBy}", "sort", "{sort}").Methods(http.MethodGet)
	// wallets.HandleFunc("/available", walletHandler.ListAvailablewallets).Methods(http.MethodGet)
	// wallets.HandleFunc("/booked", walletHandler.ListBookedwallets).Methods(http.MethodGet)
	//Satuan
	// wallet.HandleFunc("/{id}", walletHandler.wallet).Methods(http.MethodGet)
	// wallet.HandleFunc("", walletHandler.Postwallet).Methods(http.MethodPost)
	// wallet.HandleFunc("", walletHandler.Putwallet).Methods(http.MethodPut)
	// wallet.HandleFunc("/{id}", walletHandler.Deletewallet).Methods(http.MethodDelete)
	wallet.HandleFunc("/payment", walletHandler.TransactionPayment).Methods(http.MethodPost)
}

//TransactionPayment app
func (s *WalletHandler) TransactionPayment(w http.ResponseWriter, r *http.Request) {
	var inTicket model.Tickets
	var transactionPaymentResponse response.Response
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&inTicket)
	if err != nil {
		log.Println(err)
		// w.Write([]byte("cant read JSON"))
	}
	err = s.walletUsecase.TransactionPayment(&inTicket)
	if err != nil {
		transactionPaymentResponse = response.Response{Status: http.StatusBadRequest, Message: "Error", Data: err.Error()}
		response.ResponseWrite(&transactionPaymentResponse, w)
		log.Println(err)
	} else {
		transactionPaymentResponse = response.Response{Status: http.StatusAccepted, Message: "Post Transaction Payment Success", Data: inTicket}
		response.ResponseWrite(&transactionPaymentResponse, w)
	}
	log.Println("Endpoint hit: Post Transaction Payment")
}
