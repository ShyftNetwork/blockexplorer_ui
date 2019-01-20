package db

// GetAllBlocksNoLimit returns all blocks from blocks table
const GetAllBlocksNoLimit = `SELECT * FROM blocks ORDER BY number ASC`
// GetBlockCount returns a count of all rows in block table
const GetBlockCount = `SELECT COUNT(*) FROM blocks`
// GetAllBlocks returns all blocks based on offset limit
const GetAllBlocks = `SELECT * FROM blocks ORDER BY number ASC LIMIT $1 OFFSET $2`
// GetBlock returns specified block height
const GetBlock = `SELECT * FROM blocks WHERE number=$1;`
// GetRecentBlock returns the latest block mined
const GetRecentBlock = `SELECT * FROM blocks WHERE number=(SELECT MAX(number) FROM blocks);`
// GetAllBlocksMinedByAddress returns blocks moned by specified address
const GetAllBlocksMinedByAddress = `SELECT * FROM blocks WHERE coinbase=$3 ORDER BY age ASC LIMIT $1 OFFSET $2;`

// GetTransactionCount returns count of all txs from tx table
const GetTransactionCount = `SELECT COUNT(*) FROM txs`
// GetAllTransactionsFromBlock returns all txs from specified block height
const GetAllTransactionsFromBlock = `SELECT * FROM txs WHERE blocknumber=$3 ORDER BY age ASC LIMIT $1 OFFSET $2;`
// GetAllTransactionsNoLimit returns all txs
const GetAllTransactionsNoLimit = `SELECT * FROM txs ORDER BY age ASC`
// GetAllTransactions returns all txs based on offset and limit
const GetAllTransactions = `SELECT * FROM txs ORDER BY age ASC LIMIT $1 OFFSET $2`
// GetTransaction returns specified tx
const GetTransaction = `SELECT * FROM txs WHERE txhash=$1;`

// GetAllInternalTransactionsNoLimit returns all internal txs
const GetAllInternalTransactionsNoLimit = `SELECT * FROM internaltxs ORDER BY age ASC`
// GetAllInternalTransactions returns all internal txs based on offset limit
const GetAllInternalTransactions = `SELECT * FROM internaltxs ORDER BY age ASC LIMIT $1 OFFSET $2`
// GetInternalTransaction returns internal txs based on hash
const GetInternalTransaction = `SELECT * FROM internaltxs WHERE txhash=$3 LIMIT $1 OFFSET $2;`
// GetInternalTransactionLength returns count of rows in internal tx table
const GetInternalTransactionLength = `SELECT COUNT(*) FROM internaltxs`

// GetAllAccountsNoLimit returns all accounts
const GetAllAccountsNoLimit = `SELECT * FROM accounts ORDER BY balance ASC`
// GetAccountCount returns count of rows in accounts table
const GetAccountCount = `SELECT COUNT(*) FROM accounts`
// GetAllAccounts returns all accounts based on offset limit
const GetAllAccounts = `SELECT * FROM accounts LIMIT $1 OFFSET $2`
// GetAccount returns specified account
const GetAccount = `SELECT * FROM accounts WHERE addr=$1;`
// GetAccountTransactions returns all txs from specified account address
const GetAccountTransactions = `SELECT * FROM txs WHERE to_addr=$3 OR from_addr=$3 ORDER BY age ASC LIMIT $1 OFFSET $2;`
// GetAllAccountBlocks returns all account blocks
const GetAllAccountBlocks = `SELECT * FROM accountblocks`

// SearchQuery returns all relevant data based on specified search
const SearchQuery = `SELECT * FROM txs WHERE to_addr=$1 OR from_addr=$1 OR blockhash=$1 OR txhash=$1`