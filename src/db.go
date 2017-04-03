package main

import (
	"fmt"

	"os"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func openDatabase(name string) {
	os.Mkdir("/db", 0700)

	var err error
	db, err = bolt.Open("/db/"+name+".bolt", 0600, nil)
	if err != nil {
		fmt.Println("Error in opening database: ", name)
		panic(err)
	}
}

func closeDatabase() {
	db.Close()
}
