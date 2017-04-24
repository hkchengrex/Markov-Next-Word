package main

import (
	"fmt"

	"io/ioutil"

	"path/filepath"

	"github.com/boltdb/bolt"
)

func trainFile(filePattern string) {
	filePaths, err := filepath.Glob(filePattern)
	if err != nil {
		panic(err)
	}

	for _, file := range filePaths {
		fmt.Print("Training " + file + "...")
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		trained := processString(data)
		if trained {
			fmt.Println(" Done. ")
		} else {
			fmt.Println(" Failed.")
		}
	}
}

func processString(inputBytes []byte) (trained bool) {
	//Replace all with a single spaces
	inputRunes := []rune(string(inputBytes))

	err := myBoltDB.Update(func(tx *bolt.Tx) error {
		for i := 0; i < len(inputRunes)-gramNum; i++ {
			//Get the first bucket
			currBucket, err := tx.CreateBucketIfNotExists([]byte(dbBucketName))

			if err != nil {
				return err
			}

			for j := 0; j < gramNum-1; j++ {
				//Nest into the deepest bucket
				currBucket, err = currBucket.CreateBucketIfNotExists([]byte(string(inputRunes[i+j])))
				if err != nil {
					return err
				}
			}

			gotByteArray := currBucket.Get([]byte(string(inputRunes[i+gramNum-1])))
			if gotByteArray == nil {
				currBucket.Put([]byte(string(inputRunes[i+gramNum-1])), intToByteArray(1))
			} else {
				currBucket.Put([]byte(string(inputRunes[i+gramNum-1])), intToByteArray(byteArrayToInt(gotByteArray)+1))
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return true
}
