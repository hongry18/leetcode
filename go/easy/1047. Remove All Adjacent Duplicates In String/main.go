package main

func main() {
	removeDuplicates("aababaab")
}

func removeDuplicates(s string) string {
	var r []byte

	for i := 0; i < len(s); i++ {
		if len(r) > 0 && s[i] == r[len(r)-1] {
			r = r[:len(r)-1]
		} else {
			r = append(r, s[i])
		}
	}
	return string(r)
}
