package main

func detectCapitalUse(word string) bool {
	var nCap, nLow int
	for _, b := range word {
		if b >= 'a' && b <= 'z' {
			nLow++
		} else if b >= 'A' && b <= 'Z' {
			nCap++
		}
	}

	if nCap == 0 {
		return true
	}

	if len(word) == nCap {
		return true
	}

	if nCap == 1 && word[0] >= 'A' && word[0] <= 'Z' {
		return true
	}
	return false
}
