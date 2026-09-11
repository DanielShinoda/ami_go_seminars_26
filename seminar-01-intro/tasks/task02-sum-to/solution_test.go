package sumto

import "testing"

func TestSumTo(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want int
	}{
		{name: "n=1", in: 1, want: 1},
		{name: "n=5", in: 5, want: 15},
		{name: "n=100", in: 100, want: 5050},
		{name: "ноль", in: 0, want: 0},
		{name: "отрицательное", in: -7, want: 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := SumTo(c.in); got != c.want {
				t.Errorf("SumTo(%d) = %d, ожидалось %d", c.in, got, c.want)
			}
		})
	}
}
