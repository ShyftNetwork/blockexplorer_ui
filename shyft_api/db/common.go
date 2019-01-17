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

func BlockArrayQueries(db *SPGDatabase, query string, currentPage int64, pageLimit int64) ([]byte, error) {
	switch {
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

func QueryType(query string, row *sqlx.Row, rows *sqlx.Rows) []byte {
	switch query {
	case "block":
		fmt.Println("BLOCK PURPOSE")
		b, _ := BlockMarshalling(row)
		return b
	case "blocks":
		fmt.Println("BLOCK PURPOSEsssssss")
		b, _ := BlockArrayMarshalling(rows)
		return b
	default:
		t := []byte("string")
		return t
	}
}

func ExplorerQueryArray(db *SPGDatabase, args ...string) []byte {
	switch {
	case len(args) == 5:
		currentPage, pageLimit := GetPageAndLimit(args)
		queryObject := SliceStringToString(args[4:5])
		query := SliceStringToString(args[:1])
		identifier := SliceStringToString(args[3:4])

		var offset = (currentPage - 1) * pageLimit

		row := db.Db.QueryRowx(query, pageLimit, offset, identifier)
		response := QueryType(queryObject, row, nil)
		fmt.Println("length = 5", args)
		return response
	case len(args) == 4:
		fmt.Println("length = 4", args)
		t := []byte("string")
		return t
	case len(args) == 3:
		fmt.Println("length = 3", args)
		t := []byte("string")
		return t
	case len(args) == 2:
		fmt.Println("length 2", args)
		t := []byte("string")
		return t
	case len(args) == 1:
		fmt.Println("length 1", args)
		t := []byte("string")
		return t
	default:
		fmt.Println("DEFAULT", args)
		t := []byte("string")
		return t
	}
}

// func BlockQueries(db *SPGDatabase, query string, identifier string) ([]byte, error) {
// 	tx, _ := db.Db.Begin()
// 	switch {
// 	case len(identifier) > 0:
// 		row := db.Db.QueryRow(query, identifier)
// 		b, _ := BlockMarshalling(row)
// 		tx.Commit()
// 		return b, nil
// 	default:
// 		row := db.Db.QueryRow(query)
// 		b, _ := BlockMarshalling(row)
// 		tx.Commit()
// 		return b, nil
// 	}
// }

// func TransactionArrayMarshalling(rows *sql.Rows) ([]byte, error) {

// }

func TransactionArrayQueries(args ...string) {
	// page, err := strconv.ParseInt(currentPage, 10, 64)
	// limit, err := strconv.ParseInt(pageLimit, 10, 64)
	fmt.Println("VARIADIC FUNCTION", args[2:3])

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
