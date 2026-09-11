package grade

import "testing"

func TestGrade(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{in: 100, want: "отлично"},
		{in: 90, want: "отлично"},
		{in: 89, want: "хорошо"},
		{in: 75, want: "хорошо"},
		{in: 74, want: "удовлетворительно"},
		{in: 60, want: "удовлетворительно"},
		{in: 59, want: "неудовлетворительно"},
		{in: 0, want: "неудовлетворительно"},
		{in: 101, want: "некорректный балл"},
		{in: -1, want: "некорректный балл"},
	}

	for _, c := range cases {
		if got := Grade(c.in); got != c.want {
			t.Errorf("Grade(%d) = %q, ожидалось %q", c.in, got, c.want)
		}
	}
}

func TestSeason(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{in: 12, want: "зима"},
		{in: 1, want: "зима"},
		{in: 2, want: "зима"},
		{in: 3, want: "весна"},
		{in: 5, want: "весна"},
		{in: 6, want: "лето"},
		{in: 8, want: "лето"},
		{in: 9, want: "осень"},
		{in: 11, want: "осень"},
		{in: 0, want: "некорректный месяц"},
		{in: 13, want: "некорректный месяц"},
	}

	for _, c := range cases {
		if got := Season(c.in); got != c.want {
			t.Errorf("Season(%d) = %q, ожидалось %q", c.in, got, c.want)
		}
	}
}
