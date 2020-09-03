package reviewrepositories

import "finalproject/main/master/models"

type ReviewRepo interface {
	CreateReview(review *models.Review) (*models.Review, error)
}
