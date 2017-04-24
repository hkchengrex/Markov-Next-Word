package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var myBoltDB *bolt.DB
var dbBucketName string

func openDatabase(name string) {
	var err error
	myBoltDB, err = bolt.Open(name, 0600, nil)
	if err != nil {
		fmt.Println("Error in opening database: ", name)
		panic(err)
	}

	myBoltDB.Update(func(tx *bolt.Tx) error {
		dbBucketName = fmt.Sprintf("%d-gram", gramNum)
		_, err := tx.CreateBucketIfNotExists([]byte(dbBucketName))
		return err
	})
}

func closeDatabase() {
	myBoltDB.Close()
}
