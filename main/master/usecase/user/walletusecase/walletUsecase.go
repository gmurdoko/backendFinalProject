package walletusecase

import (
	"finalproject/main/master/model"
)

//WalletUsecase app
type WalletUsecase interface {
	TransactionPayment(Ticket *model.Tickets) error
}
