package review

import (
	"database/sql"
	"finalproject/main/master/models"
	"finalproject/utils/constant"
	"github.com/google/uuid"
	"log"
)

type ReviewRepoImpl struct{
	db *sql.DB
}

func InitReviewRepoImpl(mydb *sql.DB) ReviewRepo {
	return &ReviewRepoImpl{db: mydb}
}

func (r *ReviewRepoImpl) CreateReview(review *models.Review) (*models.Review, error) {
	//panic("implement me")
	query := constant.CREATE_RATING_COMMENT
	review.ID = uuid.New().String()
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(query, review.ID, review.UserID, review.AssetID,
		review.Rating, review.Comment)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return review, tx.Commit()
}