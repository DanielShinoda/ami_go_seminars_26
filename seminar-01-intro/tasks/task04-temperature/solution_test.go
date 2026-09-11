package temperature

import (
	"math"
	"testing"
)

// Числа с плавающей точкой нельзя сравнивать через ==:
// результат арифметики почти никогда не совпадает бит в бит.
// Сравниваем с допуском — так делают в настоящих проектах.
const eps = 1e-9

func TestCToF(t *testing.T) {
	cases := []struct {
		name string
		in   float64
		want float64
	}{
		{name: "замерзание воды", in: 0, want: 32},
		{name: "кипение воды", in: 100, want: 212},
		{name: "точка совпадения шкал", in: -40, want: -40},
		{name: "комнатная температура", in: 21.5, want: 70.7},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := CToF(c.in)
			if math.Abs(got-c.want) > eps {
				t.Errorf("CToF(%v) = %v, ожидалось %v", c.in, got, c.want)
			}
		})
	}
}

func TestRoundToInt(t *testing.T) {
	cases := []struct {
		name string
		in   float64
		want int
	}{
		{name: "вниз", in: 21.4, want: 21},
		{name: "ровно половина вверх", in: 21.5, want: 22},
		{name: "целое", in: -3, want: -3},
		{name: "отрицательная половина", in: -21.5, want: -22},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := RoundToInt(c.in); got != c.want {
				t.Errorf("RoundToInt(%v) = %d, ожидалось %d", c.in, got, c.want)
			}
		})
	}
}
