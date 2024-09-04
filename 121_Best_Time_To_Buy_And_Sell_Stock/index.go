package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	length := len(prices)
	profit := 0

	if length == 1 {
		return 0
	}

	if length == 2 {
		if prices[1] - prices[0] > profit {
			return prices[1] - prices[0]
		}
		return profit
	}

	l := 0
	r := 1

	for r <= length - 1 {
		if prices[r]-prices[l] > profit {
			profit = prices[r] - prices[l]
		}
		if prices[r] < prices[l] {
			l = r
		}
		r++
	}

	return profit
}

func main() {
	prices := []int{
		1,2,4,
	}
	fmt.Println(maxProfit(prices))
}
