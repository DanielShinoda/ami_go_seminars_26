package reverse

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "пустая строка", in: "", want: ""},
		{name: "один символ", in: "a", want: "a"},
		{name: "латиница", in: "hello", want: "olleh"},
		{name: "кириллица", in: "Привет", want: "тевирП"},
		{name: "двойной разворот даёт исходную", in: "Go — это просто", want: "отсорп отэ — oG"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Reverse(c.in); got != c.want {
				t.Errorf("Reverse(%q) = %q, ожидалось %q", c.in, got, c.want)
			}
		})
	}
}

// Свойство: развернув строку дважды, получаем исходную.
// Такой тест ловит ошибки, о которых вы не подумали, составляя таблицу.
func TestReverseTwice(t *testing.T) {
	for _, s := range []string{"", "a", "Привет, Go!", "🚀🎯", "ёжик"} {
		if got := Reverse(Reverse(s)); got != s {
			t.Errorf("Reverse(Reverse(%q)) = %q, ожидалось %q", s, got, s)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{name: "пустая строка", in: "", want: true},
		{name: "один символ", in: "x", want: true},
		{name: "палиндром латиницей", in: "level", want: true},
		{name: "палиндром кириллицей", in: "шалаш", want: true},
		{name: "не палиндром", in: "Привет", want: false},
		{name: "регистр важен", in: "Анна", want: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := IsPalindrome(c.in); got != c.want {
				t.Errorf("IsPalindrome(%q) = %v, ожидалось %v", c.in, got, c.want)
			}
		})
	}
}
