package main

import (
	"database/sql"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	err := db.Ping()
	if err != nil {
		t.Errorf("wanted nil, got err %v", err)
	}

}

func TestTableCreation(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	createTable(db, "expense")

	got := tableCheck(db, "expense")

	if got != true {
		t.Errorf("no table found")
	}
}

func TestInsert(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	createTable(db, "expense")
	instertExpense(db, 1, "asd")

	var got int
	err := db.QueryRow("select count(*) from expense;").Scan(&got)
	if err != nil {
		t.Fatal(err)
	}

	want := 1

	if got != want {
		t.Errorf("got %v but wanted %v rows", got, want)
	}
}
