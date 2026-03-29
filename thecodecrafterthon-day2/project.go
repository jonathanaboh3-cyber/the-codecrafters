package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	fmt.Println("[1]hex")
	fmt.Println("[2]bin")
	fmt.Println("[3]dec")
	fmt.Scan(&input)
	input, err := strconv.Atoi("input")
		if err != nil {
     fmt.Println("Invalid input")
	}
}
