package unitest_test

import (
	"testing"
)

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	got := intMin(2, -2)
	if got != -2 {
		t.Errorf("got %d but want -2 for input IntMin(2, -2)", got)
	}
}

type testCase struct {
	name string
	in1  int
	in2  int
	want int
}

func TestIntMinMiddle(t *testing.T) {
	tc := testCase{
		name: "success",
		in1:  3,
		in2:  -1,
		want: -1,
	}

	makeAssertion(t, tc)
}

func makeAssertion(t *testing.T, tc testCase) {
	t.Run(tc.name, func(t *testing.T) {
		got := intMin(tc.in1, tc.in2)
		if got != tc.want {
			t.Errorf("got %d but want %d for input (%d, %d)", got, tc.want, tc.in1, tc.in2)
		}
	})
}

func TestIntMinTableDrivenMaxRefactoring(t *testing.T) {
	tt := []struct {
		name     string
		in1, in2 int
		want     int
	}{
		{name: "success", in1: 111, in2: 110, want: 110},
		{name: "negative-success", in1: -2398, in2: -9348, want: -9348},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := intMin(tc.in1, tc.in2)
			if got != tc.want {
				t.Errorf("got %d but want %d for input (%d, %d)", got, tc.want, tc.in1, tc.in2)
			}
		})
	}
}
