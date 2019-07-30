package problem0739

/*
Source: https://leetcode.com/problems/daily-temperatures
Test Case:
[73,74,75,71,69,72,76,73]
*/

func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))
	stack := make([]int, 0, len(T))

	for i, ele := range T {
		for len(stack) > 0 && ele > T[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return result
}

// =============== Naive Version: O(n*W) time, W=100-30

// const maxTemp = 100

// func dailyTemperatures(T []int) []int {
// 	result := make([]int, len(T))
// 	nextOccurrence := make([]int, maxTemp+1)

// 	if len(T) <= 1 {
// 		return result
// 	}

// 	lastIndex := len(T) - 1
// 	nextOccurrence[T[lastIndex]] = lastIndex

// 	for i := lastIndex - 1; i >= 0; i-- {
// 		closerIndex := int(math.MaxInt32)
// 		for j := T[i] + 1; j <= maxTemp; j++ {
// 			if nextOccurrence[j] != 0 && nextOccurrence[j] < closerIndex {
// 				closerIndex = nextOccurrence[j]
// 			}
// 		}

// 		if closerIndex != int(math.MaxInt32) {
// 			result[i] = closerIndex - i
// 		}
// 		nextOccurrence[T[i]] = i
// 	}

// 	return result
// }
