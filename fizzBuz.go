package main

import (
	"fmt"
	"strconv"
)

func main() {
	var userInput int
	fmt.Print("Input max value: ")
	_,err := fmt.Scanf("%d", &userInput)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i:=0; i<userInput; i++ {
		fmt.Println(i, "\t", multiplicityCheck(i))
	}

	return
}

func multiplicityCheck(x int) string {
	var resultString string

	if x%3 == 0 {
		resultString += "Fizz"
	}

	if x%5 == 0 {
		resultString += "Buzz"
	}

	if resultString == "" {
		resultString = strconv.Itoa(x)
	}

	return resultString
}