package db

import (
	"strconv"
	"strings"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
)

// ConnectShyftDatabase connects to open postgres db
func ConnectShyftDatabase() *SPGDatabase {
	db, err := ReturnShyftDatabase()
	if err != nil {
		logger.Warn("Unable to connect to database: " + err.Error())
		return nil
	}
	return db
}

func RecordCountQuery(db *SPGDatabase, query string) int {
	tx, _ := db.Db.Begin()
	row := db.Db.QueryRow(query)
	tx.Commit()

	var count int
	row.Scan(&count)

	return count
}

func GetPageAndLimit(pagination []string) (int64, int64) {
	page, err := strconv.ParseInt(strings.Join(pagination[:1], " "), 10, 64)
	limit, err := strconv.ParseInt(strings.Join(pagination[1:2], " "), 10, 64)
	if err != nil {
		logger.Warn("Unable to query: " + err.Error())
	}
	return page, limit
}

func SliceStringToString(query []string) string {
	q := strings.Join(query, " ")
	return q
}
