package main

func canConstruct(ransomNote string, magazine string) bool {
	ar := [26]int{}

	for _, b := range magazine {
		ar[b-'a']++
	}

	for _, b := range ransomNote {
		if ar[b-'a'] == 0 {
			return false
		}
		ar[b-'a']--
	}

	return true
}
