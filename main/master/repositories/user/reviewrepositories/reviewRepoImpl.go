package reviewrepositories

import (
	"database/sql"
	"finalproject/main/master/models"
	constanta "finalproject/utils/constant"
	"fmt"
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

func (r *ReviewRepoImpl) GetStatusReview(user_id, asset_id string) error {
	query := "SELECT id, user_id, asset_id, rating, comment, created_at FROM m_review WHERE user_id = ? AND asset_id = ?;"
	reviewData := new(models.Review)
	err := r.db.QueryRow(query, user_id, asset_id).Scan(&reviewData.ID, &reviewData.UserID, &reviewData.AssetID, &reviewData.Rating,
		&reviewData.Comment, &reviewData.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
