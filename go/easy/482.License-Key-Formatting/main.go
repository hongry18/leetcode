package main

func licenseKeyFormatting(s string, k int) string {
	var cnt int
	var res []byte
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}

		if cnt == k {
			res = append(res, '-')
			cnt = 0
		}

		if s[i] >= 'a' && s[i] <= 'z' {
			res = append(res, s[i]-'a'+'A')
		} else {
			res = append(res, s[i])
		}
		cnt++
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return string(res)
}
