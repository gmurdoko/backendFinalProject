package walletrepository

import "finalproject/main/master/model"

//WalletRepository interface for ticket
type WalletRepository interface {
	Payment(Wallet *model.Wallets) error
	Receive(Wallet *model.Wallets) error

	//SelectWallet
	SelectWalletByID(id string) (*model.Wallets, error)

	CheckWalletIDByAssetID(id string) (*string, error)
	CheckWalletIDByUserID(id string) (*string, error)

	//Cek Harga Parkir Perjam
	CheckFeePerHour(id string) (*int, error)
}
