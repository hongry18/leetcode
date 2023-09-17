package main

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	div := numRows*2 - 2
	arr := make([][]byte, numRows)

	for i := range s {
		pos := i % div
		if pos == 0 || pos == numRows-1 || pos < numRows {
			arr[pos] = append(arr[pos], s[i])
			continue
		}
		pos2 := pos - (((i%div)%numRows)+1)*2
		arr[pos2] = append(arr[pos2], s[i])
	}

	var buf bytes.Buffer
	for i := range arr {
		buf.WriteString(string(arr[i]))
	}

	return buf.String()
}

func bestPractiecs(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	arr := make([][]byte, numRows)
	pos := 0
	move := -1
	n := len(s)

	for i := 0; i < n; i++ {
		if pos == 0 || pos == numRows-1 {
			move *= -1
		}

		arr[pos] = append(arr[pos], s[i])
		pos += move
	}

	return string(bytes.Join(arr, nil))
}
