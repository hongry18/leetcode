package main

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "babad"}, want: "bab"},
		{args: args{s: "aacabdkacaa"}, want: "aca"},
		{args: args{s: "cbbd"}, want: "bb"},
		{args: args{s: "aaa"}, want: "aaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valid(t *testing.T) {
	type args struct {
		s string
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "abcba", i: 0, j: 4}, want: true},
		{args: args{s: "abcbb", i: 0, j: 4}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid(tt.args.s, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bestPractices(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "babed"}, want: "bab"},
		{args: args{s: "aacabdkacaa"}, want: "aca"},
		{args: args{s: "cbbd"}, want: "bb"},
		{args: args{s: "aaa"}, want: "aaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestPractices(tt.args.s); got != tt.want {
				t.Errorf("bestPractices() = %v, want %v", got, tt.want)
			}
		})
	}
}
