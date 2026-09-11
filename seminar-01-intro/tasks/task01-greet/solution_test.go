package greet

import "testing"

func TestGreet(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "обычное имя", in: "Аня", want: "Привет, Аня!"},
		{name: "латиница", in: "Bob", want: "Привет, Bob!"},
		{name: "пустая строка", in: "", want: "Привет, мир!"},
		{name: "пробел — это не пустая строка", in: " ", want: "Привет,  !"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Greet(c.in); got != c.want {
				t.Errorf("Greet(%q) = %q, ожидалось %q", c.in, got, c.want)
			}
		})
	}
}
