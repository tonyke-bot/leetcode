package leetcode

// Source: https://leetcode.com/problems/implement-strstr

func strStr(haystack string, needle string) int {
	src := []byte(haystack)
	target := []byte(needle)

	if len(target) == 0 {
		return 0
	}

	for pos, char := range src {
		if len(src)-pos < len(target) {
			break
		}

		if char != target[0] {
			continue
		}

		same := true
		for i := 1; i < len(target) && same; i++ {
			same = target[i] == src[pos+i]
		}

		if same {
			return pos
		}
	}

	return -1
}

/*
Test Case:
"hello"
"ll"
*/
