package leetcode

// Source: https://leetcode.com/problems/unique-morse-code-words

var code = [][]byte{
	[]byte(".-"), []byte("-..."), []byte("-.-."), []byte("-.."),
	[]byte("."), []byte("..-."), []byte("--."), []byte("...."),
	[]byte(".."), []byte(".---"), []byte("-.-"), []byte(".-.."),
	[]byte("--"), []byte("-."), []byte("---"), []byte(".--."),
	[]byte("--.-"), []byte(".-."), []byte("..."), []byte("-"),
	[]byte("..-"), []byte("...-"), []byte(".--"), []byte("-..-"),
	[]byte("-.--"), []byte("--.."),
}

func uniqueMorseRepresentations(words []string) int {
	count := 0
	exists := make(map[string]bool, len(words))

	for _, word := range words {
		buffer := make([]byte, 0, len(word)*5)
		for _, c := range []byte(word) {
			buffer = append(buffer, code[c-'a']...)
		}

		cipher := string(buffer)
		if _, ok := exists[cipher]; !ok {
			count++
			exists[cipher] = true
		}
	}

	return count
}

/*
Test Case:
["gin", "zen", "gig", "msg"]
*/
