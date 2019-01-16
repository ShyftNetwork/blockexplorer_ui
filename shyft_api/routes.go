package main

//@NOTE Shyft setting up endpoints
import "net/http"

//Route stuct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes routes
type Routes []Route

var routes = Routes{
	Route{
		"GetAccount",
		"GET",
		"/api/get_account/{address}",
		GetAccount,
	},
	Route{
		"GetAccountTxs",
		"GET",
		"/api/get_account_txs/{address}",
		GetAccountTxs,
	},
	Route{
		"GetAllAccounts",
		"GET",
		"/api/get_all_accounts",
		GetAllAccounts,
	},
	Route{
		"GetAllBlocksWithoutLimit",
		"GET",
		"/api/get_all_blocks_nolimit",
		GetAllBlocksWithoutLimit,
	},
	Route{
		"GetAllBlocks",
		"GET",
		"/api/get_all_blocks/{limit}/{offset}",
		GetAllBlocks,
	},
	Route{
		"GetAllBlocksLength",
		"GET",
		"/api/get_all_blocks_length",
		GetAllBlocksLength,
	},
	Route{
		"GetBlock",
		"GET",
		"/api/get_block/{blockNumber}",
		GetBlock,
	},
	Route{
		"GetAllTransactions",
		"GET",
		"/api/get_all_transactions",
		GetAllTransactions,
	},
	Route{
		"GetTransaction",
		"GET",
		"/api/get_transaction/{txHash}",
		GetTransaction,
	},
	Route{
		Name:        "GetRecentBlock",
		Method:      "GET",
		Pattern:     "/api/get_recent_block",
		HandlerFunc: GetRecentBlock,
	},
	Route{
		Name:        "GetAllTransactionsFromBlock",
		Method:      "GET",
		Pattern:     "/api/get_all_transactions_from_block/{blockNumber}",
		HandlerFunc: GetAllTransactionsFromBlock,
	},
	Route{
		Name:        "GetAllBlocksMinedByAddress",
		Method:      "GET",
		Pattern:     "/api/get_blocks_mined/{coinbase}",
		HandlerFunc: GetAllBlocksMinedByAddress,
	},
	Route{
		"GetInternalTransactions",
		"GET",
		"/api/get_internal_transactions/",
		GetInternalTransactions,
	},
	Route{
		"GetInternalTransactionsByHash",
		"GET",
		"/api/get_internal_transactions/{txHash}",
		GetInternalTransactionsByHash,
	},
	Route{
		"BroadcastTx",
		"GET",
		"/api/broadcast_tx/{transaction_hash}",
		BroadcastTx,
	},
}
