package main

import "fmt"
import "os"
import "strconv"
import "io/ioutil"

const gramNum = 4

func main() {
	fmt.Println("Next-word prediction using Markov's Chain.")

	arguments := os.Args[1:]

	if len(arguments) < 3 {
		//At least operation, database name and data set
		fmt.Println("Usage: [Operation] [Database] [Extra]")
		os.Exit(0)
	}

	operation := arguments[0]
	dbName := arguments[1]
	openDatabase(dbName)
	defer closeDatabase()

	if operation == "train" {
		allFiles := arguments[2:]
		for _, f := range allFiles {
			trainFile(f)
		}

	} else if operation == "write" {
		length, err := strconv.Atoi(arguments[2])
		if len(arguments) < 4 || err != nil {
			fmt.Println("Usage: [Operation] [Database] [Length] [Output]")
		}
		result := writeForLength(length)
		fmt.Println(result)
		ioutil.WriteFile(arguments[3], []byte(result), 0600)

	} else if operation == "predict" {

	} else if operation == "compare" {

	} else {
		fmt.Println("Operation " + operation + " not supported.")
		os.Exit(0)
	}
}
