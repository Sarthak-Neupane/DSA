package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	l := 0
	subString := ""

	if len(needle) > len(haystack) {
		return -1
	}

	if haystack == needle {
		return 0
	}

	if len(needle) == 1 {
		for i := 0; i < len(haystack); i++ {
			if string(haystack[i])== needle {
				return i
			}
		}
	}

	for r := 1; (r - l) < len(needle); r++ {
		if l == len(haystack) || r == len(haystack) {
			break
		}
		if haystack[l] != needle[0] && len(subString) == 0 {
			l++
			r = l
			continue
		}
		if subString == "" {
			subString += string(needle[0])
		}
		fmt.Println("Currently checking ", "left:", l, "right:", r, "needle char position:", r-l)
		if haystack[r] == needle[r-l] {
			subString += string(needle[r-l])
			fmt.Println(subString, "currently added is ", string(needle[r-l]))
			if subString == needle {
				return l
			}
		} else {
			fmt.Println("Milena. Current Letter namileko", string(needle[r-l]))
			subString = ""
			l++
			r = l
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("aabaaabaaac", "aabaaac"))
}
