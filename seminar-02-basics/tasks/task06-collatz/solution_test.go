package collatz

import "testing"

func TestSteps(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{in: 1, want: 0},
		{in: 2, want: 1},
		{in: 3, want: 7},
		{in: 6, want: 8},
		{in: 27, want: 111},
		{in: 0, want: -1},
		{in: -5, want: -1},
	}

	for _, c := range cases {
		if got := Steps(c.in); got != c.want {
			t.Errorf("Steps(%d) = %d, ожидалось %d", c.in, got, c.want)
		}
	}
}

func TestMaxValue(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{in: 1, want: 1},
		{in: 2, want: 2},
		{in: 3, want: 16},
		{in: 7, want: 52},
		{in: 27, want: 9232},
		{in: 0, want: -1},
	}

	for _, c := range cases {
		if got := MaxValue(c.in); got != c.want {
			t.Errorf("MaxValue(%d) = %d, ожидалось %d", c.in, got, c.want)
		}
	}
}
