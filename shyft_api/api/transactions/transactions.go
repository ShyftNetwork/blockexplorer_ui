package transactions

import (
	"github.com/jmoiron/sqlx"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/types"
	"encoding/json"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/db"
)

func TransactionArrayMarshalling(rows *sqlx.Rows) ([]byte, error) {
	var t types.TransactionPayload
	var txs []byte

	defer rows.Close()
	for rows.Next() {
		tx := types.Transaction{}
		err := rows.StructScan(&tx)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil, err
		}
		t.Payload = append(t.Payload, tx)
		serializedPayload, err := json.Marshal(t.Payload)
		txs = serializedPayload
	}
	return txs, nil
}

func TransactionArrayQueries(db *db.SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) ([]byte, error) {
	switch {
	case len(identifier) > 0 && currentPage > 0:
		var offset = (currentPage - 1) * pageLimit
		rows, err := db.Db.Queryx(query, pageLimit, offset, identifier)
		if err != nil {
			logger.Warn("Unable to query: " + err.Error())
			return nil, err
		}
		txs, _ := TransactionArrayMarshalling(rows)
		return txs, nil
	case currentPage > 0:
		var offset = (currentPage - 1) * pageLimit
		rows, err := db.Db.Queryx(query, pageLimit, offset)
		if err != nil {
			logger.Warn("Unable to query: " + err.Error())
			return nil, err
		}
		txs, _ := TransactionArrayMarshalling(rows)
		return txs, nil
	default:
		rows, err := db.Db.Queryx(query)
		if err != nil {
			logger.Warn("Unable to query: " + err.Error())
			return nil, err
		}
		txs, _ := TransactionArrayMarshalling(rows)
		return txs, nil
	}
}

func TransactionMarshalling(row *sqlx.Row) ([]byte, error) {
	tx := types.Transaction{}
	err := row.StructScan(&tx)
	if err != nil {
		logger.Warn("Unable to retrieve row: " + err.Error())
		return nil, err
	}
	serializedPayload, err := json.Marshal(tx)
	return serializedPayload, nil
}

func TransactionQuery(db *db.SPGDatabase, query string, identifier string) ([]byte, error) {
	tx, _ := db.Db.Begin()
	row := db.Db.QueryRowx(query, identifier)
	tx.Commit()
	transaction, err := TransactionMarshalling(row)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func InternalTransactionArrayMarshalling(rows *sqlx.Rows) []byte {
	var i types.InternalTransactionPayload
	var internals []byte

	defer rows.Close()

	for rows.Next() {
		internal := types.InteralTransaction{}
		err := rows.StructScan(&internal)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
		}
		i.Payload = append(i.Payload, internal)
		serializedPayload, err := json.Marshal(i.Payload)
		internals = serializedPayload
	}
	return internals
}

func InternalTransactionArrayQuery(db *db.SPGDatabase, query string, currentPage int64, pageLimit int64, identifer string) ([]byte, error) {
	var offset = (currentPage - 1) * pageLimit
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Queryx(query, pageLimit, offset, identifer)
	tx.Commit()
	if err != nil {
		logger.Warn("Unable to query: " + err.Error())
		return nil, err
	}
	txs := InternalTransactionArrayMarshalling(rows)
	return txs, nil
}

