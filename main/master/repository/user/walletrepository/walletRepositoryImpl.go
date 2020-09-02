package walletrepository

import (
	"database/sql"
	"finalproject/main/master/model"
	"time"
)

//walletRepositoryImpl is for init Repository
type walletRepositoryImpl struct {
	db *sql.DB
}

//Payment app
func (s walletRepositoryImpl) Payment(Wallet *model.Wallets) error {
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
func (s walletRepositoryImpl) Receive(Wallet *model.Wallets) error {
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

//SelectWalletByID app
func (s walletRepositoryImpl) SelectWalletByID(id string) (*model.Wallets, error) {
	var wallet = new(model.Wallets)
	err := s.db.QueryRow("SELECT * FROM m_wallet WHERE id = ?", id).Scan(&wallet.ID, &wallet.Saldo, &wallet.Debit, &wallet.Kredit, &wallet.CreatedAt, &wallet.EditedAt, &wallet.DeletedAt, &wallet.Status)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

//CheckWalletIDByAssetID For Checking Amount Fee By ID
func (s walletRepositoryImpl) CheckWalletIDByAssetID(id string) (*string, error) {
	var idWallet = new(string)
	err := s.db.QueryRow("SELECT id_wallet FROM m_asset WHERE id = ?", id).Scan(idWallet)
	if err != nil {
		return nil, err
	}
	return idWallet, nil
}

//CheckWalletIDByUserID For Checking Amount Fee By ID
func (s walletRepositoryImpl) CheckWalletIDByUserID(id string) (*string, error) {
	var idWallet = new(string)
	err := s.db.QueryRow("SELECT id_wallet FROM m_user_account WHERE id = ?", id).Scan(idWallet)
	if err != nil {
		return nil, err
	}
	return idWallet, nil
}

//CheckFeePerHour For Checking Amount Fee By ID
func (s walletRepositoryImpl) CheckFeePerHour(id string) (*int, error) {
	var feePerHour = new(int)
	err := s.db.QueryRow("SELECT fee FROM m_fee WHERE id = ?", id).Scan(feePerHour)
	if err != nil {
		return nil, err
	}
	return feePerHour, nil
}

//InitWalletRepositoryImpl is init gate for repository
func InitWalletRepositoryImpl(db *sql.DB) WalletRepository {
	return &walletRepositoryImpl{db}
}
