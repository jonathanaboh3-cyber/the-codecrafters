package main
import (
	"fmt"
	"strconv"
)
func main() {
	var convert string
	fmt.Println("pick a base to convert")
    fmt.Scan(&convert)

	value, err := strconv.ParseInt(convert, 16, 64)

	if err != nil {
	  fmt.Println("invalid")
    }
	 fmt.Println (strconv.FormatInt(value, 10))

	var base string

	fmt.Println("enter a base:")
	fmt.Scan(&base)

	if base == "hex" || base == "bin" || base == "dec" {
		fmt.Println("let continue operator")
	}
}