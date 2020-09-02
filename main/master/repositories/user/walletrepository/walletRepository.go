package walletrepository

import "finalproject/main/master/models"

//WalletRepository interface for ticket
type WalletRepository interface {
	Payment(Wallet *models.WalletModel) error
	Receive(Wallet *models.WalletModel) error

	//SelectWallet
	SelectWalletByID(id string) (*models.WalletModel, error)

	CheckWalletIDByAssetID(id string) (*string, error)
	CheckWalletIDByUserID(id string) (*string, error)

	//Cek Harga Parkir Perjam
	CheckFeePerHour(id string) (*int, error)

	//Update ticket to transaction done
	TransactionDone(ticket *models.Tickets) error
}
