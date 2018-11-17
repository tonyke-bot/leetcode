package leetcode

import (
	"strings"
)

// Source: https://leetcode.com/problems/unique-email-addresses

func parseEmail(email string) string {
	atPos := strings.LastIndex(email, "@")
	domain := email[atPos:]

	username := make([]byte, 0, len(email)-atPos-1)
iter:
	for _, v := range []byte(email[:atPos]) {
		switch v {
		case '+':
			break iter
		case '.':
			continue
		default:
			username = append(username, v)
		}
	}
	return string(username) + domain
}

func numUniqueEmails(emails []string) int {
	exists := map[string]bool{}
	count := 0

	for _, rawEmail := range emails {
		email := parseEmail(rawEmail)
		if _, ok := exists[email]; ok {
			continue
		}

		exists[email] = true
		count++
	}

	return count
}

/*
Test Case:
["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
*/
