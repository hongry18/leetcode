package main

func canConstruct(ransomNote string, magazine string) bool {
	if ransomNote == magazine {
		return true
	}

	var cnt int
	ar := make([]bool, len(ransomNote))

	for _, mc := range magazine {
		for j, rc := range ransomNote {
			if mc == rc && !ar[j] {
				ar[j] = true
				cnt++
				break
			}
		}
	}

	return len(ar) == cnt
}
