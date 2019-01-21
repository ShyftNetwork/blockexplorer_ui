package api

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

// Endpoints outline all endpoints for api calls
var Endpoints = Routes{
	Route{
		"GetAllAccountLength",
		"GET",
		"/api/get_all_accounts_length",
		SGetAllAccountsLength,
	},
	Route{
		"GetAccount",
		"GET",
		"/api/get_account/{address}",
		GetAccount,
	},
	Route{
		"GetAccountTxs",
		"GET",
		"/api/get_account_txs/{currentPage}/{pageLimit}/{address}",
		GetAccountTxs,
	},
	Route{
		"GetAllAccounts",
		"GET",
		"/api/get_all_accounts/{currentPage}/{pageLimit}",
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
		"/api/get_all_blocks/{currentPage}/{pageLimit}",
		GetAllBlocks,
	},
	Route{
		"GetAllBlocksLength",
		"GET",
		"/api/get_all_blocks_length",
		SGetAllBlocksLength,
	},
	Route{
		"GetBlock",
		"GET",
		"/api/get_block/{blockNumber}",
		GetBlock,
	},
	Route{
		"GetAllTransactionsWithoutLimit",
		"GET",
		"/api/get_all_transactions_nolimit",
		GetAllTransactionsWithoutLimit,
	},
	Route{
		"GetAllTransactionsLength",
		"GET",
		"/api/get_all_transactions_length",
		SGetAllTransactionsLength,
	},
	Route{
		"GetAllTransactions",
		"GET",
		"/api/get_all_transactions/{currentPage}/{pageLimit}",
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
		Pattern:     "/api/get_all_transactions_from_block/{currentPage}/{pageLimit}/{blockNumber}",
		HandlerFunc: GetAllTransactionsFromBlock,
	},
	Route{
		Name:        "GetAllBlocksMinedByAddress",
		Method:      "GET",
		Pattern:     "/api/get_blocks_mined/{currentPage}/{pageLimit}/{coinbase}",
		HandlerFunc: GetAllBlocksMinedByAddress,
	},
	Route{
		"GetInternalTransactions",
		"GET",
		"/api/get_internal_transactions/{currentPage}/{pageLimit}",
		GetInternalTransactions,
	},
	Route{
		"GetInternalTransactionsByHash",
		"GET",
		"/api/get_internal_transactions/{currentPage}/{pageLimit}/{txHash}",
		GetInternalTransactionsByHash,
	},
	Route{
		"GetAllInternalTransactionsLength",
		"GET",
		"/api/get_internal_transactions_length",
		GetAllInternalTransactionsLength,
	},
	Route{
		"BroadcastTx",
		"GET",
		"/api/broadcast_tx/{transaction_hash}",
		BroadcastTx,
	},
	Route{
		"GetSearchQuery",
		"GET",
		"/api/search/{query}",
		GetSearchQuery,
	},
}
