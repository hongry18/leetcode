package main

func detectCapitalUse(word string) bool {
	for _, r := range word {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}
