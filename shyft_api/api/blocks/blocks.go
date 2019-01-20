package blocks

import (
	"encoding/json"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
	"github.com/jmoiron/sqlx"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/types"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/db"
)

// BlockArrayMarshalling marshalls into account struct
func BlockArrayMarshalling(rows *sqlx.Rows) []byte {
	var b types.BlockPayload
	var blocks []byte

	for rows.Next() {
		block := types.Block{}
		err := rows.StructScan(&block)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil
		}
		b.Payload = append(b.Payload, block)
		serializedPayload, err := json.Marshal(b.Payload)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil
		}
		blocks = serializedPayload
	}
	if err := rows.Close(); err != nil {
		logger.Warn("Unable to close row connection: " + err.Error())
	}
	return blocks
}

// BlockMarshalling marshalls into account struct
func BlockMarshalling(row *sqlx.Row) ([]byte, error) {
	block := types.Block{}
	err := row.StructScan(&block)
	if err != nil {
		logger.Warn("Unable to retrieve rows: " + err.Error())
		return nil, err
	}
	serializedPayload, err := json.Marshal(block)
	if err != nil {
		logger.Warn("Unable to serialize row: " + err.Error())
		return nil, err
	}
	return serializedPayload, nil
}

// BlockArrayQueries queries db
func BlockArrayQueries(db *db.SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) ([]byte, error) {
	switch {
	case len(identifier) > 0 && currentPage > 0:
		var offset = (currentPage - 1) * pageLimit
		rows, err := db.Db.Queryx(query, pageLimit, offset, identifier)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil, err
		}
		blocks := BlockArrayMarshalling(rows)
		return blocks, nil
	case currentPage > 0:
		var offset = (currentPage - 1) * pageLimit
		rows, err := db.Db.Queryx(query, pageLimit, offset)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil, err
		}
		blocks := BlockArrayMarshalling(rows)
		return blocks, nil
	default:
		rows, err := db.Db.Queryx(query)
		if err != nil {
			logger.Warn("Unable to retrieve rows: " + err.Error())
			return nil, err
		}
		blocks := BlockArrayMarshalling(rows)
		return blocks, nil
	}
}

// BlockQueries queries db
func BlockQueries(db *db.SPGDatabase, query string, identifier string) ([]byte, error) {
	switch {
	case len(identifier) > 0:
		row := db.Db.QueryRowx(query, identifier)
		b, err := BlockMarshalling(row)
		if err != nil {
			logger.Warn("Unable to query: " + err.Error())
			return nil, err
		}
		return b, nil
	default:
		row := db.Db.QueryRowx(query)
		b, err := BlockMarshalling(row)
		if err != nil {
			logger.Warn("Unable to query: " + err.Error())
			return nil, err
		}
		return b, nil
	}
}

