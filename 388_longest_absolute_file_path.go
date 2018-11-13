package leetcode

import (
	"strings"
)

/*
Source: https://leetcode.com/problems/longest-absolute-file-path
Test Case:
"dir
\tsubdir1
\tsubdir2
\t\tfile.ext"
*/

func lengthLongestPath(input string) int {
	dirs := make([]int, 256)
	lines := strings.Split(input, "\n")
	longest := 0

	for _, line := range lines {
		depth := 0
		for line[depth] == '\t' {
			depth++
		}

		dirs[depth] = len(line) - depth

		dotPos := strings.LastIndex(line, ".")
		isFile := dotPos != -1 && dotPos != len(line)-1

		if isFile {
			currentLength := depth
			for _, length := range dirs[:depth+1] {
				currentLength += length
			}

			if currentLength > longest {
				longest = currentLength
			}
		}
	}

	return longest
}
