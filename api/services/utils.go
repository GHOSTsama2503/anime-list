package services

import "database/sql"

func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func NullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: i != 0}
}
