package adminaccountmanagementrepo

import (
	"database/sql"
	constanta "finalproject/utils/constant"
	"log"
)

type AccountManagementRepoImpl struct {
	db *sql.DB
}

func InitAccountManagementImpl(mydb *sql.DB) AccountManagementRepo {
	return &AccountManagementRepoImpl{db: mydb}
}

func (s *AccountManagementRepoImpl) DeleteUser(userId string) error {
	query := constanta.DELETEUSER
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = tx.Exec(query, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *AccountManagementRepoImpl) DeleteAsset(assetId string) error {
	query := constanta.DELETEASSET
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = tx.Exec(query, assetId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *AccountManagementRepoImpl) DeleteProvider(providerId string) error {
	query := constanta.DELETEPROVIDER
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = tx.Exec(query, providerId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *AccountManagementRepoImpl) DeleteComment(reviewId string) error {
	query := constanta.DELETECOMMENT
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = tx.Exec(query, reviewId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *AccountManagementRepoImpl) ApproveAssets(assetId string) error {
	query := constanta.APPROVEASSETS
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = tx.Exec(query, assetId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *AccountManagementRepoImpl) ApproveAssetsUpdate(assetId string) error {
	query := constanta.APPROVEASSETS
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = tx.Exec(query, assetId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
