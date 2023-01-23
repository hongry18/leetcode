package main

import "testing"

func Test_findContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test1",
			args: args{
				g: []int{1, 2, 3},
				s: []int{1, 3, 2},
			},
			want: 3,
		},
		{name: "test2",
			args: args{
				g: []int{1, 2, 3},
				s: []int{1, 1},
			},
			want: 1,
		},
		{name: "test3",
			args: args{
				g: []int{1, 2},
				s: []int{1, 2, 3},
			},
			want: 2,
		},
		{name: "test4",
			args: args{
				g: []int{10, 9, 8, 7, 10, 9, 8, 7},
				s: []int{10, 9, 8, 7},
			},
			want: 4,
		},
		{name: "test5",
			args: args{
				g: []int{10, 9, 8, 7},
				s: []int{5, 6, 7, 8},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(tt.args.g, tt.args.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
