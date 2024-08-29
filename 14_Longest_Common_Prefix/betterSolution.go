package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, s := range strs[1:] {

		for len(s) < len(prefix) || s[:len(prefix)] != prefix {

			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}

		}

	}

	return prefix
}

func main() {
	strs := []string{
		"flower",
		"flow",
		"flight",
	}
	fmt.Println(longestCommonPrefix(strs))
}