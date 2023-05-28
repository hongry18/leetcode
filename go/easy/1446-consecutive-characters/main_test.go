package main

import "testing"

func Test_maxPower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{args: args{s: "leetcode"}, want: 2},
		{args: args{s: "abbcccddddeeeeedcba"}, want: 5},
		{args: args{s: "letcodec"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.args.s); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
