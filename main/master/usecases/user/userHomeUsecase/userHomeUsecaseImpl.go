package userHomeUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/user/userHomeRepo"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type UserHomeUsecaseImpl struct {
	userRepo userHomeRepo.UserHome
}

func InitUserHomeUsecase(userRepo userHomeRepo.UserHome) UserHome {
	return &UserHomeUsecaseImpl{userRepo: userRepo}
}
func (uc *UserHomeUsecaseImpl) GetSaldo(id string) (*models.UserWallet, error) {
	data, err := uc.userRepo.GetSaldo(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (uc *UserHomeUsecaseImpl) DeleteUserPhoto(id string) error {
	err := uc.userRepo.DeleteUserPhoto(id)
	return err
}
func (uc *UserHomeUsecaseImpl) UpdateUserData(user *models.UserModel, id string) (*models.UserModel, error) {
	data, err := uc.userRepo.UpdateUserData(user, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (uc *UserHomeUsecaseImpl) UpdateUserSaldoTopUp(wallet *models.Wallets, id string) (string, error) {
	data, err := uc.userRepo.UpdateUserSaldoTopUp(wallet, id)
	if err != nil {
		return "0", err
	}
	return data, nil
}
func (uc *UserHomeUsecaseImpl) GetUserPhoto(id string) (*string, error) {
	photo, err := uc.userRepo.GetUserPhoto(id)
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (uc *UserHomeUsecaseImpl) UpdateUserPhoto(photo multipart.File, handler *multipart.FileHeader, id string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	rand.Seed(time.Now().UnixNano())
	min := 11111111111
	max := 99999999999
	photoString := "user-" + strconv.Itoa(rand.Intn(max-min+1)+min) + filepath.Ext(handler.Filename)
	fileLocation := filepath.Join(dir, "files", photoString)

	log.Println(`FileLocation ->`, fileLocation)

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, photo); err != nil {
		log.Println(`Error While Coping File to Local Storage`, err)
		return err
	}

	err = uc.userRepo.UpdateUserPhoto(photoString, id)
	if err != nil {
		return err
	}
	return nil
}
func (uc *UserHomeUsecaseImpl) GetUserTicket(id string) (*models.TicketUser, error) {
	data, err := uc.userRepo.GetUserTicket(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
