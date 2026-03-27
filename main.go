package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("----------WELCOME TO JONATHAN'S CALCULATION SYSTEM----------")
	for {
		var quit string
		
		fmt.Println("type quit to exit or type continue for the rest of tne operations")
		
		fmt.Scan(&quit)
		if quit == "quit" {
			fmt.Println("Goodbye Codecrafters! ")
			break
		} else {
			fmt.Println("continue operations")
			goto start
		}

	start:
		var input string
		fmt.Println("enter first index")
		fmt.Scan(&input)
		num1, err := strconv.Atoi(input)
		if err != nil {
		fmt.Println("only numbers allowed")
		goto start
		}

		var operator string
		fmt.Println("enter an operator: +, -, /, *,")
		fmt.Scan(&operator)
		if operator != ("+, *, -, /") {
			fmt.Println("invalid operator")
		}

		var index string
		fmt.Println("enter second index")
		fmt.Scan(&index)
		num2, err := strconv.Atoi(index)
		if err != nil {
		fmt.Println("only numbers allowed")
		goto start
		}

		if operator == "+" {
			result := (num1 + num2)
			fmt.Println(result)
		}
		if operator == "-" {
			result := (num1 - num2)
			fmt.Println(result)
		}
		if operator == "*" {
			result := (num1 * num2)
			fmt.Println(result)
		}

		if operator == "/" {
			if num1 == 0 || num2 == 0 {
				fmt.Println("syntax Error")
			} else {
				fmt.Println("Result", num1/num2)
			}

		}
		
		

	}

}
