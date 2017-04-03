package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startTraining() {
	fmt.Println()
	fmt.Println("1. Manual Input")
	fmt.Println("2. Read from file")
	fmt.Print("Give me your choice: ")

	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		trainOnSentence()
	} else if choice == 2 {

	} else {
		panic("Wrong input. Panicking.")
	}
}

func trainOnSentence() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Training sentence (type nothing to stop) (You need EOF): ")
		done, trained := processString(reader)
		if done {
			break
		}
		if trained {
			fmt.Println("Trained.")
		}
	}
}

func processString(r *bufio.Reader) (done, trained bool) {
	list := make([]string, gramNum)
	index := 1

	var err error
	list[0], err = r.ReadString(' ')
	list[0] = strings.TrimSpace(list[0])
	if err != nil {
		fmt.Println("Halted.")
		return true, false
	}

	//Fill the array first
	for ; err == nil && index < gramNum; index++ {
		list[index], err = r.ReadString(' ')
		list[index] = strings.TrimSpace(list[index])
	}

	if index < gramNum {
		//Not enough item filled
		fmt.Println("Not enough items in string.")
		return false, false
	}

	//Start processing
	for err == nil {

		for i := 1; i < gramNum; i++ {
			list[i-1] = list[i]
		}

		list[gramNum-1], err = r.ReadString(' ')
		list[gramNum-1] = strings.TrimSpace(list[gramNum-1])
		fmt.Println(list)
	}

	return false, true
}
