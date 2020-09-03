package reviewusecase

import (
	"finalproject/main/master/models"
	"finalproject/main/master/repositories/user/reviewrepositories"
)

type ReviewUsecaseImpl struct {
	reviewRepo reviewrepositories.ReviewRepo
}

func InitReviewUsecaseImpl(repo reviewrepositories.ReviewRepo) ReviewUsecase {
	return &ReviewUsecaseImpl{reviewRepo: repo}
}

func (r *ReviewUsecaseImpl) CreateReview(review *models.Review) (*models.Review, error) {
	data, err := r.reviewRepo.CreateReview(review)
	if err != nil {
		return nil, err
	}
	return data, nil
}
