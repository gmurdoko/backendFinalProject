package adminassetsreportsrepo

import (
	"database/sql"
	"finalproject/main/master/models"
	constanta "finalproject/utils/constant"
	"fmt"
	"log"
)

type AdminAssetReportRepoImpl struct {
	db *sql.DB
}

func InitAdminAssetReportRepoImpl(mydb *sql.DB) AdminAssetReportRepo {
	return &AdminAssetReportRepoImpl{db: mydb}
}

func (s *AdminAssetReportRepoImpl) GetReportDaily(start, end, id string) ([]*models.ReportAssetDaily, error) {
	query := constanta.ADMINASSETSDAILYREPORT
	rows, err := s.db.Query(query, id, start, end)
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
		fmt.Println(report.AssetName)
		reports = append(reports, &report)
	}
	return reports, nil
}

func (s *AdminAssetReportRepoImpl) GetReportMonthly(id string) ([]*models.ReportAssetMonthly, error) {
	query := constanta.ADMINASSETSMONTHLYREPORT
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
