package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		// 이부분은 잘 이해 안감
		mp := map[byte]int{}
		for i := range s {
			if _, ok := mp[s[i]]; ok {
				return true
			}
			mp[s[i]]++
		}
		return false
	}

	diff := []int{}
	for i := range s {
		if s[i] != goal[i] {
			diff = append(diff, i)
		}
	}
	if len(diff) != 2 {
		return false
	}
	return s[diff[0]] == goal[diff[1]] && goal[diff[0]] == s[diff[1]]
}
