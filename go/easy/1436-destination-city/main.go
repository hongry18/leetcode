package main

func destCity(paths [][]string) string {
	out := map[string]int{}
	for _, path := range paths {
		out[path[0]]++
	}

	for _, path := range paths {
		if _, ok := out[path[1]]; !ok {
			return path[1]
		}
	}
	return ""
}
