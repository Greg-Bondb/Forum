package main

import (
	"database/sql"
	"fmt"
	"log"
)

//conect to db
func connectDB() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/Forum")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}
