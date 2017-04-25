package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var gramNum int

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Next-word prediction using Markov's Chain.")

	arguments := os.Args[1:]

	if len(arguments) < 4 {
		//At least operation, database name and data set
		fmt.Println("Usage: [Operation] [N-gram] [Database] [Extra]")
		os.Exit(0)
	}

	operation := arguments[0]
	var err error
	gramNum, err = strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Usage: [Operation] [N-gram] [Database] [Extra]")
		os.Exit(0)
	}

	dbName := arguments[2]
	openDatabase(dbName)
	defer closeDatabase()

	if operation == "train" {
		allFiles := arguments[3:]
		for _, f := range allFiles {
			trainFile(f)
		}

	} else if operation == "write" {
		length, err := strconv.Atoi(arguments[3])
		if len(arguments) < 5 || err != nil {
			fmt.Println("Usage: [Operation] [N-gram] [Database] [Length] [Output]")
		}
		result := writeForLength(length)
		fmt.Println(result)
		ioutil.WriteFile(arguments[4], []byte(result), 0600)

	} else if operation == "predict" {

	} else if operation == "compare" {

	} else {
		fmt.Println("Operation " + operation + " not supported.")
		os.Exit(0)
	}
}
