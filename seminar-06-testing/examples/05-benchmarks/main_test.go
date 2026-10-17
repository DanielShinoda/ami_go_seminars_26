// Бенчмарки: измеряем, а не гадаем.
//
//	go test -bench=. -benchmem ./seminar-06-testing/examples/05-benchmarks
package main

import (
	"strconv"
	"strings"
	"testing"
)

// Бенчмарк — функция BenchmarkXxx(b *testing.B). Раннер сам подбирает
// b.N так, чтобы замер занял ~1 секунду: прогоняет, удваивает N,
// прогоняет снова. Внутри цикла — только измеряемая работа, вся
// подготовка — до b.ResetTimer().

// JoinInts — склейка чисел через strings.Builder.
func JoinInts(n int) string {
	var b strings.Builder
	for i := range n {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(strconv.Itoa(i))
	}
	return b.String()
}

// JoinIntsNaive — та же склейка через конкатенацию строк.
func JoinIntsNaive(n int) string {
	s := ""
	for i := range n {
		if i > 0 {
			s += ","
		}
		s += strconv.Itoa(i)
	}
	return s
}

func BenchmarkJoinInts(b *testing.B) {
	for b.Loop() { // Go 1.24+: цикл без b.N и без ResetTimer
		JoinInts(100)
	}
}

func BenchmarkJoinIntsNaive(b *testing.B) {
	for b.Loop() {
		JoinIntsNaive(100)
	}
}

// Подбенчмарки через b.Run — таблица размеров, как в табличном тесте:
//
//	go test -bench='Join/builder/1000' ...
func BenchmarkJoin(b *testing.B) {
	for _, n := range []int{10, 100, 1000} {
		b.Run("builder/"+strconv.Itoa(n), func(b *testing.B) {
			for b.Loop() {
				JoinInts(n)
			}
		})
		b.Run("concat/"+strconv.Itoa(n), func(b *testing.B) {
			for b.Loop() {
				JoinIntsNaive(n)
			}
		})
	}
}

// Читаем вывод:
//
//	BenchmarkJoinInts-10   500000   2400 ns/op   1792 B/op   3 allocs/op
//	                       итераций  нс/операцию  байт/оп      аллокаций/оп
//
// -10 — значение GOMAXPROCS, не количество прогонов.
// -benchmem добавляет колонки про память — аллокации часто важнее
// наносекунд: меньше аллокаций — меньше работы сборщику мусора.
//
// Правила честного бенчмарка:
//   - результат из цикла должен КУДА-ТО деваться (присвойте в пакетную
//     переменную sink), иначе компилятор выбросит вызов целиком;
//   - сравнивайте версии на ОДНОЙ машине и в одном прогоне;
//   - числа в ±5% — это шум, а не оптимизация.
//
//nolint:unused // sink нужен только бенчмаркам ниже
var sink string

func BenchmarkJoinIntsWithSink(b *testing.B) {
	var s string
	for b.Loop() {
		s = JoinInts(100)
	}
	sink = s
}
