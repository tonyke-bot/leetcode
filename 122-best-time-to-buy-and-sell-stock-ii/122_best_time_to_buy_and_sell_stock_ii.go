package main

// Source: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii

func maxProfitII(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}

	totalProfit := 0
	for index := 0; index < length; index++ {
		// find valley
		for index+1 < length && prices[index] > prices[index+1] {
			index++
		}

		valley := index
		// find next peek
		for index+1 < length && prices[index] < prices[index+1] {
			index++
		}

		if index == length {
			totalProfit += prices[index-1] - prices[valley]
		} else {
			totalProfit += prices[index] - prices[valley]
		}
	}

	return totalProfit
}

/*
Test Case:
[7,1,5,3,6,4]
*/
