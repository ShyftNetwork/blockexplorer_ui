package db

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/types"
	"github.com/jmoiron/sqlx"
)

// ConnectShyftDatabase connects to open postgres db
func ConnectShyftDatabase() *SPGDatabase {
	db, err := ReturnShyftDatabase()
	if err != nil {
		fmt.Println("err", err)
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

func BlockArrayMarshalling(rows *sqlx.Rows) ([]byte, error) {
	var arr types.BlockRes
	var blocks []byte

	defer rows.Close()

	for rows.Next() {
		var hash, coinbase, parentHash, uncleHash, difficulty, size, rewards, num string
		var gasUsed, gasLimit, nonce uint64
		var txCount, uncleCount int
		var age time.Time

		err := rows.Scan(
			&hash, &coinbase, &gasUsed, &gasLimit, &txCount, &uncleCount, &age, &parentHash, &uncleHash, &difficulty, &size, &nonce, &rewards, &num)

		if err != nil {
			fmt.Println("Error:: ", err)
			return nil, err
		}

		arr.Blocks = append(arr.Blocks, types.SBlock{
			Hash:       hash,
			Coinbase:   coinbase,
			GasUsed:    gasUsed,
			GasLimit:   gasLimit,
			TxCount:    txCount,
			UncleCount: uncleCount,
			Age:        age,
			ParentHash: parentHash,
			UncleHash:  uncleHash,
			Difficulty: difficulty,
			Size:       size,
			Nonce:      nonce,
			Rewards:    rewards,
			Number:     num,
		})
		b, _ := json.Marshal(arr.Blocks)
		blocks = b
	}
	return blocks, nil
}

func BlockMarshalling(row *sqlx.Row) ([]byte, error) {
	var hash, coinbase, parentHash, uncleHash, difficulty, size, rewards, num string
	var gasUsed, gasLimit, nonce uint64
	var txCount, uncleCount int
	var age time.Time

	row.Scan(
		&hash, &coinbase, &gasUsed, &gasLimit, &txCount, &uncleCount, &age, &parentHash, &uncleHash, &difficulty, &size, &nonce, &rewards, &num)

	block := types.SBlock{
		Hash:       hash,
		Coinbase:   coinbase,
		GasUsed:    gasUsed,
		GasLimit:   gasLimit,
		TxCount:    txCount,
		UncleCount: uncleCount,
		Age:        age,
		ParentHash: parentHash,
		UncleHash:  uncleHash,
		Difficulty: difficulty,
		Size:       size,
		Nonce:      nonce,
		Rewards:    rewards,
		Number:     num,
	}
	b, _ := json.Marshal(block)
	return b, nil
}

func BlockArrayQueries(db *SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) ([]byte, error) {
	switch {
	case len(identifier) > 0 && currentPage > 0:
		var offset = (currentPage - 1) * pageLimit
		rows, err := db.Db.Queryx(query, pageLimit, offset, identifier)
		if err != nil {
			fmt.Println("Error:: ", err)
			return nil, err
		}
		blocks, _ := BlockArrayMarshalling(rows)
		return blocks, nil
	case currentPage > 0:
		var offset = (currentPage - 1) * pageLimit
		rows, err := db.Db.Queryx(query, pageLimit, offset)
		if err != nil {
			fmt.Println("Error:: ", err)
			return nil, err
		}
		blocks, _ := BlockArrayMarshalling(rows)
		return blocks, nil
	default:
		rows, err := db.Db.Queryx(query)
		if err != nil {
			fmt.Println("Error:: ", err)
			return nil, err
		}
		blocks, _ := BlockArrayMarshalling(rows)
		return blocks, nil
	}
}

func BlockQueries(db *SPGDatabase, query string, identifier string) ([]byte, error) {
	tx, _ := db.Db.Begin()
	switch {
	case len(identifier) > 0:
		row := db.Db.QueryRowx(query, identifier)
		b, _ := BlockMarshalling(row)
		tx.Commit()
		return b, nil
	default:
		row := db.Db.QueryRowx(query)
		b, _ := BlockMarshalling(row)
		tx.Commit()
		return b, nil
	}
}

func TransactionArrayMarshalling(rows *sqlx.Rows) ([]byte, error) {
	var arr types.TxRes
	var txs []byte

	defer rows.Close()
	for rows.Next() {
		var txhash, to_addr, from_addr, txfee, blockhash, blocknumber, amount, status string
		var gasprice, gas, gasLimit, nonce uint64
		var isContract bool
		var age time.Time
		var data []byte

		err := rows.Scan(
			&txhash, &to_addr, &from_addr, &blockhash, &blocknumber, &amount, &gasprice, &gas, &gasLimit, &txfee, &nonce, &status, &isContract, &age, &data,
		)
		if err != nil {
			fmt.Println("Error:: ", err)
			return nil, err
		}

		arr.TxEntry = append(arr.TxEntry, types.ShyftTxEntryPretty{
			TxHash:      txhash,
			To:          to_addr,
			From:        from_addr,
			BlockHash:   blockhash,
			BlockNumber: blocknumber,
			Amount:      amount,
			GasPrice:    gasprice,
			Gas:         gas,
			GasLimit:    gasLimit,
			Cost:        txfee,
			Nonce:       nonce,
			Status:      status,
			IsContract:  isContract,
			Age:         age,
			Data:        data,
		})
		txData, _ := json.Marshal(arr.TxEntry)
		txs = txData
	}
	return txs, nil
}

func TransactionArrayQueries(db *SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) ([]byte, error) {
	switch {
	case len(identifier) > 0 && currentPage > 0:
		var offset = (currentPage - 1) * pageLimit
		rows, err := db.Db.Queryx(query, pageLimit, offset, identifier)
		if err != nil {
			fmt.Println("Error:: ", err)
			return nil, err
		}
		txs, _ := TransactionArrayMarshalling(rows)
		return txs, nil
	case currentPage > 0:
		var offset = (currentPage - 1) * pageLimit
		rows, err := db.Db.Queryx(query, pageLimit, offset)
		if err != nil {
			fmt.Println("Error:: ", err)
			return nil, err
		}
		txs, _ := TransactionArrayMarshalling(rows)
		return txs, nil
	default:
		rows, err := db.Db.Queryx(query)
		if err != nil {
			fmt.Println("Error:: ", err)
			return nil, err
		}
		txs, _ := TransactionArrayMarshalling(rows)
		return txs, nil
	}
}

func TransactionMarshalling(row *sqlx.Row) []byte {
	var txhash, to_addr, from_addr, txfee, blockhash, blocknumber, amount, status string
	var gasprice, gas, gasLimit, nonce uint64
	var isContract bool
	var age time.Time
	var data []byte

	row.Scan(
		&txhash, &to_addr, &from_addr, &blockhash, &blocknumber, &amount, &gasprice, &gas, &gasLimit, &txfee, &nonce, &status, &isContract, &age, &data)

	txData := types.ShyftTxEntryPretty{
		TxHash:      txhash,
		To:          to_addr,
		From:        from_addr,
		BlockHash:   blockhash,
		BlockNumber: blocknumber,
		Amount:      amount,
		GasPrice:    gasprice,
		Gas:         gas,
		GasLimit:    gasLimit,
		Cost:        txfee,
		Nonce:       nonce,
		Status:      status,
		IsContract:  isContract,
		Age:         age,
		Data:        data,
	}
	json, _ := json.Marshal(txData)
	return json
}

func TransactionQuery(db *SPGDatabase, query string, identifier string) []byte {
	tx, _ := db.Db.Begin()
	row := db.Db.QueryRowx(query, identifier)
	tx.Commit()
	transaction := TransactionMarshalling(row)

	return transaction
}

func AccountArrayMarshalling(rows *sqlx.Rows) []byte {
	var accounts []byte
	var array types.AccountRes

	defer rows.Close()

	for rows.Next() {
		var addr, balance, nonce string
		err := rows.Scan(
			&addr, &balance, &nonce,
		)
		if err != nil {
			fmt.Println("ERROR::", err)
		}
		array.AllAccounts = append(array.AllAccounts, types.SAccounts{
			Addr:         addr,
			Balance:      balance,
			AccountNonce: nonce,
		})
		acc, _ := json.Marshal(array.AllAccounts)
		accounts = acc
	}
	return accounts
}

func AccountArrayQueries(db *SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) []byte {
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

func AccountMarshalling(row *sqlx.Row) []byte {
	var addr, balance, nonce string
	row.Scan(&addr, &balance, &nonce)
	account := types.SAccounts{
		Addr:         addr,
		Balance:      balance,
		AccountNonce: nonce,
	}
	acc, _ := json.Marshal(account)
	return acc
}

func AccountQuery(db *SPGDatabase, query string, currentPage int64, pageLimit int64, identifier string) []byte {
	tx, _ := db.Db.Begin()
	row := db.Db.QueryRowx(query, identifier)
	tx.Commit()
	account := AccountMarshalling(row)
	return account
}

func InternalTransactionArrayMarshalling(rows *sqlx.Rows) []byte {
	var arr types.InternalArray
	var internalTx []byte

	defer rows.Close()
	for rows.Next() {
		var txhash, blockhash, action, to_addr, from_addr, amount, input, output string
		var gas, gasUsed uint64
		var id int
		var age string

		err := rows.Scan(
			&id, &txhash, &blockhash, &action, &to_addr, &from_addr, &amount, &gas, &gasUsed, &age, &input, &output,
		)
		if err != nil {
			fmt.Println("err")
		}

		arr.InternalEntry = append(arr.InternalEntry, types.InteralWrite{
			ID:        id,
			Hash:      txhash,
			BlockHash: blockhash,
			Action:    action,
			To:        to_addr,
			From:      from_addr,
			Value:     amount,
			Gas:       gas,
			GasUsed:   gasUsed,
			Time:      age,
			Input:     input,
			Output:    output,
		})

		txData, _ := json.Marshal(arr.InternalEntry)
		internalTx = txData
	}
	return internalTx
}

func InternalTransactionArrayQuery(db *SPGDatabase, query string, currentPage int64, pageLimit int64) []byte {
	var offset = (currentPage - 1) * pageLimit
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Queryx(query, pageLimit, offset)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	txs := InternalTransactionArrayMarshalling(rows)
	return txs
}

func GetPageAndLimit(pagination []string) (int64, int64) {
	page, err := strconv.ParseInt(strings.Join(pagination[:1], " "), 10, 64)
	limit, err := strconv.ParseInt(strings.Join(pagination[1:2], " "), 10, 64)
	if err != nil {
		println("ERROR", err)
	}
	return page, limit
}

func SliceStringToString(query []string) string {
	q := strings.Join(query, " ")
	return q
}
