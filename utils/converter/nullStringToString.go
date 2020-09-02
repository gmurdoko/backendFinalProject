package converter

import (
	"database/sql"
)

//NullStringToString app
func NullStringToString(input sql.NullString) string {
	if input.Valid {
		return ""
	}
	return input.String
}
