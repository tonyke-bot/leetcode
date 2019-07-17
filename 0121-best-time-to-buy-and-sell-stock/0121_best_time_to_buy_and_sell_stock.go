package problem0121

// Source: https://leetcode.com/problems/best-time-to-buy-and-sell-stock

func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	minPrice := make([]int, length)
	maxPrice := make([]int, length)

	minPrice[0] = prices[0]
	maxPrice[length-1] = prices[length-1]

	for i := 1; i < length; i++ {
		if prices[i] < minPrice[i-1] {
			minPrice[i] = prices[i]
		} else {
			minPrice[i] = minPrice[i-1]
		}

		lastIndex := length - 1 - i
		if prices[lastIndex] > maxPrice[lastIndex+1] {
			maxPrice[lastIndex] = prices[lastIndex]
		} else {
			maxPrice[lastIndex] = maxPrice[lastIndex+1]
		}
	}

	maxDelta := 0
	for i := 0; i < length; i++ {
		delta := maxPrice[i] - minPrice[i]
		if delta > maxDelta {
			maxDelta = delta
		}
	}
	return maxDelta
}

/*
Test Case:
[7,1,5,3,6,4]
*/
