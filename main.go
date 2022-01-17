package main

import (
	"fmt"
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

// func main() {
// 	hextodec("")

// }

// function to convert (bin) to dec

func main() {
	bin_num := "111"

	num, err := strconv.ParseInt(bin_num, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("decimal num: ", num)
}

// Every instance of (up) converts the word placed before in the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !")

func ToUpper(s string) string {
	characters := []rune(s)
	for i := 0; i < len(characters); i++ {
		currentLetter := characters[i]
		if currentLetter >= 'a' && currentLetter <= 'z' {
			characters[i] = characters[i] - 32
		}
	}
	return string(characters)
}

// Every instance of (low) converts the word placed before in the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")

func ToLower(s string) string {
	characters := []rune(s)
	for i := 0; i < len(characters); i++ {
		currentLetter := characters[i]
		if currentLetter >= 'A' && currentLetter <= 'Z' {
			characters[i] = characters[i] + 32
		}
	}
	return string(characters)
}

// Every instance of (cap) transforms the previous word in the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")
func capitalise(s string) string {
	runes := []rune(s)
	strlen := 0
	for i := range runes {
		strlen = i + 1
	}
	for i := 0; i < strlen; i++ {
		if i != 0 && (!((runes[i-1] >= 'a' && runes[i-1] <= 'z') || (runes[i-1] >= 'A' && runes[i-1] <= 'Z'))) {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else if i == 0 {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = rune(runes[i] - 32)
			}
		} else {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = rune(runes[i] + 32)
			}
		}
	}
	return string(runes)
}
