package main

import "fmt"

const gramNum = 3

func main() {
	fmt.Println("Next-word prediction using Markov's Chain.")
	fmt.Println("1. Train on new data.")
	fmt.Println("2. Predict based on trained database.")
	fmt.Print("Give me your choice: ")

	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println()
		fmt.Println("Input the name of the database. Using an existing database would train the model based on existing data.")
		var dbName string
		fmt.Scanln(&dbName)
		openDatabase(dbName)
		startTraining()
		defer closeDatabase()

	} else if choice == 2 {
		fmt.Println()
		fmt.Println("Name of the dataset to be used: ")
		var dbName string
		fmt.Scanln(&dbName)
		openDatabase(dbName)
		defer closeDatabase()

	} else {
		panic("Wrong input. Panicking.")
	}
}
