package counter

import "testing"

func TestNew(t *testing.T) {
	next := New()
	for i := 1; i <= 5; i++ {
		if got := next(); got != i {
			t.Fatalf("вызов №%d вернул %d, ожидалось %d", i, got, i)
		}
	}
}

// Ключевая проверка: два счётчика не делят состояние.
// Если внутри New использована глобальная переменная — тест упадёт.
func TestNewIndependent(t *testing.T) {
	a := New()
	b := New()

	a()
	a()
	a() // a досчитал до 3

	if got := b(); got != 1 {
		t.Errorf("второй счётчик вернул %d, ожидалось 1 — счётчики делят состояние", got)
	}
	if got := a(); got != 4 {
		t.Errorf("первый счётчик вернул %d, ожидалось 4", got)
	}
}

func TestAccumulator(t *testing.T) {
	acc := Accumulator(10)

	steps := []struct {
		add  int
		want int
	}{
		{add: 5, want: 15},
		{add: -3, want: 12},
		{add: 0, want: 12},
		{add: 100, want: 112},
	}

	for _, s := range steps {
		if got := acc(s.add); got != s.want {
			t.Errorf("acc(%d) = %d, ожидалось %d", s.add, got, s.want)
		}
	}

	if got := Accumulator(0)(7); got != 7 {
		t.Errorf("новый аккумулятор вернул %d, ожидалось 7", got)
	}
}
