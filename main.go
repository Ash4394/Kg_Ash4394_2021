package main

import (
	"fmt"
	"os"
	"strconv"
)

func findingNumberOfDigits(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}

	return count
}

func ConvertNumberToWord(number int) string {

	var ConvertedNumber string
	if number == 1 {
		ConvertedNumber = "One"

	} else if number == 2 {
		ConvertedNumber = "Two"

	} else if number == 3 {
		ConvertedNumber = "Three"

	} else if number == 4 {
		ConvertedNumber = "Four"

	} else if number == 5 {
		ConvertedNumber = "Five"

	} else if number == 6 {
		ConvertedNumber = "Six"

	} else if number == 7 {
		ConvertedNumber = "Seven"

	} else if number == 8 {
		ConvertedNumber = "Eight"

	} else if number == 9 {
		ConvertedNumber = "Nine"

	} else if number == 0 {
		ConvertedNumber = "Zero"

	}

	return ConvertedNumber

}

func main() {

	StoreArray := []string{}
	var x int
	length := len(os.Args)

	for i, a := range os.Args[1:] {

		n, err := strconv.Atoi(a)

		if err != nil {
			panic(err)
		}

		Digits := findingNumberOfDigits(n)

		for n != 0 {
			x = n % 10
			StoreArray = append(StoreArray, ConvertNumberToWord(x))
			n /= 10

		}

		for Digits != 0 {
			fmt.Printf("%s", StoreArray[Digits-1])
			Digits -= 1
		}

		StoreArray = nil

		if i != length-2 {
			fmt.Print(", ")
		}

	}

}
