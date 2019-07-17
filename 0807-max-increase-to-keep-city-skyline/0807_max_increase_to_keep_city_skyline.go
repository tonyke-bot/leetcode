package problem0807

/*
Source: https://leetcode.com/problems/max-increase-to-keep-city-skyline
Test Case:
[[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
*/

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	n := len(grid)
	m := len(grid[0])

	horizontal := make([]int, n)
	vertical := make([]int, m)

	for i := 0; i < n; i++ {
		max := grid[i][0]
		for j := 1; j < m; j++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		horizontal[i] = max
	}

	for i := 0; i < m; i++ {
		max := grid[0][i]
		for j := 1; j < n; j++ {
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
		vertical[i] = max
	}

	maxSum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			v := vertical[j]
			h := horizontal[i]
			if v > h {
				maxSum += h
			} else {
				maxSum += v
			}
		}
	}

	return maxSum
}
