package main

import (
	"bytes"
	"testing"
)

func TestXxx(t *testing.T) {
	// t.Log(reformat("a0b1c22"))
	t.Log(reformat("uzmlfxspduzb2621199741214"))
}

func reformat(s string) string {
	var ar1, ar2 []byte

	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			ar1 = append(ar1, s[i])
		} else {
			ar2 = append(ar2, s[i])
		}
	}

	if len(ar2) > len(ar1) {
		ar1, ar2 = ar2, ar1
	}

	if len(ar1)-len(ar2) > 1 {
		return ""
	}

	var b bytes.Buffer
	var i, j int
	for ; i < len(ar1) && j < len(ar2); i, j = i+1, j+1 {
		b.WriteByte(ar1[i])
		b.WriteByte(ar2[j])
	}

	if i < len(ar1) {
		b.WriteByte(ar1[i])
	}
	return b.String()
}
