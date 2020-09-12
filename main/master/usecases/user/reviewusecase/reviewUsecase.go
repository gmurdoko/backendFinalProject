package reviewusecase

import (
	"finalproject/main/master/models"
)

type ReviewUsecase interface {
	CreateReview(review *models.Review) (*models.Review, error)
	GetStatusReview(user_id, asset_id string) error
}
