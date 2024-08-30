package main

import (
	"fmt"
)

func convertToTitle(columnNumber int) string {
	var answer string
	newNumber := columnNumber

	if columnNumber <= 26 {
		return string(columnNumber + 64)
	}

	for {
		floatValue := float64((newNumber) % 26)
		if floatValue == 0 {
			floatValue = 26
		}
		answer = string(int((floatValue))+64) + answer
		if floatValue == 26 {
			newNumber = (newNumber / 26) - 1
		} else {
			newNumber = newNumber / 26
		}
		if newNumber <= 26 {
			answer = string(newNumber+64) + answer
			break
		}
	}

	return string(answer)
}

func main() {
	n := 1
	for n < 1000 {
		fmt.Println(convertToTitle(n))
		n++
	}
}
