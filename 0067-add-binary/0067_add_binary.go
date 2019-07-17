package problem0067

// Source: https://leetcode.com/problems/add-binary

func addBinary(a string, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}

	result := ""
	bytesA := []byte(a)
	bytesB := []byte(b)
	carry := byte(0)

	for i := 0; ; i++ {
		if len(bytesA) <= i && len(bytesB) <= i {
			if carry > 0 {
				result = string(carry+'0') + result
			}
			return result
		}

		valueA := byte(0)
		valueB := byte(0)

		if len(bytesA) > i {
			valueA = bytesA[len(bytesA)-1-i] - '0'
		}
		if len(bytesB) > i {
			valueB = bytesB[len(bytesB)-1-i] - '0'
		}

		sum := valueA + valueB + carry
		result = string(sum%2+'0') + result
		carry = sum / 2
	}
}

/*
Test Case:
"11"
"1"
*/
