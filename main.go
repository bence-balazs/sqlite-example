package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFile = "sqlite-database.db"
)

//go:embed sql/expense_information.sql
var expenseQuery string

//go:embed sql/create_expense_table.sql
var createTableExpense string

//go:embed sql/insert_into_expense.sql
var insertIntoExpense string

func main() {

	// create database if not exists
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		file, err := os.Create(dbFile)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println(dbFile, "created")
	}

	sqliteDatabase, _ := sql.Open("sqlite3", dbFile)
	defer sqliteDatabase.Close()

	if !tableCheck(sqliteDatabase, "expense") {
		createTable(sqliteDatabase, "expense")
	}

	instertExpense(sqliteDatabase, 15, "alma")

	// displayExpenses(sqliteDatabase)
	displayExpensesQuery(sqliteDatabase)

}

func createTable(db *sql.DB, tableName string) {
	log.Println("Create table..." + tableName)
	statement, err := db.Prepare(createTableExpense)

	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("expense table created")
}

func instertExpense(db *sql.DB, value int, title string) {
	statement, err := db.Prepare(insertIntoExpense)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(value, title)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayExpensesQuery(db *sql.DB) {
	row, err := db.Query(expenseQuery)

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var id int
		var value int
		var title string
		row.Scan(&id, &value, &title)
		fmt.Println("id:", id, value, ".-", title)
	}
}
func tableCheck(db *sql.DB, table string) bool {
	tableCheck, _ := db.Query("select * from " + table + ";")
	return tableCheck != nil
}
