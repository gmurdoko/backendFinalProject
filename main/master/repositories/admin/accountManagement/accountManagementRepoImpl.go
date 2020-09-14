package adminaccountmanagementrepo

import (
	"database/sql"
	"finalproject/main/master/models"
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
		log.Println(err)
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

// Get all users, providers, etc.

func (s *AccountManagementRepoImpl) GetAllUsers() ([]*models.UserManagement, error) {
	query := constanta.GET_ALL_USERS
	var listUsers []*models.UserManagement
	var bd sql.NullString

	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.UserManagement{}
		err := rows.Scan(&user.ID, &user.IdWallet, &user.Username,
			&user.Email, &user.Fullname, &bd, &user.PhoneNumber,
			&user.Address, &user.CreatedAt, &user.Status)
		user.BornDate = bd.String
		if err != nil {
			log.Println(err)
			return nil, err
		}
		listUsers = append(listUsers, &user)
	}
	return listUsers, nil
}

func (s *AccountManagementRepoImpl) GetAllProviders() ([]*models.ProvidersManagement, error) {
	query := constanta.GET_ALL_PROVIDERS
	var listProviders []*models.ProvidersManagement
	var bd sql.NullString
	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		provider := models.ProvidersManagement{}
		err := rows.Scan(&provider.ID, &provider.Username, &provider.Email,
			&provider.Fullname, &bd, &provider.PhoneNumber,
			&provider.Address, &provider.CreatedAt, &provider.Status)
		provider.Borndate = bd.String
		if err != nil {
			log.Println(err)
			return nil, err
		}
		listProviders = append(listProviders, &provider)
	}
	return listProviders, nil
}

func (s *AccountManagementRepoImpl) GetAllAssetsNotApproved() ([]*models.AssetManagement, error) {
	query := constanta.GET_ALL_ASSET_NOT_APPROVED
	var listAssets []*models.AssetManagement

	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		asset := models.AssetManagement{}
		err := rows.Scan(&asset.ID, &asset.IdWallet, &asset.ProviderId,
			&asset.AssetName, &asset.AssetArea, &asset.Longitude,
			&asset.Latitude, &asset.CarCapacity, &asset.MotorcycleCapacity,
			&asset.BicycleCapacity, &asset.CreatedAt, &asset.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		listAssets = append(listAssets, &asset)
	}
	return listAssets, nil
}

func (s *AccountManagementRepoImpl) GetAllAssets() ([]*models.AssetManagement, error) {
	query := constanta.GET_ALL_ASSETS_APPROVED
	var listAssets []*models.AssetManagement

	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		asset := models.AssetManagement{}
		err := rows.Scan(&asset.ID, &asset.IdWallet, &asset.ProviderId,
			&asset.AssetName, &asset.AssetArea, &asset.Longitude,
			&asset.Latitude, &asset.CarCapacity, &asset.MotorcycleCapacity,
			&asset.BicycleCapacity, &asset.CreatedAt, &asset.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		listAssets = append(listAssets, &asset)
	}
	return listAssets, nil
}

func (s *AccountManagementRepoImpl) GetAllReviews() ([]*models.ReviewManagement, error) {
	query := constanta.GET_ALL_REVIEWS
	var listReviews []*models.ReviewManagement

	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		review := models.ReviewManagement{}
		err := rows.Scan(&review.ID, &review.UserID, &review.AssetID,
			&review.Rating, &review.Comment, &review.CreatedAt,
			&review.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		listReviews = append(listReviews, &review)
	}
	return listReviews, nil
}
