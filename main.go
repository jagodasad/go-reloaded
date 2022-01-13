package main

import (
	"strconv"
)

// function to convert hex to dec

func hextodec(hexi string) string {
	// convert string (hex) to dec
	num, err := strconv.ParseInt(hexi, 16, 64)
	if err != nil {
		panic(err)
		// panic(err), aborts a function and returns an error value. Prints error message
	}
	dec := strconv.Itoa(int(num))
	// convert int to string via 'Itoa'
	return dec
	// returning string
}

func main() {
	hextodec("")

}

// function to convert (bin) to dec

func bintodec(bin string) string {
	// 	// convert string (bin) to dec
	num, err := strconv.ParseInt(bin, 16, 64)
	if err != nil {
		panic(err)
		// 		// panic(err), aborts a function and returns an error value. Prints error message
	}
	dec := strconv.Itoa(int(num))
	// 	// convert int to string via 'Itoa'
	return dec
	// 	// returning string
}
