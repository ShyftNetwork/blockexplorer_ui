package accounts

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/types"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/db"
)

// AccountArrayMarshalling marshalls into account struct
func AccountArrayMarshalling(rows *sqlx.Rows) []byte {
	var a types.AccountPayload
	var accounts []byte

	for rows.Next() {
		account := types.Account{}
		err := rows.StructScan(&account)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
		}
		a.Payload = append(a.Payload, account)
		serializedPayload, err := json.Marshal(a.Payload)
		if err != nil {
			logger.Warn("Unable to marshal payload: " + err.Error())
		}
		accounts = serializedPayload
	}
	if err := rows.Err(); err != nil {
		logger.Warn("Unable to retrieve row: " + err.Error())
	}
	if err := rows.Close(); err != nil {
		logger.Warn("Unable to close row connection: " + err.Error())
	}
	return accounts
}

// AccountArrayQueries queries db
func AccountArrayQueries(db *db.SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) ([]byte, error) {
	var offset = (currentPage - 1) * pageLimit
	rows, err := db.Db.Queryx(query, pageLimit, offset)
	if err != nil {
		logger.Warn("Unable to connect: " + err.Error())
		return nil, err
	}
	accounts := AccountArrayMarshalling(rows)
	return accounts, nil
}

// AccountMarshalling marshalls bytes to struct
func AccountMarshalling(row *sqlx.Row) ([]byte, error) {
	account := types.Account{}
	err := row.StructScan(&account)
	if err != nil {
		logger.Warn("Unable to retrieve row: " + err.Error())
		return nil, err
	}
	serializedPayload, err := json.Marshal(account)
	if err != nil {
		logger.Warn("Unable to marshal payload: " + err.Error())
	}
	return serializedPayload, nil
}

// AccountQuery queries db
func AccountQuery(db *db.SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) ([]byte, error) {
	row := db.Db.QueryRowx(query, identifier)
	account, err := AccountMarshalling(row)
	if err != nil {
		logger.Warn("Unable to connect: " + err.Error())
		return nil, err
	}
	return account, nil
}