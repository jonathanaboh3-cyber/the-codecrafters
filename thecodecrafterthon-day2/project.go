
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("--------- Welcome To Jonathan's Converting System-------")

    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }

        input := strings.TrimSpace(scanner.Text())
        if input == "" {
            continue
        }

        if input == "quit" {
            break
        }

        parts := strings.Fields(input)
        if len(parts) != 3 || parts[0] != "convert" {
            fmt.Println("Invalid input. Use: convert <value> <base>c")
            continue
        }

        value := parts[1]
        baseStr := parts[2]

        var base int
        switch baseStr {
        case "bin":
            base = 2
        case "dec":
            base = 10
        case "hex":
            base = 16
        default:
            fmt.Println("Unknown base. Use bin, dec, or hex.")
            continue
        }

        num, err := strconv.ParseInt(value, base, 64)
        if err != nil {
            fmt.Println("Invalid number for base:", baseStr)
            continue
        }

        switch baseStr {
        case "dec":
            fmt.Println("✦ Binary: ", strconv.FormatInt(num, 2))
            fmt.Println("✦ Hex:    ", strings.ToUpper(strconv.FormatInt(num, 16)))
        default:
            fmt.Println("✦ Decimal:", num)
        }
    }
}

