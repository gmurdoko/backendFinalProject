package review_usecases

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/review"
)

type ReviewUsecaseImpl struct {
	reviewRepo review_repositories.ReviewRepo
}

func InitReviewUsecaseImpl(repo review_repositories.ReviewRepo) ReviewUsecase {
	return &ReviewUsecaseImpl{reviewRepo: repo}
}

func (r *ReviewUsecaseImpl) CreateReview(review *models.Review) (*models.Review, error) {
	data, err := r.reviewRepo.CreateReview(review)
	if err != nil {
		return nil, err
	}
	return data, nil
}
