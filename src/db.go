package main

import (
	"fmt"

	"os"

	"github.com/boltdb/bolt"
)

var myBoltDB *bolt.DB
var dbBucketName string

func openDatabase(name string) {
	os.Mkdir("/db", 0700)

	var err error
	myBoltDB, err = bolt.Open("/db/"+name+".bolt", 0600, nil)
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
