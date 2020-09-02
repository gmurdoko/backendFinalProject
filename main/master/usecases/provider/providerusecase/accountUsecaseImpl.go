package providerusecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/provider/providerrepository"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//providerUsecaseImpl app
type providerUsecaseImpl struct {
	providerRepository providerrepository.ProviderRepository
}

//DeleteProvider app
func (s providerUsecaseImpl) DeleteProviderFoto(id string) error {
	err := s.providerRepository.DeletePhotoProvider(id)
	if err != nil {
		return err
	}
	return nil
}

func (s providerUsecaseImpl) UpdateDataProvider(Provider *models.Providers) error {
	err := s.providerRepository.UpdateDataProvider(Provider)
	if err != nil {
		return err
	}
	return nil
}

func (s providerUsecaseImpl) UpdateProviderFoto(photo multipart.File, handler *multipart.FileHeader, id string) error {

	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	rand.Seed(time.Now().UnixNano())
	min := 11111111111
	max := 99999999999
	photoString := "provider-" + strconv.Itoa(rand.Intn(max-min+1)+min) + filepath.Ext(handler.Filename)
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

	err = s.providerRepository.UpdatePhotoProvider(photoString, id)
	if err != nil {
		return err
	}
	return nil
}

func (s providerUsecaseImpl) GetPhotoByID(id string) (*string, error) {
	photo, err := s.providerRepository.SelectPhotoByID(id)
	if err != nil {
		return nil, err
	}
	return photo, nil
}

//InitProviderUsecaseImpl app
func InitProviderUsecaseImpl(providerRepository providerrepository.ProviderRepository) ProviderUsecase {
	return &providerUsecaseImpl{providerRepository}
}
