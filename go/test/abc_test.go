package test

import "testing"

func TestAbc(t *testing.T) {
	var a *int
	b := 6
	a = &b

	c := new(int)
	d := new(int)
	*c = *a
	*d = *a

	*c = 7
	*d = 2

	t.Log(*a)
	t.Log(*c)
	t.Log(*d)
}
