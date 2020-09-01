package walletusecase

import (
	"finalproject/main/master/model"
	"finalproject/main/master/repository/walletrepository"
	"math"
	"time"
)

//WalletUsecaseImpl app
type WalletUsecaseImpl struct {
	walletRepository walletrepository.WalletRepository
}

//TransactionPayment app
func (s WalletUsecaseImpl) TransactionPayment(ticket *model.Tickets) error {
	layout := `2006-01-02 15:04:05`
	startAt, err := time.Parse(layout, ticket.StartAt)
	if err != nil {
		return err
	}
	finishAt, err := time.Parse(layout, ticket.FinishedAt)
	if err != nil {
		return err
	}
	HourDifferent := math.Ceil(startAt.Sub(finishAt).Seconds() / 3600)

	userWallet := new(model.Wallets)

	err = s.walletRepository.Payment(userWallet)
	if err != nil {
		return err
	}
	return nil
}

//InitWalletUsecaseImpl app
func InitWalletUsecaseImpl(walletRepository walletrepository.WalletRepository) WalletUsecase {
	return &WalletUsecaseImpl{walletRepository}
}
