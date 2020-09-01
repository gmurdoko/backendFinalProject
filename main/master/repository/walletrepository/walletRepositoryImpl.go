package walletrepository

import (
	"database/sql"
	"finalproject/main/master/model"
	"time"
)

//WalletRepositoryImpl is for init Repository
type WalletRepositoryImpl struct {
	db *sql.DB
}

//Payment app
func (s WalletRepositoryImpl) Payment(Wallet *model.Wallets) error {
	tx, err := s.db.Begin()
	Wallet.EditedAt = time.Now().Format(`2006-01-02 15:04:05`)
	if err != nil {
		return err
	}
	query := "UPDATE m_wallet SET kredit = ?, saldo = ?, edited_at =? WHERE id = ?;"
	_, err = tx.Exec(query, Wallet.Kredit, Wallet.Saldo, Wallet.EditedAt, Wallet.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//Receive app
func (s WalletRepositoryImpl) Receive(Wallet *model.Wallets) error {
	tx, err := s.db.Begin()
	Wallet.EditedAt = time.Now().Format(`2006-01-02 15:04:05`)
	if err != nil {
		return err
	}
	query := "UPDATE m_wallet SET debit = ?, saldo = ?, edited_at =? WHERE id = ?;"
	_, err = tx.Exec(query, Wallet.Kredit, Wallet.Saldo, Wallet.EditedAt, Wallet.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//InitWalletRepositoryImpl is init gate for repository
func InitWalletRepositoryImpl(db *sql.DB) WalletRepository {
	return &WalletRepositoryImpl{db}
}
