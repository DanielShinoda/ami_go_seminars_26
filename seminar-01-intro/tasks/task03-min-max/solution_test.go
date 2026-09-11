package minmax

import "testing"

func TestMinMax(t *testing.T) {
	cases := []struct {
		name    string
		in      []int
		wantMin int
		wantMax int
		wantOK  bool
	}{
		{name: "обычный слайс", in: []int{3, 1, 4, 1, 5}, wantMin: 1, wantMax: 5, wantOK: true},
		{name: "один элемент", in: []int{42}, wantMin: 42, wantMax: 42, wantOK: true},
		{name: "отрицательные", in: []int{-3, -10, -1}, wantMin: -10, wantMax: -1, wantOK: true},
		{name: "все одинаковые", in: []int{7, 7, 7}, wantMin: 7, wantMax: 7, wantOK: true},
		{name: "пустой слайс", in: []int{}, wantMin: 0, wantMax: 0, wantOK: false},
		{name: "nil-слайс", in: nil, wantMin: 0, wantMax: 0, wantOK: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotMin, gotMax, gotOK := MinMax(c.in)
			if gotMin != c.wantMin || gotMax != c.wantMax || gotOK != c.wantOK {
				t.Errorf("MinMax(%v) = (%d, %d, %v), ожидалось (%d, %d, %v)",
					c.in, gotMin, gotMax, gotOK, c.wantMin, c.wantMax, c.wantOK)
			}
		})
	}
}

// Тест проверяет, что функция не портит входной слайс:
// слайс передаётся «по ссылке на массив», и его легко случайно изменить.
func TestMinMaxDoesNotModifyInput(t *testing.T) {
	in := []int{5, 2, 9}
	MinMax(in)
	want := []int{5, 2, 9}
	for i := range in {
		if in[i] != want[i] {
			t.Fatalf("входной слайс изменён: %v, ожидалось %v", in, want)
		}
	}
}
