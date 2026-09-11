package countrunes

import "testing"

func TestCountRunes(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want int
	}{
		{name: "пустая строка", in: "", want: 0},
		{name: "латиница", in: "hello", want: 5},
		{name: "кириллица", in: "Привет", want: 6},
		{name: "смешанная строка", in: "Привет, Go!", want: 11},
		{name: "эмодзи — одна руна", in: "go🚀", want: 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := CountRunes(c.in); got != c.want {
				t.Errorf("CountRunes(%q) = %d, ожидалось %d", c.in, got, c.want)
			}
		})
	}
}

func TestCountOccurrences(t *testing.T) {
	cases := []struct {
		name   string
		in     string
		target rune
		want   int
	}{
		{name: "латинская буква", in: "banana", target: 'a', want: 3},
		{name: "кириллица", in: "молоко", target: 'о', want: 3},
		{name: "регистр важен", in: "Анна", target: 'а', want: 1},
		{name: "нет вхождений", in: "abc", target: 'z', want: 0},
		{name: "пустая строка", in: "", target: 'a', want: 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := CountOccurrences(c.in, c.target); got != c.want {
				t.Errorf("CountOccurrences(%q, %q) = %d, ожидалось %d", c.in, c.target, got, c.want)
			}
		})
	}
}
