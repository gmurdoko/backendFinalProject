package providerassetreportrepo

import (
	"database/sql"
)

type ProviderAssetReportRepoImpl struct {
	db *sql.DB
}
