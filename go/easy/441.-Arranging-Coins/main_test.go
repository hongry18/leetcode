package main

import "testing"

func Test_arrangeCoins(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{n: 5}, want: 2},
		{name: "test2", args: args{n: 8}, want: 3},
		{name: "test3", args: args{n: 1}, want: 1},
		{name: "test4", args: args{n: 6}, want: 3},
		{name: "test5", args: args{n: 10}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoins(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
