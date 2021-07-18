package main

import "testing"

func TestConvert(t *testing.T) {
	from := "ksh"
	to := "ghs"

	res := Convert(from, to, 100)

	got := res.To.Value
	want := 5.50

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
