package main

import "testing"

func TestXxx(t *testing.T) {
	s := "123456"
	t.Log(s[0 : len(s)/2])
	t.Log(s[len(s)/2 : len(s)])
}

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{s: "abab"}, want: true},
		{name: "test2", args: args{s: "aba"}, want: false},
		{name: "test3", args: args{s: "abcabcabc"}, want: true},
		{name: "test3", args: args{s: "abcdbcabc"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
