package review_repositories

import "finalproject/main/master/models"

type ReviewRepo interface {
	CreateReview(review *models.Review) (*models.Review, error)
}



