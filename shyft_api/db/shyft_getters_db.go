package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/types"
)

//SGetAllBlocksWithoutLimit returns all blocks in database; data dump
func SGetAllBlocksWithoutLimit() (string, error) {
	db := ConnectShyftDatabase()
	// arg1 := "10"
	// arg2 := "20"
	// arg3 := "30"
	// query := "SELECT THIS POOPSICLE"
	// purpose := "block"

	//ExplorerQuery(db, arg1, arg2, arg3, query, purpose)
	blocks, err := BlockArrayQueries(db, GetAllBlocksNoLimit, 0, 0)
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

	blocks, err := BlockArrayQueries(db, GetAllBlocks, currentPage, pageLimit)
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	stringifyBlocks := string(blocks)

	return stringifyBlocks, nil
}

//GetBlock queries to send single block info
//TODO provide blockHash arg passed from handler.go
func SGetBlock(blockNumber string) (string, error) {
	db := ConnectShyftDatabase()
	block := ExplorerQueryArray(db, GetBlock, blockNumber, "block")

	return string(block), nil
}

// func SGetRecentBlock() (string, error) {
// 	db := ConnectShyftDatabase()

// 	block, _ := BlockQueries(db, GetBlock, "")

// 	return string(block), nil
// }

// GetAllTransactionsLength returns number of records in blocks table
// This tells the client-UI how many pages to show for pagination
func GetAllTransactionsLength() (int, error) {
	db := ConnectShyftDatabase()

	count := RecordCountQuery(db, GetTransactionCount)

	return count, nil
}

func SGetAllTransactionsFromBlock(currentPage int64, pageLimit int64, blockNumber string) (string, error) {
	db := ConnectShyftDatabase()

	var arr types.TxRes
	var txx string
	sqlStatement := `SELECT * FROM txs WHERE blocknumber=$1`
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Query(sqlStatement, blockNumber)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	for rows.Next() {
		var txhash, to_addr, from_addr, txfee, blockhash, blocknumber, amount, status string
		var gasprice, gas, gasLimit, nonce uint64
		var isContract bool
		var age time.Time
		var data []byte

		err = rows.Scan(
			&txhash, &to_addr, &from_addr, &blockhash, &blocknumber, &amount, &gasprice, &gas, &gasLimit, &txfee, &nonce, &status, &isContract, &age, &data,
		)

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
		newtx := string(txData)
		txx = newtx
	}
	return txx, nil
}

func SGetAllBlocksMinedByAddress(coinbase string) (string, error) {
	db := ConnectShyftDatabase()

	var arr types.BlockRes
	var blockArr string
	sqlStatement := `SELECT * FROM blocks WHERE coinbase=$1`
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Query(sqlStatement, coinbase)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()

	for rows.Next() {
		var hash, coinbase, parentHash, uncleHash, difficulty, size, rewards, num string
		var gasUsed, gasLimit, nonce uint64
		var txCount, uncleCount int
		var age time.Time

		err = rows.Scan(
			&hash, &coinbase, &gasUsed, &gasLimit, &txCount, &uncleCount, &age, &parentHash, &uncleHash, &difficulty, &size, &nonce, &rewards, &num)

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

		blocks, _ := json.Marshal(arr.Blocks)
		blocksFmt := string(blocks)
		blockArr = blocksFmt
	}
	return blockArr, nil
}

//GetAllTransactions getter fn for API
func SGetAllTransactionsWithoutLimit() (string, error) {
	db := ConnectShyftDatabase()

	var arr types.TxRes
	var txx string
	sqlStatement := `SELECT * FROM txs`
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Queryx(sqlStatement)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	for rows.Next() {
		var txhash, to_addr, from_addr, txfee, blockhash, blocknumber, amount, status string
		var gasprice, gas, gasLimit, nonce uint64
		var isContract bool
		var age time.Time
		var data []byte

		err = rows.Scan(
			&txhash, &to_addr, &from_addr, &blockhash, &blocknumber, &amount, &gasprice, &gas, &gasLimit, &txfee, &nonce, &status, &isContract, &age, &data,
		)

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
		newtx := string(txData)
		txx = newtx
	}
	return txx, nil
}

//GetAllTransactions getter fn for API
func SGetAllTransactions(currentPage int64, pageLimit int64) (string, error) {
	db := ConnectShyftDatabase()

	var arr types.TxRes
	var txx string
	var offset = (currentPage - 1) * pageLimit
	sqlStatement := `SELECT * FROM txs ORDER BY age ASC LIMIT $1 OFFSET $2`
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Queryx(sqlStatement, pageLimit, offset)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	for rows.Next() {
		var txhash, to_addr, from_addr, txfee, blockhash, blocknumber, amount, status string
		var gasprice, gas, gasLimit, nonce uint64
		var isContract bool
		var age time.Time
		var data []byte

		err = rows.Scan(
			&txhash, &to_addr, &from_addr, &blockhash, &blocknumber, &amount, &gasprice, &gas, &gasLimit, &txfee, &nonce, &status, &isContract, &age, &data,
		)

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
		newtx := string(txData)
		txx = newtx
	}
	return txx, nil
}

//GetTransaction fn returns single tx
func SGetTransaction(txHash string) (string, error) {
	db := ConnectShyftDatabase()

	sqlStatement := `SELECT * FROM txs WHERE txhash=$1;`
	tx, _ := db.Db.Begin()
	row := db.Db.QueryRow(sqlStatement, txHash)
	tx.Commit()
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

	return string(json), nil
}

//GetAllAccounts returns all accounts and balances
func SGetAllAccounts() (string, error) {
	db := ConnectShyftDatabase()

	var array types.AccountRes
	var accountsArr, nonce string
	tx, _ := db.Db.Begin()
	accs, err := db.Db.Query(`
		SELECT
			addr,
			balance,
			nonce
		FROM accounts
		ORDER BY balance ASC`)
	tx.Commit()
	if err != nil {
		fmt.Println(err)
	}

	defer accs.Close()

	for accs.Next() {
		var addr, balance string
		err = accs.Scan(
			&addr, &balance, &nonce,
		)

		array.AllAccounts = append(array.AllAccounts, types.SAccounts{
			Addr:         addr,
			Balance:      balance,
			AccountNonce: nonce,
		})

		accounts, _ := json.Marshal(array.AllAccounts)
		accountsFmt := string(accounts)
		accountsArr = accountsFmt
	}
	return accountsArr, nil
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

	var account, _ = InnerSGetAccount(db, address)
	json, _ := json.Marshal(account)
	return string(json), nil
}

//GetAccount returns account balances
func SGetAccountTxs(address string) (string, error) {
	db := ConnectShyftDatabase()

	var arr types.TxRes
	var txx string
	sqlStatement := `SELECT * FROM txs WHERE to_addr=$1 OR from_addr=$1;`
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Query(sqlStatement, address)
	tx.Commit()
	if err != nil {
		fmt.Println("err", err)
	}
	defer rows.Close()
	for rows.Next() {
		var txhash, to_addr, from_addr, txfee, blockhash, blocknumber, amount, status string
		var gasprice, gas, gasLimit, nonce uint64
		var isContract bool
		var age time.Time
		var data []byte

		err = rows.Scan(
			&txhash, &to_addr, &from_addr, &blockhash, &blocknumber, &amount, &gasprice, &gas, &gasLimit, &txfee, &nonce, &status, &isContract, &age, &data,
		)

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
		newtx := string(txData)
		txx = newtx
	}
	return txx, nil
}

//GetAllInternalTransactions getter fn for API
func SGetAllInternalTransactions() (string, error) {
	db := ConnectShyftDatabase()

	var arr types.InternalArray
	var internaltx string
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Query(`SELECT * FROM internaltxs`)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	for rows.Next() {
		var txhash, blockhash, action, to_addr, from_addr, amount, input, output string
		var gas, gasUsed uint64
		var id int
		var age string

		err = rows.Scan(
			&id, &txhash, &blockhash, &action, &to_addr, &from_addr, &amount, &gas, &gasUsed, &age, &input, &output,
		)

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
		newtx := string(txData)
		internaltx = newtx
	}
	return internaltx, nil
}

//GetInternalTransaction fn returns single tx
func SGetInternalTransaction(txHash string) (string, error) {
	db := ConnectShyftDatabase()

	var arr types.InternalArray
	var internaltx string

	sqlStatement := `SELECT * FROM internaltxs WHERE txhash=$1;`
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Query(sqlStatement, txHash)
	tx.Commit()
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()

	for rows.Next() {
		var txhash, blockhash, action, to_addr, from_addr, amount, input, output string
		var id int
		var gas, gasUsed uint64
		var age string

		err = rows.Scan(
			&id, &txhash, &blockhash, &action, &to_addr, &from_addr, &amount, &gas, &gasUsed, &age, &input, &output,
		)

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
		newtx := string(txData)
		internaltx = newtx
	}
	return internaltx, nil
}

func SGetAllAccountBlocks() (string, error) {
	db := ConnectShyftDatabase()

	var arr types.AccountBlockArray
	var accountBlockJSON string
	tx, _ := db.Db.Begin()
	rows, err := db.Db.Query(`SELECT * FROM accountblocks`)
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
