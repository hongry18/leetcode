package main

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	var isUpper bool
	if word[0] >= 'A' && word[0] <= 'Z' {
		isUpper = true
	}
	for i := 1; i < len(word); i++ {
		if !isUpper && word[i] >= 'A' && word[i] <= 'Z' {
			return false
		}

		if i == 1 && isUpper && word[i] >= 'a' && word[i] <= 'z' {
			isUpper = false
			continue
		}

		if isUpper && (word[i] < 'A' || word[i] > 'Z') {
			return false
		}

		if !isUpper && (word[i] < 'a' || word[i] > 'z') {
			return false
		}
	}
	return true
}
