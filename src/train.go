package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

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
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
	}
}

func trainOnSentence() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Training sentence (type nothing to stop): ")
		s, err := reader.ReadBytes('\n')
		if err != nil || len(s) <= 1 {
			fmt.Println("Halt.")
			break
		}

		if processString(s) {
			fmt.Println("Trained.")
		}
	}
}

var textOnlyRegex = regexp.MustCompile("[^a-zA-Z]{2,}") //Matches non english characters, greedy
func processString(inputBytes []byte) (trained bool) {
	//Replace all with a single spaces
	inputString := string(textOnlyRegex.ReplaceAll(inputBytes, []byte(" ")))

	wordList := strings.Split(inputString, " ")

	if len(wordList) < gramNum {
		fmt.Println("Not enough words in string.")
		return false
	}

	err := myBoltDB.Update(func(tx *bolt.Tx) error {
		for i := 0; i < len(wordList)-gramNum; i++ {
			//Get the first bucket
			currBucket, err := tx.CreateBucketIfNotExists([]byte(dbBucketName))

			if err != nil {
				return err
			}

			for j := 0; j < gramNum-1; j++ {
				//Nest into the deepest bucket
				currBucket, err = currBucket.CreateBucketIfNotExists([]byte(wordList[i+j]))
				if err != nil {
					return err
				}
			}

			gotByteArray := currBucket.Get([]byte(wordList[i+gramNum-1]))
			if gotByteArray == nil {
				currBucket.Put([]byte(wordList[i+gramNum-1]), intToByteArray(1))
			} else {
				currBucket.Put([]byte(wordList[i+gramNum-1]), intToByteArray(byteArrayToInt(gotByteArray)+1))
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error in inserting words into database")
		panic(err)
	}

	return true
}
