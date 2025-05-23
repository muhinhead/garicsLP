package utils

import (
	"database/sql"
	"strconv"
)

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func NullStrToStr(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func NullIntToStr(ni sql.NullInt64) string {
	if ni.Valid {
		return strconv.FormatInt(ni.Int64, 10)
	}
	return ""
}
