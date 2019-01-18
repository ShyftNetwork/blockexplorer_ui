package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/types"
)

//SGetAllBlocksWithoutLimit returns all blocks in database; data dump
func SGetAllBlocksWithoutLimit() (string, error) {
	db := ConnectShyftDatabase()

	blocks, err := BlockArrayQueries(db, GetAllBlocksNoLimit, 0, 0, "")
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	stringifyBlocks := string(blocks)

	return stringifyBlocks, nil
}

// GetAllBlocksLength returns number of records in blocks table
// This tells the client-UI how many pages to show for pagination
func GetAllBlocksLength() (int, error) {
	db := ConnectShyftDatabase()
	count := RecordCountQuery(db, GetBlockCount)

	return count, nil
}

func SGetAllBlocks(currentPage int64, pageLimit int64) (string, error) {
	db := ConnectShyftDatabase()

	blocks, err := BlockArrayQueries(db, GetAllBlocks, currentPage, pageLimit, "")
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	stringifyBlocks := string(blocks)

	return stringifyBlocks, nil
}

//SGetBlock queries to send single block info
//TODO provide blockHash arg passed from handler.go
func SGetBlock(blockNumber string) (string, error) {
	db := ConnectShyftDatabase()
	block, _ := BlockQueries(db, GetBlock, blockNumber)

	return string(block), nil
}

//SGetRecentBlock Returns recent block height
func SGetRecentBlock() (string, error) {
	db := ConnectShyftDatabase()

	block, _ := BlockQueries(db, GetBlock, "")

	return string(block), nil
}

// GetAllTransactionsLength returns number of records in blocks table
// This tells the client-UI how many pages to show for pagination
func GetAllTransactionsLength() (int, error) {
	db := ConnectShyftDatabase()

	count := RecordCountQuery(db, GetTransactionCount)

	return count, nil
}

func SGetAllTransactionsFromBlock(currentPage int64, pageLimit int64, blockNumber string) string {
	db := ConnectShyftDatabase()

	transactions, _ := TransactionArrayQueries(db, GetAllTransactionsFromBlock, currentPage, pageLimit, blockNumber)
	return string(transactions)
}

func SGetAllBlocksMinedByAddress(currentPage int64, pageLimit int64, coinbase string) string {
	db := ConnectShyftDatabase()

	blocks, _ := BlockArrayQueries(db, GetAllBlocksMinedByAddress, currentPage, pageLimit, coinbase)
	return string(blocks)
}

//GetAllTransactions getter fn for API
func SGetAllTransactionsWithoutLimit() string {
	db := ConnectShyftDatabase()

	transactions, _ := TransactionArrayQueries(db, GetAllTransactionsNoLimit, 0, 0, "")
	return string(transactions)
}

//GetAllTransactions getter fn for API
func SGetAllTransactions(currentPage int64, pageLimit int64) string {
	db := ConnectShyftDatabase()

	transactions, _ := TransactionArrayQueries(db, GetAllTransactions, currentPage, pageLimit, "")
	return string(transactions)
}

//GetTransaction fn returns single tx
func SGetTransaction(txHash string) string {
	db := ConnectShyftDatabase()

	transaction := TransactionQuery(db, GetTransaction, txHash)
	return string(transaction)
}

// GetAllTransactionsLength returns number of records in blocks table
// This tells the client-UI how many pages to show for pagination
func GetAllAccountsLength() (int, error) {
	db := ConnectShyftDatabase()

	count := RecordCountQuery(db, GetAccountCount)

	return count, nil
}

//GetAllAccounts returns all accounts and balances
func SGetAllAccounts(currentPage int64, pageLimit int64) string {
	db := ConnectShyftDatabase()

	accounts := AccountArrayQueries(db, GetAllAccounts, currentPage, pageLimit, "")
	return string(accounts)
}

func InnerSGetAccount(db *SPGDatabase, address string) (types.SAccounts, bool) {
	sqlStatement := `SELECT * FROM accounts WHERE addr=$1;`
	var addr, balance, nonce string
	tx, _ := db.Db.Begin()
	err := db.Db.QueryRow(sqlStatement, address).Scan(&addr, &balance, &nonce)
	tx.Commit()
	if err == sql.ErrNoRows {
		return types.SAccounts{}, false
	} else {
		account := types.SAccounts{
			Addr:         addr,
			Balance:      balance,
			AccountNonce: nonce,
		}
		return account, true
	}
}

//GetAccount returns account balances
func SGetAccount(address string) (string, error) {
	db := ConnectShyftDatabase()

	account := AccountQuery(db, GetAccount, 0, 0, address)
	return string(account), nil
}

//GetAccount returns account balances
func SGetAccountTxs(currentPage int64, pageLimit int64, address string) (string, error) {
	db := ConnectShyftDatabase()

	transactions, _ := TransactionArrayQueries(db, GetAccountTransactions, currentPage, pageLimit, address)
	return string(transactions), nil
}

// GetAllTransactionsLength returns number of records in blocks table
// This tells the client-UI how many pages to show for pagination
func GetAllInternalTransactionsLength() (int, error) {
	db := ConnectShyftDatabase()

	count := RecordCountQuery(db, GetInternalTransactionLength)

	return count, nil
}

//GetAllInternalTransactions getter fn for API
func SGetAllInternalTransactions(currentPage int64, pageLimit int64) (string, error) {
	db := ConnectShyftDatabase()

	transactions := InternalTransactionArrayQuery(db, GetAllInternalTransactions, currentPage, pageLimit)
	return string(transactions), nil
}

//GetInternalTransaction fn returns single tx
func SGetInternalTransaction(currentPage int64, pageLimit int64, txHash string) (string, error) {
	db := ConnectShyftDatabase()

	transactions := InternalTransactionArrayQuery(db, GetAllInternalTransactions, currentPage, pageLimit)
	return string(transactions), nil
}

func SGetAllAccountBlocks() (string, error) {
	db := ConnectShyftDatabase()

	var arr types.AccountBlockArray
	var accountBlockJSON string
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Query(GetAllAccountBlocks)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	for rows.Next() {
		var acct, blockhash string
		var delta, txCount int64

		err = rows.Scan(
			&acct, &blockhash, &delta, &txCount,
		)

		arr.AccountBlocks = append(arr.AccountBlocks, types.AccountBlock{
			Acct:      acct,
			Blockhash: blockhash,
			Delta:     delta,
			TxCount:   txCount,
		})

		accountBlocks, _ := json.Marshal(arr.AccountBlocks)
		accountBlocksJSON := string(accountBlocks)
		accountBlockJSON = accountBlocksJSON
	}
	return accountBlockJSON, nil
}
