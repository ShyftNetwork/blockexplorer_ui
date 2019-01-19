package accounts

import (
	"github.com/jmoiron/sqlx"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/types"
	"encoding/json"
	"fmt"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/db"
)

func AccountArrayMarshalling(rows *sqlx.Rows) []byte {
	var a types.AccountPayload
	var accounts []byte

	defer rows.Close()

	for rows.Next() {
		account := types.Account{}
		err := rows.StructScan(&account)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
		}
		a.Payload = append(a.Payload, account)
		serializedPayload, err := json.Marshal(a.Payload)
		accounts = serializedPayload
	}
	return accounts
}

func AccountArrayQueries(db *db.SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) []byte {
	var offset = (currentPage - 1) * pageLimit
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Queryx(query, pageLimit, offset)
	tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
	accounts := AccountArrayMarshalling(rows)
	return accounts
}

func AccountMarshalling(row *sqlx.Row) ([]byte, error) {
	account := types.Account{}
	err := row.StructScan(&account)
	if err != nil {
		logger.Warn("Unable to retrieve row: " + err.Error())
		return nil, err
	}
	serializedPayload, err := json.Marshal(account)
	return serializedPayload, nil
}

func AccountQuery(db *db.SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) ([]byte, error) {
	tx, _ := db.Db.Begin()
	row := db.Db.QueryRowx(query, identifier)
	tx.Commit()
	account, err := AccountMarshalling(row)
	if err != nil {
		return nil, err
	}
	return account, nil
}