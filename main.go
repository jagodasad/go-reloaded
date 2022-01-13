package main

import (
	"fmt"
	"strconv"
)

func main() {

	hexa := "F1"

	deciman, err := strconv.ParseInt(hexa, 16, 32)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Hexadecimal %s (base16) = Decimal %v (base10", hexa, decimal)
}
