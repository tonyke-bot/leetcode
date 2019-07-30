package problem0682

import "strconv"

/*
Source: https://leetcode.com/problems/baseball-game
Test Case:
["5","2","C","D","+"]
*/

func calPoints(ops []string) int {
	points := [1001]int{0}
	pos := 0
	sum := 0

	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "C":
			sum -= points[pos]
			pos--
			continue
		case "D":
			points[pos+1] = points[pos] << 1
		case "+":
			points[pos+1] = points[pos] + points[pos-1]
		default:
			points[pos+1], _ = strconv.Atoi(ops[i])
		}
		pos++
		sum += points[pos]
	}

	return sum
}
