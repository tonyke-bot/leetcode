package problem0709

// Source: https://leetcode.com/problems/to-lower-case

var delta byte = 'a' - 'A'

func toLowerCase(str string) string {
	buffer := make([]byte, 0, len(str))
	for _, v := range []byte(str) {
		if 'A' <= v && v <= 'Z' {
			v += delta
		}

		buffer = append(buffer, v)
	}
	return string(buffer)
}

/*
Test Case:
"Hello"
*/
