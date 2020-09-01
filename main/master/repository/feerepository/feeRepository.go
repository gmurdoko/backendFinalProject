package feerepository

import (
	"finalproject/main/master/model"
)

// FeeRepository app
type FeeRepository interface {
	SelectFeeByID(id string) *model.Fees
}
