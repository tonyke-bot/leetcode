package leetcode

/*
Source: https://leetcode.com/problems/lemonade-change
Test Case:
[5,5,5,10,20]
*/

func lemonadeChange(bills []int) bool {
	balance := map[int]int{5: 0, 10: 0, 20: 0}

	for _, bill := range bills {
		switch bill {
		case 5:
		case 10:
			if balance[5] == 0 {
				return false
			}
			balance[5]--
		case 20:
			if balance[10] > 0 && balance[5] > 0 {
				balance[5]--
				balance[10]--
			} else if balance[5] > 2 {
				balance[5] -= 3
			} else {
				return false
			}
		default:
			return false
		}
		balance[bill]++
	}

	return true
}
