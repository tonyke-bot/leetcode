package leetcode

/*
Source: https://leetcode.com/problems/find-and-replace-pattern
Test Case:
["abc","deq","mee","aqq","dkd","ccc"]
"abb"
*/

func findAndReplacePattern(words []string, pattern string) []string {
	result := make([]string, 0, len(words))
	patternBytes := []byte(pattern)

	for i := 0; i < len(words); i++ {
		bytes := []byte(words[i])
		if len(bytes) != len(patternBytes) {
			continue
		}

		match := true
		mapping := map[byte]byte{}
		usedByte := map[byte]bool{}
		for j := 0; j < len(bytes) && match; j++ {
			pByte := patternBytes[j]
			wByte := bytes[j]
			if mapping[pByte] == 0 && !usedByte[wByte] {
				mapping[pByte] = wByte
				usedByte[wByte] = true
			}

			match = mapping[pByte] == wByte
		}

		if match {
			result = append(result, words[i])
		}
	}

	return result
}
