// Файл с суффиксом _test.go компилируется ТОЛЬКО при `go test`.
// В обычную сборку он не попадает — тесты не увеличивают ваш бинарь.
//
// Запуск:
//
//	go test ./seminar-01-intro/examples/05-first-test
//	go test -v ./seminar-01-intro/examples/05-first-test   # с именами подтестов
package mathutil

import "testing"

// Тест — это функция TestXxx(t *testing.T) в файле _test.go.
// Никаких фреймворков, аннотаций и регистрации: так работает сам `go test`.
func TestAbs(t *testing.T) {
	got := Abs(-3)
	want := 3
	if got != want {
		// t.Errorf помечает тест проваленным и продолжает выполнение.
		// t.Fatalf — помечает и немедленно останавливает эту тестовую функцию.
		t.Errorf("Abs(-3) = %d, ожидалось %d", got, want)
	}
}

// Идиома Go — табличный тест: список случаев + один цикл.
// Именно так написаны тесты во всех ваших задачах.
func TestAbsTable(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want int
	}{
		{name: "отрицательное", in: -5, want: 5},
		{name: "ноль", in: 0, want: 0},
		{name: "положительное", in: 7, want: 7},
	}

	for _, c := range cases {
		// t.Run создаёт подтест: в выводе будет видно, какой случай упал.
		t.Run(c.name, func(t *testing.T) {
			if got := Abs(c.in); got != c.want {
				t.Errorf("Abs(%d) = %d, ожидалось %d", c.in, got, c.want)
			}
		})
	}
}

// Тест может проверять и неэкспортируемые функции:
// тестовый файл лежит в том же пакете и видит всё.
func TestClamp(t *testing.T) {
	if got := clamp(15, 0, 10); got != 10 {
		t.Errorf("clamp(15, 0, 10) = %d, ожидалось 10", got)
	}
}
