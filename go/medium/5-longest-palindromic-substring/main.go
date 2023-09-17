package main

/*
0부터 s의 길이까지 반복을 한다.
idx를 중심으로 palindrom이 valid할때까지 반복한다
idx, idx+1을 중심으로 palindrome가 valid할때까지 반복한다.
*/
func bestPractices(s string) string {
	var result string
	var n int = len(s)
	var max int

	for i := 0; i < n; i++ {
		var left, right int

		left, right = i, i
		for ; left >= 0 && right < n && s[left] == s[right]; left, right = left-1, right+1 {
			if (right - left + 1) > max {
				max = right - left + 1
				result = s[left : right+1]
			}
		}

		left, right = i, i+1
		for ; left >= 0 && right < n && s[left] == s[right]; left, right = left-1, right+1 {
			if (right - left + 1) > max {
				max = right - left + 1
				result = s[left : right+1]
			}
		}
	}

	if result == "" {
		return s[0:1]
	}
	return result
}

/*
윈도우 크기를 최대로 한 다음 윈도우 크기를 1씩 줄여가면서 palindrome가 valid한지 확인한다. 이건 느리다.
*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return string(s[0])
	}

	last := len(s)
	size := last

	for size > 0 {
		for i := 0; i < last-size; i++ {
			if valid(s, i, i+size) {
				return s[i : i+size+1]
			}
		}
		size--
	}

	return string(s[0])
}

func valid(s string, i, j int) bool {
	for ; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
