package providerHomeUsecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/provider/providerHomeRepo"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type ProviderHomeUsecaseImpl struct {
	providerRepo providerHomeRepo.ProviderHome
}

func InitProviderHomeUsecase(providerRepo providerHomeRepo.ProviderHome) ProviderHome {
	return &ProviderHomeUsecaseImpl{providerRepo: providerRepo}
}
func (pu *ProviderHomeUsecaseImpl) CreateAssetProvider(photo multipart.File, handler *multipart.FileHeader, provider *models.AssetModel) (*models.AssetModel, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	min := 11111111111
	max := 99999999999
	photoString := "asset-" + strconv.Itoa(rand.Intn(max-min+1)+min) + filepath.Ext(handler.Filename)
	fileLocation := filepath.Join(dir, "files", photoString)

	log.Println(`FileLocation ->`, fileLocation)

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, photo); err != nil {
		log.Println(`Error While Coping File to Local Storage`, err)
		return nil, err
	}

	provider.Photo = photoString
	data, err := pu.providerRepo.CreateAssetProvider(provider)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (pu *ProviderHomeUsecaseImpl) GetProviderSaldo(id string) (string, error) {
	data, err := pu.providerRepo.GetProviderSaldo(id)
	if err != nil {
		return "0", err
	}
	return data, nil
}
func (pu *ProviderHomeUsecaseImpl) GetAssetReview(id string) ([]*models.AssetReview, error) {
	data, err := pu.providerRepo.GetAssetReview(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
