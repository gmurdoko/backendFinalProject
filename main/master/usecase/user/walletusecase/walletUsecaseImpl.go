package walletusecase

import (
	"errors"
	"finalproject/main/master/model"
	"finalproject/main/master/repository/user/walletrepository"
	"fmt"
	"math"
	"strconv"
	"time"
)

//walletUsecaseImpl app
type walletUsecaseImpl struct {
	walletRepository walletrepository.WalletRepository
	// feeRepository    feerepository.FeeRepository
}

//TransactionPayment app
func (s walletUsecaseImpl) TransactionPayment(ticket *model.Tickets) error {
	//
	if ticket.Status != "A" {
		return errors.New("Ticket Not Active")
	}
	//TimeDifferent
	layout := `2006-01-02 15:04:05`
	// Update ticket add FinishedAt
	ticket.FinishedAt = time.Now().Format(`2006-01-02 15:04:05`)
	startAt, err := time.Parse(layout, ticket.StartAt)
	if err != nil {
		return err
	}
	finishAt, err := time.Parse(layout, ticket.FinishedAt)
	if err != nil {
		return err
	}
	//HourDifferent
	HourDifferent := int(math.Ceil(finishAt.Sub(startAt).Seconds()) / 3600)
	//Looking for fee per hour
	feePerHour, err := s.walletRepository.CheckFeePerHour(ticket.FeeID)
	if err != nil {
		return err
	}
	// Total Payment
	paymentMoney := *feePerHour * HourDifferent
	fmt.Println("Payment Money:", paymentMoney)
	fmt.Println("Fee Per Hour", feePerHour)
	fmt.Println("HourDifferent:", HourDifferent)
	//Updating Wallet User
	userWalletID, err := s.walletRepository.CheckWalletIDByUserID(ticket.UserID)
	if err != nil {
		return err
	}
	userWallet, err := s.walletRepository.SelectWalletByID(*userWalletID)
	if err != nil {
		return err
	}
	kreditInt, err := strconv.Atoi(userWallet.Kredit)
	if err != nil {
		return err
	}
	fmt.Println("kredit+payment", kreditInt+paymentMoney)
	kreditInt += paymentMoney
	fmt.Println("Kredit", kreditInt)

	kreditFinal := strconv.Itoa(kreditInt)
	userWallet.Kredit = kreditFinal
	saldoInt, err := strconv.Atoi(userWallet.Saldo)
	if err != nil {
		return err
	}
	if saldoInt-paymentMoney < 0 {
		return errors.New("Saldo tidak mencukupi")
	}
	saldoInt -= paymentMoney
	saldoFinal := strconv.Itoa(saldoInt)
	userWallet.Saldo = saldoFinal

	fmt.Println("bisa sampai sini")

	//Updating Wallet Asset
	assetWalletID, err := s.walletRepository.CheckWalletIDByAssetID(ticket.AssetID)
	if err != nil {
		return err
	}

	assetWallet, err := s.walletRepository.SelectWalletByID(*assetWalletID)
	if err != nil {
		return err
	}

	debitInt, err := strconv.Atoi(assetWallet.Debit)
	if err != nil {
		return err
	}
	debitInt += paymentMoney
	debitFinal := strconv.Itoa(debitInt)
	assetWallet.Debit = debitFinal
	saldoInt, err = strconv.Atoi(assetWallet.Saldo)
	if err != nil {
		return err
	}
	saldoInt += paymentMoney
	saldoFinal = strconv.Itoa(saldoInt)
	assetWallet.Saldo = saldoFinal

	//Commit
	err = s.walletRepository.Payment(userWallet)
	if err != nil {
		return err
	}
	err = s.walletRepository.Receive(assetWallet)
	if err != nil {
		return err
	}

	//Update transaction
	err = s.walletRepository.TransactionDone(ticket)
	if err != nil {
		return err
	}

	return nil
}

//InitWalletUsecaseImpl app
func InitWalletUsecaseImpl(walletRepository walletrepository.WalletRepository) WalletUsecase {
	return &walletUsecaseImpl{walletRepository}
}
