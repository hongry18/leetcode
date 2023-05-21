package main

import (
	"bytes"
	"testing"
)

func TestXxx(t *testing.T) {
	t.Log(sortString("aaaabbbbcccc"))
}

func sortString(s string) string {
	var ar = [26]int{}
	for _, v := range s {
		ar[v-'a']++
	}

	var buf bytes.Buffer
	var i, direction = 0, 1
	for buf.Len() < len(s) {
		if i == 26 {
			i = 25
			direction = -1
		}
		if i == -1 {
			i = 0
			direction = 1
		}
		if ar[i] > 0 {
			buf.WriteByte(byte(i) + 'a')
			ar[i]--
		}

		i += direction
	}
	return buf.String()
}
