package main

import (
	"fmt"
)

func main() {
	lengthOfLongestSubstring("aab")
}

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	} 

	substrings := make(map[byte]bool)
	l := 0
	maxLength := 0

	for r := 0; r <= len(s) - 1; r++ {
		for substrings[s[r]] {
			delete(substrings, s[l])
			l++;
		}

		substrings[s[r]] = true
		if (r-l)+1 >= maxLength {
			maxLength = (r - l) + 1
		}
	}

	fmt.Println(maxLength)
	return 0
}
