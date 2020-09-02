package walletusecase

import (
	"finalproject/main/master/models"
)

//WalletUsecase app
type WalletUsecase interface {
	TransactionPayment(Ticket *models.Tickets) error
}
