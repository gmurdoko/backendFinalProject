package walletrepository

import "finalproject/main/master/model"

//WalletRepository interface for ticket
type WalletRepository interface {
	Payment(Wallet *model.Wallets) error
	Receive(Wallet *model.Wallets) error
}
