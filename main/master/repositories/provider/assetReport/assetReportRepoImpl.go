package providerassetreportrepo

import (
	"database/sql"
)

type AssetReportRepoImpl struct {
	db *sql.DB
}
