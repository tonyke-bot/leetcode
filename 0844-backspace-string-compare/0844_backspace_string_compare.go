package problem0844

/*
Source: https://leetcode.com/problems/backspace-string-compare
Test Case:
"ab#c"
"ad#c"
*/

func backspaceCompare(S string, T string) bool {
	bytesS := []byte(S)
	lenS := len(bytesS)
	posS := lenS - 1
	offsetS := 0
	bytesT := []byte(T)
	lenT := len(bytesT)
	posT := lenT - 1
	offsetT := 0
	matched := true

	for (posT >= 0 || posS >= 0) && matched {
		for posT >= 0 && (bytesT[posT] == '#' || offsetT != 0) {
			if bytesT[posT] == '#' {
				offsetT--
			} else {
				offsetT++
			}

			posT--
		}

		for posS >= 0 && (bytesS[posS] == '#' || offsetS != 0) {
			if bytesS[posS] == '#' {
				offsetS--
			} else {
				offsetS++
			}

			posS--
		}

		if posT == -1 && posS == -1 {
			return true
		}

		if !(posT >= 0 && posS >= 0 && bytesS[posS] == bytesT[posT]) {
			return false
		}

		posT--
		posS--
	}

	return posT == posS
}
