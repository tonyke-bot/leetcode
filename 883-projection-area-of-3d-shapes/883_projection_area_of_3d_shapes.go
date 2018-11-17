package main

/*
Source: https://leetcode.com/problems/projection-area-of-3d-shapes
Test Case:
[[2]]
*/

func projectionArea(grid [][]int) int {
	size := len(grid)
	area := 0

	for i := 0; i < size; i++ {
		maxY := 0
		maxX := 0
		for j := 0; j < size; j++ {
			if grid[i][j] > 0 {
				area++
			}

			if grid[i][j] > maxY {
				maxY = grid[i][j]
			}

			if grid[j][i] > maxX {
				maxX = grid[j][i]
			}
		}
		area += maxY + maxX
	}

	return area
}
