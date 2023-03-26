package main

import (
	"bytes"
)

func main() {
	numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})
	numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"})
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]int)
	for _, email := range emails {
		var b bytes.Buffer
		var findPlus bool
		var cnt int
		for j, c := range email {
			if c == '@' {
				cnt = j
				break
			}

			if findPlus || c == '.' {
				continue
			}

			if c == '+' {
				findPlus = true
				continue
			}

			b.WriteRune(c)
		}
		b.WriteString(email[cnt:])
		m[b.String()]++
	}

	return len(m)
}
