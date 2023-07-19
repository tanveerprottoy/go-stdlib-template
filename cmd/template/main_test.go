package main

import "testing"

type TestCase struct {
	arg1 int
	arg2 int
	want int
}

// Test table example
func TestMultiply(t *testing.T) {
	cases := []TestCase{
		{2, 3, 6},
		{10, 5, 50},
		{-8, -3, 24},
		{0, 9, 0},
		{-7, 6, -42},
	}

	for _, tc := range cases {
		g := Multiply(tc.arg1, tc.arg2)
		if tc.want != g {
			t.Errorf("Expected '%d', but got '%d'", tc.want, g)
		}
	}
}
