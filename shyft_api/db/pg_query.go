package db

const GetAllBlocksNoLimit = `SELECT * FROM blocks ORDER BY number ASC`
const GetBlockCount = `SELECT COUNT(*) FROM blocks`
const GetAllBlocks = `SELECT * FROM blocks ORDER BY number ASC LIMIT $1 OFFSET $2`
const GetBlock = `SELECT * FROM blocks WHERE number=$1;`
const GetRecentBlock = `SELECT * FROM blocks WHERE number=(SELECT MAX(number) FROM blocks);`
const GetAllBlocksMinedByAddress = `SELECT * FROM blocks WHERE coinbase=$3 ORDER BY age ASC LIMIT $1 OFFSET $2;`

const GetTransactionCount = `SELECT COUNT(*) FROM txs`
const GetAllTransactionsFromBlock = `SELECT * FROM txs WHERE blocknumber=$3 ORDER BY age ASC LIMIT $1 OFFSET $2;`
const GetAllTransactionsNoLimit = `SELECT * FROM txs ORDER BY age ASC`
const GetAllTransactions = `SELECT * FROM txs ORDER BY age ASC LIMIT $1 OFFSET $2`
const GetTransaction = `SELECT * FROM txs WHERE txhash=$1;`

const GetAllInternalTransactionsNoLimit = `SELECT * FROM internaltxs ORDER BY age ASC`
const GetAllInternalTransactions = `SELECT * FROM internaltxs ORDER BY age ASC LIMIT $1 OFFSET $2`
const GetInternalTransaction = `SELECT * FROM internaltxs WHERE txhash=$3 LIMIT $1 OFFSET $2;`
const GetInternalTransactionLength = `SELECT COUNT(*) FROM internaltxs`

const GetAllAccountsNoLimit = `SELECT * FROM accounts ORDER BY balance ASC`
const GetAccountCount = `SELECT COUNT(*) FROM accounts`
const GetAllAccounts = `SELECT * FROM accounts LIMIT $1 OFFSET $2`
const GetAccount = `SELECT * FROM accounts WHERE addr=$1;`
const GetAccountTransactions = `SELECT * FROM txs WHERE to_addr=$3 OR from_addr=$3 ORDER BY age ASC LIMIT $1 OFFSET $2;`

const GetAllAccountBlocks = `SELECT * FROM accountblocks`
