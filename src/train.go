package main

import (
	"fmt"
	"strings"

	"io/ioutil"

	"path/filepath"

	"github.com/boltdb/bolt"
)

func trainFile(filePattern string) {
	filePaths, err := filepath.Glob(filePattern)
	if err != nil {
		fmt.Println("Cannot process " + filePattern)
		return
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
	inputRunes := []rune(obtainStartOfText() + strings.Replace(string(inputBytes), "\r\n", "\n", -1) + obtainEndOfText())

	err := myBoltDB.Update(func(tx *bolt.Tx) error {
		rootBucket, err := tx.CreateBucketIfNotExists([]byte(dbBucketName))
		if err != nil {
			return err
		}

		for i := 0; i < len(inputRunes)-gramNum+1; i++ {
			destBucket, err := rootBucket.CreateBucketIfNotExists([]byte(string(inputRunes[i : i+gramNum-1])))
			if err != nil {
				return err
			}

			keyBytes := []byte(string(inputRunes[i+gramNum-1]))
			gotByteArray := destBucket.Get(keyBytes)
			if gotByteArray == nil {
				destBucket.Put((keyBytes), intToByteArray(1))
			} else {
				destBucket.Put((keyBytes), intToByteArray(byteArrayToInt(gotByteArray)+1))
			}
		}

		return nil
	})

	if err != nil {
		return false
	}
	return true
}
