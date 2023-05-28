package main

import "testing"

func Test_kLengthApart(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args: args{nums: []int{1, 0, 0, 0, 1, 0, 0, 1}, k: 2}, want: true},
		{args: args{nums: []int{1, 0, 0, 1, 0, 1}, k: 2}, want: false},
		{args: args{nums: []int{0, 0, 1}, k: 3}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kLengthApart(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kLengthApart() = %v, want %v", got, tt.want)
			}
		})
	}
}
