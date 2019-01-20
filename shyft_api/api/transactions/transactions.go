package transactions

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/types"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/db"
)

// TransactionArrayMarshalling marshalls into tx struct
func TransactionArrayMarshalling(rows *sqlx.Rows) []byte {
	var t types.TransactionPayload
	var txs []byte

	for rows.Next() {
		tx := types.Transaction{}
		err := rows.StructScan(&tx)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil
		}
		t.Payload = append(t.Payload, tx)
		serializedPayload, err := json.Marshal(t.Payload)
		if err != nil {
			logger.Warn("Unable to serialize payload: " + err.Error())
		}
		txs = serializedPayload
	}
	if err := rows.Close(); err != nil {
		logger.Warn("Unable to close row connection: " + err.Error())
	}
	return txs
}

// TransactionArrayQueries queries db
func TransactionArrayQueries(db *db.SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) ([]byte, error) {
	switch {
	case len(identifier) > 0 && currentPage > 0:
		var offset = (currentPage - 1) * pageLimit
		rows, err := db.Db.Queryx(query, pageLimit, offset, identifier)
		if err != nil {
			logger.Warn("Unable to query: " + err.Error())
			return nil, err
		}
		txs := TransactionArrayMarshalling(rows)
		return txs, nil
	case currentPage > 0:
		var offset = (currentPage - 1) * pageLimit
		rows, err := db.Db.Queryx(query, pageLimit, offset)
		if err != nil {
			logger.Warn("Unable to query: " + err.Error())
			return nil, err
		}
		txs := TransactionArrayMarshalling(rows)
		return txs, nil
	default:
		rows, err := db.Db.Queryx(query)
		if err != nil {
			logger.Warn("Unable to query: " + err.Error())
			return nil, err
		}
		txs := TransactionArrayMarshalling(rows)
		return txs, nil
	}
}

// TransactionMarshalling marshalls into tx struct
func TransactionMarshalling(row *sqlx.Row) []byte {
	tx := types.Transaction{}
	err := row.StructScan(&tx)
	if err != nil {
		logger.Warn("Unable to retrieve row: " + err.Error())
		return nil
	}
	serializedPayload, err := json.Marshal(tx)
	if err != nil {
		logger.Warn("Unable to serialize payload: " + err.Error())
	}
	return serializedPayload
}

// TransactionQuery queries db
func TransactionQuery(db *db.SPGDatabase, query string, identifier string) ([]byte, error) {
	row := db.Db.QueryRowx(query, identifier)
	transaction := TransactionMarshalling(row)
	return transaction, nil
}

// InternalTransactionArrayMarshalling marshalls into internaltx struct
func InternalTransactionArrayMarshalling(rows *sqlx.Rows) []byte {
	var i types.InternalTransactionPayload
	var internals []byte

	for rows.Next() {
		internal := types.InteralTransaction{}
		err := rows.StructScan(&internal)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
		}
		i.Payload = append(i.Payload, internal)
		serializedPayload, err := json.Marshal(i.Payload)
		if err != nil {
			logger.Warn("Unable to serialize payload: " + err.Error())
		}
		internals = serializedPayload
	}
	if err := rows.Close(); err != nil {
		logger.Warn("Unable to close row connection: " + err.Error())
	}
	return internals
}

// InternalTransactionArrayQuery queries db
func InternalTransactionArrayQuery(db *db.SPGDatabase, query string, currentPage int64, pageLimit int64, identifer string) ([]byte, error) {
	var offset = (currentPage - 1) * pageLimit
	rows, err := db.Db.Queryx(query, pageLimit, offset, identifer)
	if err != nil {
		logger.Warn("Unable to query: " + err.Error())
		return nil, err
	}
	txs := InternalTransactionArrayMarshalling(rows)
	return txs, nil
}

