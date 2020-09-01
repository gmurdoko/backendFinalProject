package feerepository

import (
	"database/sql"
	"finalproject/main/master/model"
)

//FeeRepositoryImpl is for init Repository
type FeeRepositoryImpl struct {
	db *sql.DB
}

//SelectFeeByID app
func (s FeeRepositoryImpl) SelectFeeByID(id string) (*model.Fees, error) {
	var fee = new(model.Fees)
	err := s.db.QueryRow("SELECT * FROM m_fee WHERE id = ?", id).Scan(&fee.ID, &fee.VehicleID, &fee.Fee, &fee.CreatedAt, &fee.EditedAt, &fee.DeletedAt, &fee.Status)
	if err != nil {
		return nil, err
	}
	return fee, nil
}

//InitFeeRepositoryImpl is init gate for repository
func InitFeeRepositoryImpl(db *sql.DB) FeeRepository {
	return &FeeRepositoryImpl{db}
}
