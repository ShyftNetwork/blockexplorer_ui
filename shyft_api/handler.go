package main

///@NOTE Shyft handler functions when endpoints are hit
import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ShyftNetwork/blockexplorer_ui/shyft_api/db"
	_ "github.com/lib/pq"

	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
)

// GetTransaction gets txs
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txHash := vars["txHash"]
	getTxResponse, err := db.SGetTransaction(txHash)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, getTxResponse)
}

// GetAllTransactions gets txs
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	txs, err := db.SGetAllTransactions()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, txs)
}

// GetAllTransactions gets txs
func GetAllTransactionsFromBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNumber := vars["blockNumber"]
	txsFromBlock, err := db.SGetAllTransactionsFromBlock(blockNumber)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, txsFromBlock)
}

func GetAllBlocksMinedByAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coinbase := vars["coinbase"]

	blocksMined, err := db.SGetAllBlocksMinedByAddress(coinbase)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, blocksMined)
}

// GetAccount gets balance
func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	getAccountBalance, err := db.SGetAccount(address)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, getAccountBalance)
}

// GetAccount gets balance
func GetAccountTxs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	getAccountTxs, err := db.SGetAccountTxs(address)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, getAccountTxs)
}

// GetAllAccounts gets balances
func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	allAccounts, err := db.SGetAllAccounts()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, allAccounts)
}

//GetBlock returns block json
func GetBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNumber := vars["blockNumber"]

	getBlockResponse, err := db.SGetBlock(blockNumber)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, getBlockResponse)
}

// GetAllBlocks response
func GetAllBlocksWithoutLimit(w http.ResponseWriter, r *http.Request) {
	blocks, err := db.SGetAllBlocksWithoutLimit()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, blocks)
}

// Count all rows in Blocks Table
func SGetAllBlocksLength(w http.ResponseWriter, r *http.Request) {
	count, err := db.GetAllBlocksLength()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, count)
}

func GetAllBlocks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currentPage := vars["currentPage"]
	pageLimit := vars["pageLimit"]
	page, err := strconv.ParseInt(currentPage, 10, 64)
	limit, err := strconv.ParseInt(pageLimit, 10, 64)

	blocks, err := db.SGetAllBlocks(page, limit)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, blocks)
}

func GetRecentBlock(w http.ResponseWriter, r *http.Request) {
	mostRecentBlock, err := db.SGetRecentBlock()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, mostRecentBlock)

}

//GetInternalTransactions gets internal txs
func GetInternalTransactionsByHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txHash := vars["txHash"]

	internalTxs, err := db.SGetInternalTransaction(txHash)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, internalTxs)
}

//GetInternalTransactionsHash gets internal txs hash
func GetInternalTransactions(w http.ResponseWriter, r *http.Request) {
	internalTxs, err := db.SGetAllInternalTransactions()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, internalTxs)
}

func BroadcastTx(w http.ResponseWriter, r *http.Request) {
	// Example return result (returns tx hash):
	// {"jsonrpc":"2.0","id":1,"result":"0xafa4c62f29dbf16bbfac4eea7cbd001a9aa95c59974043a17f863172f8208029"}

	// http params
	vars := mux.Vars(r)
	transactionHash := vars["transaction_hash"]

	// format the transactionHash into a proper sendRawTransaction jsonrpc request
	formatted_json := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["%s"],"id":0}`, transactionHash))

	// send json rpc request
	resp, _ := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer(formatted_json))
	body, _ := ioutil.ReadAll(resp.Body)
	byt := []byte(string(body))

	// read json and return result as http response, be it an error or tx hash
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ERROR parsing json")
	}
	tx_hash := dat["result"]
	if tx_hash == nil {
		errMap := dat["error"].(map[string]interface{})
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ERROR:", errMap["message"])
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Transaction Hash:", tx_hash)
	}
}
