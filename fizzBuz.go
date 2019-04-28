package main

import (
	"fmt"
	"strconv"
)

func main() {
	var userInput int
	fmt.Print("Input max value: ")
	_,err := fmt.Scanf("%d", &userInput)

	if err == nil { //If Scanf got value without errors
		for x:=0; x<userInput; x++ {
			fmt.Println(x, "\t", multiplicityCheck(x))
		}
	}
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