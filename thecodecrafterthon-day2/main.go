package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to d's Base convanter")

initial:

	fmt.Println("enter base to convert: (e.g hex bin dec)")
	fmt.Println("Incase you need more help")
	fmt.Println("enter help")
	var converter string

	var conv string

	fmt.Scan(&conv)

	switch conv {

	case "help":
		fmt.Println("hex -> dec")
		fmt.Println("bin -> dec")
		fmt.Println("dec -> bin and hex")
		goto initial
	case "quit":
		fmt.Println("thank you for your time and goodbye")
		return
	case "bin", "hex", "dec":

	default:
		fmt.Println("invalid base")
		goto initial

	}

	fmt.Println("enter number to convert")
	fmt.Scan(&converter)

	switch strings.ToLower(conv) {
	case "hex":

		conn, err := strconv.ParseInt(converter, 16, 64)
		if err != nil {
			fmt.Println("invalid hexadecimal number")
		} else {
			fmt.Println("result", strconv.FormatInt(conn, 10))
		}
		goto initial

	case "bin":

		output, err := strconv.ParseInt(converter, 2, 64)
		if err != nil {
			fmt.Println("invalide binary number")
		} else {
			fmt.Println("result", strconv.FormatInt(output, 10))
		}

		goto initial
	case "dec":

		convy, err := strconv.ParseInt(converter, 10, 64)
		if err != nil {
			fmt.Println("invalid decimal number")
			goto initial
		}
		fmt.Println("result", strconv.FormatInt(convy, 2))
		fmt.Println("result", strings.ToUpper(strconv.FormatInt(convy, 16)))
		goto initial
	}

}
