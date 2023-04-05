package main

func longestCommonPrefix(strs []string) string {
	min := 200

	for _, s := range strs {
		if len(s) < min {
			min = len(s)
		}
	}

	if min == 0 {
		return ""
	}

	for i := 0; i < min; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:min]
}
