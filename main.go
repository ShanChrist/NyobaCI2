package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	if x := 1; x == 1 {
		fmt.Println("x is equal to 1")
	}

	a := 10
	b := 10
	log.Println(a + b)

	username := "budi"
	query := fmt.Sprintf("SELECT * FROM users WHERE username = %s", username)

	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

}
