package main

import "fmt"
import "os"

const gramNum = 3

func main() {
	fmt.Println("Next-word prediction using Markov's Chain.")

	arguments := os.Args[1:]

	if len(arguments) < 3 {
		//At least operation, database name and data set
		fmt.Println("Usage: [operation] [database] [dataset]")
		os.Exit(0)
	}

	operation := arguments[0]
	dbName := arguments[1]
	openDatabase(dbName)
	defer closeDatabase()

	if operation == "train" {
		startTraining()

	} else if operation == "predict" {

	} else if operation == "compare" {

	} else {
		fmt.Println("Operation " + operation + " not supported.")
		os.Exit(0)
	}
}
