package reviewrepositories

import (
	"database/sql"
	"finalproject/main/master/models"
	constanta "finalproject/utils/constant"
	"log"

	"github.com/google/uuid"
)

type ReviewRepoImpl struct {
	db *sql.DB
}

func InitReviewRepoImpl(mydb *sql.DB) ReviewRepo {
	return &ReviewRepoImpl{db: mydb}
}

func (r *ReviewRepoImpl) CreateReview(review *models.Review) (*models.Review, error) {
	query := constanta.CREATE_RATING_COMMENT
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
