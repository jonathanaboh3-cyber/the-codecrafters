package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("----------WELCOME TO JONATHAN'S CALCULATION SYSTEM----------")
	for {
		var quit string

		fmt.Println("type quit to exit or type continue to goahead or type help to show menu")

		fmt.Scan(&quit)
		if quit == "quit" {
			fmt.Println("Goodbye Codecrafters! ")
			break
		}
		if quit == "help" {

			fmt.Println("For Addition: index, operation, index. eg (1 + 1)")
			fmt.Println("For Subtraction: index, operation, index. eg (1 - 1)")
			fmt.Println("For Division: index, operation, index. eg (1 / 1)")
			fmt.Println("For Multiplication: index, operation, index. eg (1 * 1)")

			goto start
		}

	start:
		var input string
		fmt.Println("enter first index")
		fmt.Scan(&input)
		num1, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input only want digit", err)
			continue
		}

		var operator string
		fmt.Println("enter an operator: +, -, /, *,")
		fmt.Scan(&operator)
		if operator == "+" || operator == "*" || operator == "-" || operator == "/" {
			fmt.Println("let continue operator")
		} else if operator != "+" || operator != "*" || operator != "-" || operator != "/" {
			fmt.Println("Invalid operation")
			goto start
		}

		var index string
		fmt.Println("enter second index")
		fmt.Scan(&index)
		num2, err := strconv.Atoi(index)
		if err != nil {
			fmt.Println("Invalid input only want digit", err)
			continue
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
