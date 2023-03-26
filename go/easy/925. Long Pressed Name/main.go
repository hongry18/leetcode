package main

func main() {
	// isLongPressedName("alex", "aaleex")
	isLongPressedName("saeed", "ssaaedd")
	// isLongPressedName("alex", "aaleexa")
	// isLongPressedName("saeed", "ssaaaaaaaaeeeeeeeeedd")
}

func isLongPressedName(name string, typed string) bool {
	var i, j int
	var match bool

	for i < len(name) && j < len(typed) {
		if !match && name[i] == typed[j] {
			match = true
			continue
		} else if !match && name[i] != typed[j] {
			return false
		}

		if i < len(name)-1 && name[i] == name[i+1] && name[i] == typed[j] {
			i++
			j++
			match = false
			continue
		}

		if name[i] == typed[j] {
			j++
		} else {
			i++
			match = false
		}
	}

	return i == len(name)-1 && j == len(typed)
}
