package adminaccountmanagementusecase

import (
	"finalproject/main/master/models"
	adminaccountmanagementrepo "finalproject/main/master/repositories/admin/accountManagement"
	"log"
)

type AccountManagementUsecaseImpl struct {
	accountManagementRepo adminaccountmanagementrepo.AccountManagementRepo
}

func InitAccountManagementUsecaseImpl(accountManagementRepo adminaccountmanagementrepo.AccountManagementRepo) AccountManagementUsecase {
	return &AccountManagementUsecaseImpl{accountManagementRepo: accountManagementRepo}
}

func (s *AccountManagementUsecaseImpl) DeleteUser(userId string) (string, error) {
	err := s.accountManagementRepo.DeleteUser(userId)
	if err != nil {
		log.Println(err)
		return "Delete User Unsuccessful", err
	}
	return "Delete User Successful", nil
}
func (s *AccountManagementUsecaseImpl) DeleteAsset(assetsId string) (string, error) {
	err := s.accountManagementRepo.DeleteAsset(assetsId)
	if err != nil {
		log.Println(err)
		return "Delete Asset Unsuccessful", err
	}
	return "Delete Asset Successful", nil
}
func (s *AccountManagementUsecaseImpl) DeleteProvider(provId string) (string, error) {
	err := s.accountManagementRepo.DeleteProvider(provId)
	if err != nil {
		log.Println(err)
		return "Delete Provider Unsuccessful", err
	}
	return "Delete Provider Successful", nil
}
func (s *AccountManagementUsecaseImpl) DeleteComment(commentId string) (string, error) {
	err := s.accountManagementRepo.DeleteComment(commentId)
	if err != nil {
		log.Println(err)
		return "Delete Comment Unsuccessful", err
	}
	return "Delete Comment Successful", nil
}
func (s *AccountManagementUsecaseImpl) ApproveAssets(assetID string) (string, error) {
	err := s.accountManagementRepo.ApproveAssets(assetID)
	if err != nil {
		log.Println(err)
		return "Asset Approved", err
	}
	return "Asset Not Approved", nil
}
func (s *AccountManagementUsecaseImpl) ApproveAssetsUpdate(assetId string) (string, error) {
	err := s.accountManagementRepo.ApproveAssetsUpdate(assetId)
	if err != nil {
		log.Println(err)
		return "Update Asset Unsuccessful", err
	}
	return "Update Asset Successful", nil
}

// Get all users, providers, etc.

func (s *AccountManagementUsecaseImpl) GetAllUsers() ([]*models.UserManagement, error) {
	listUsers, err := s.accountManagementRepo.GetAllUsers()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listUsers, err
}

func (s *AccountManagementUsecaseImpl) GetAllProviders() ([]*models.ProvidersManagement, error) {
	listProviders, err := s.accountManagementRepo.GetAllProviders()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listProviders, nil
}

func (s *AccountManagementUsecaseImpl) GetAllAssetsNotApproved() ([]*models.AssetManagement, error) {
	listAssets, err := s.accountManagementRepo.GetAllAssetsNotApproved()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listAssets, nil
}

func (s *AccountManagementUsecaseImpl) GetAllAssets() ([]*models.AssetManagement, error) {
	listAssets, err := s.accountManagementRepo.GetAllAssets()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listAssets, nil
}

func (s *AccountManagementUsecaseImpl) GetAllReviews() ([]*models.ReviewManagement, error) {
	listReviews, err := s.accountManagementRepo.GetAllReviews()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return listReviews, nil
}
