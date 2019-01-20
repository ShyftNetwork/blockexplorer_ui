package api

///@NOTE Shyft handler functions when endpoints are hit
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq" //github.com/lib/pq needed for sqlx transactions
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/db"
	"github.com/gorilla/mux"
	b "github.com/ShyftNetwork/blockexplorer_ui/shyft_api/api/blocks"
	tx "github.com/ShyftNetwork/blockexplorer_ui/shyft_api/api/transactions"
	acc "github.com/ShyftNetwork/blockexplorer_ui/shyft_api/api/accounts"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/logger"
	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/api/common"
)

// SGetAllTransactionsLength Count all rows in Blocks Table
func SGetAllTransactionsLength(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	count := b.RecordCountQuery(dbase, db.GetBlockCount)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(count))
}

// GetAllTransactionsWithoutLimit returns all rows in Blocks Table
func GetAllTransactionsWithoutLimit(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	txs, err := tx.TransactionArrayQueries(dbase, db.GetAllTransactionsNoLimit, 0, 0, "")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(txs))
}

// GetTransaction gets txs
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	txHash := vars["txHash"]
	transaction, err := tx.TransactionQuery(dbase, db.GetTransaction, txHash)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(transaction))
}

// GetAllTransactions gets txs
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]
	page := common.StringToInteger(currentPage)
	limit := common.StringToInteger(pageLimit)

	txs, err := tx.TransactionArrayQueries(dbase, db.GetAllTransactions, page, limit, "")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(txs))
}

// GetAllTransactionsFromBlock returns all txs from specified block
func GetAllTransactionsFromBlock(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]
	blockNumber := vars["blockNumber"]
	page := common.StringToInteger(currentPage)
	limit := common.StringToInteger(pageLimit)

	transactions, err := tx.TransactionArrayQueries(dbase, db.GetAllTransactionsFromBlock, page, limit, blockNumber)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(transactions))
}

// GetAllBlocksMinedByAddress returns all blocks mined by specific address
func GetAllBlocksMinedByAddress(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	coinbase := vars["coinbase"]
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]
	page := common.StringToInteger(currentPage)
	limit := common.StringToInteger(pageLimit)

	blocks, err := b.BlockArrayQueries(dbase, db.GetAllBlocksMinedByAddress, page, limit, coinbase)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(blocks))
}

// SGetAllAccountsLength Count all rows in accounts Table
func SGetAllAccountsLength(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	count := b.RecordCountQuery(dbase, db.GetAccountCount)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(count))
}

// GetAccount returns specific account data; balance, nonce
func GetAccount(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	address := vars["address"]

	account, err := acc.AccountQuery(dbase, db.GetAccount, 0, 0, address)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(account))
}

// GetAccountTxs returns account txs
func GetAccountTxs(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	address := vars["address"]
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]
	page := common.StringToInteger(currentPage)
	limit := common.StringToInteger(pageLimit)

	transactions, err := tx.TransactionArrayQueries(dbase, db.GetAccountTransactions, page, limit, address)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(transactions))
}

// GetAllAccounts returns all accounts
func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]
	page := common.StringToInteger(currentPage)
	limit := common.StringToInteger(pageLimit)

	accounts, err := acc.AccountArrayQueries(dbase, db.GetAllAccounts, page, limit, "")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(accounts))
}

//GetBlock returns contextual block data
func GetBlock(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	blockNumber := vars["blockNumber"]

	block, err := b.BlockQueries(dbase, db.GetBlock, blockNumber)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(block))
}

// GetAllBlocksWithoutLimit returns all blocks in table
func GetAllBlocksWithoutLimit(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	blocks, err := b.BlockArrayQueries(dbase, db.GetAllBlocksNoLimit, 0, 0, "")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(blocks))
}

// SGetAllBlocksLength Count all rows in Blocks Table
func SGetAllBlocksLength(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	count := b.RecordCountQuery(dbase, db.GetBlockCount)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(count))
}

// GetAllBlocks returns all blocks
func GetAllBlocks(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]
	page := common.StringToInteger(currentPage)
	limit := common.StringToInteger(pageLimit)

	blocks, err := b.BlockArrayQueries(dbase, db.GetAllBlocks, page, limit, "")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(blocks))
}

// GetRecentBlock returns most recent block height
func GetRecentBlock(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	block, err := b.BlockQueries(dbase, db.GetBlock, "")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(block))
}

// SGetAllInternalTransactionsLength Count all rows in Blocks Table
func SGetAllInternalTransactionsLength(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	count := b.RecordCountQuery(dbase, db.GetInternalTransactionLength)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.WriteLogger(w.Write(count))
}

//GetInternalTransactionsByHash gets internal txs specified by hash
func GetInternalTransactionsByHash(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	txHash := vars["txHash"]
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]

	page := common.StringToInteger(currentPage)
	limit := common.StringToInteger(pageLimit)

	transactions, err := tx.InternalTransactionArrayQuery(dbase, db.GetInternalTransaction, page, limit, txHash)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	logger.WriteLogger(w.Write(transactions))
}

//GetInternalTransactions gets internal txs
func GetInternalTransactions(w http.ResponseWriter, r *http.Request) {
	dbase := db.ConnectShyftDatabase()

	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]

	page := common.StringToInteger(currentPage)
	limit := common.StringToInteger(pageLimit)

	transactions, err := tx.InternalTransactionArrayQuery(dbase, db.GetAllInternalTransactions, page, limit, "")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	logger.WriteLogger(w.Write(transactions))
}

// SGetAllInternalTransactionsLength Count all rows in Blocks Table
func GetSearchQuery(w http.ResponseWriter, r *http.Request) {
	//dbase := db.ConnectShyftDatabase()
	vars := mux.Vars(r)
	query := vars["query"]
	//count := b.RecordCountQuery(dbase, db.GetInternalTransactionLength)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println("worked", query)
	fmt.Fprintf(w, "worked")
	//logger.WriteLogger(w.Write(count))
}

// BroadcastTx broadcasts tx
func BroadcastTx(w http.ResponseWriter, r *http.Request) {
	// Example return result (returns tx hash):
	// {"jsonrpc":"2.0","id":1,"result":"0xafa4c62f29dbf16bbfac4eea7cbd001a9aa95c59974043a17f863172f8208029"}

	// http params
	vars := mux.Vars(r)
	transactionHash := vars["transaction_hash"]

	// format the transactionHash into a proper sendRawTransaction jsonrpc request
	formattedJSON := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["%s"],"id":0}`, transactionHash))

	// send json rpc request
	resp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer(formattedJSON))
	if err != nil {
		logger.Warn("Error: " + err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warn("Error: " + err.Error())
	}
	byt := []byte(string(body))

	// read json and return result as http response, be it an error or tx hash
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	txHash := dat["result"]
	if txHash == nil {
		errMap := dat["error"].(map[string]interface{})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintln(w, "ERROR:", errMap["message"]); err != nil {
			logger.Warn("Error: " + err.Error())
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintln(w, "Transaction Hash:", txHash); err != nil {
			logger.Warn("Error: " + err.Error())
		}
	}
}
