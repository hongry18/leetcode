package main

func isValid(s string) bool {

	ar := []byte{}

	for i := range s {
		switch s[i] {
		case '(':
			ar = append(ar, ')')
		case '[':
			ar = append(ar, ']')
		case '{':
			ar = append(ar, '}')
		default:
			size := len(ar)
			if size == 0 && ar[size-1] != s[i] {
				return false
			}

			ar = ar[:size-1]
		}
	}

	return len(ar) == 0
}
