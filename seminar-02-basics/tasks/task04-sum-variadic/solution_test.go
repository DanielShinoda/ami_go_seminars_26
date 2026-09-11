package variadic

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d, ожидалось 0", got)
	}
	if got := Sum(1, 2, 3); got != 6 {
		t.Errorf("Sum(1, 2, 3) = %d, ожидалось 6", got)
	}
	if got := Sum(-5, 5); got != 0 {
		t.Errorf("Sum(-5, 5) = %d, ожидалось 0", got)
	}

	// Проверяем «раскрытие» слайса: Sum(nums...)
	nums := []int{10, 20, 30}
	if got := Sum(nums...); got != 60 {
		t.Errorf("Sum(nums...) = %d, ожидалось 60", got)
	}
}

func TestMaxOf(t *testing.T) {
	cases := []struct {
		name  string
		first int
		rest  []int
		want  int
	}{
		{name: "один аргумент", first: 5, rest: nil, want: 5},
		{name: "максимум в конце", first: 1, rest: []int{2, 3}, want: 3},
		{name: "максимум в начале", first: 9, rest: []int{2, 3}, want: 9},
		{name: "отрицательные", first: -7, rest: []int{-3, -10}, want: -3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := MaxOf(c.first, c.rest...); got != c.want {
				t.Errorf("MaxOf(%d, %v) = %d, ожидалось %d", c.first, c.rest, got, c.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	if _, ok := Average(); ok {
		t.Error("Average() вернул ok = true, ожидалось false")
	}

	got, ok := Average(1, 2, 3, 4)
	if !ok {
		t.Fatal("Average(1,2,3,4) вернул ok = false, ожидалось true")
	}
	if math.Abs(got-2.5) > 1e-9 {
		t.Errorf("Average(1,2,3,4) = %v, ожидалось 2.5", got)
	}
}
