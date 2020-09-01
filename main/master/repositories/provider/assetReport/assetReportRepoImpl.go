package providerassetreportrepo

import (
	"database/sql"
	"finalproject/main/master/models"
	constanta "finalproject/utils/constant"
	"log"
)

type ProviderAssetReportRepoImpl struct {
	db *sql.DB
}

func InitProviderAssetReportRepoImpl(mydb *sql.DB) ProviderAssetReportRepo {
	return &ProviderAssetReportRepoImpl{db: mydb}
}

func (s *ProviderAssetReportRepoImpl) GetReportDaily(id string) ([]*models.ReportAssetDaily, error) {
	query := constanta.PROVIDERASSETSDAILYREPORT
	rows, err := s.db.Query(query, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var reports []*models.ReportAssetDaily
	for rows.Next() {
		report := models.ReportAssetDaily{}
		err := rows.Scan(&report.AssetName, &report.Date, &report.TotalParked, &report.TotalRevenue)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		reports = append(reports, &report)
	}
	return reports, nil
}

func (s *ProviderAssetReportRepoImpl) GetReportMonthly(id string) ([]*models.ReportAssetMonthly, error) {
	query := constanta.PROVIDERASSETSMONTHLYREPORT
	rows, err := s.db.Query(query, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var reports []*models.ReportAssetMonthly
	for rows.Next() {
		report := models.ReportAssetMonthly{}
		err := rows.Scan(&report.AssetName, &report.Months, &report.TotalParked, &report.TotalRevenue)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		reports = append(reports, &report)
	}
	return reports, nil
}
