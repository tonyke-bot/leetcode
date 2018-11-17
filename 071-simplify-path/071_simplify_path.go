package main

import (
	"strings"
)

/*
Source: https://leetcode.com/problems/simplify-path
Test Case:
"/home/"
*/

func simplifyPath(path string) string {
	stack := make([]string, 0, 128)

	for _, dir := range strings.Split(path, "/") {
		switch dir {
		case "", ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, dir)
		}
	}

	return "/" + strings.Join(stack, "/")
}
