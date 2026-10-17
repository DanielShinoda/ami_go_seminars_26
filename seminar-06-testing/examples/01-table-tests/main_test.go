// Табличные тесты глубже: subtests, t.Helper, Errorf против Fatalf.
//
//	go test -v ./seminar-06-testing/examples/01-table-tests
package main

import (
	"strings"
	"testing"
)

// SplitWords — функция, которую мы тестируем (в задачах она лежала бы
// в solution.go).
func SplitWords(s string) []string {
	return strings.Fields(s)
}

// Канонический табличный тест вы уже видели во всех задачах курса.
// Здесь разбираем детали, на которые до сих пор не хватало времени.
func TestSplitWords(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		want []string
	}{
		{name: "несколько слов", in: "go is fun", want: []string{"go", "is", "fun"}},
		{name: "лишние пробелы", in: "  a  b ", want: []string{"a", "b"}},
		{name: "пустая строка", in: "", want: nil},
	}

	for _, tc := range testCases {
		// t.Run создаёт подтест. Имя по-русски читается в выводе -v.
		// tc обязательно передаём через замыкание — до Go 1.22 это был
		// источник классического бага с общей переменной цикла.
		t.Run(tc.name, func(t *testing.T) {
			got := SplitWords(tc.in)
			if len(got) != len(tc.want) {
				// Fatalf прерывает ЭТОТ подтест: дальше сравнивать
				// поэлементно бессмысленно, индексы разъехались.
				t.Fatalf("SplitWords(%q) = %v, ожидалось %v", tc.in, got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					// Errorf помечает тест упавшим, но продолжает:
					// увидим все несовпадения, а не первое.
					t.Errorf("элемент %d = %q, ожидалось %q", i, got[i], tc.want[i])
				}
			}
		})
	}
}

// Правило выбора: Fatalf — когда дальнейшие проверки бессмысленны или
// упадут сами (nil-указатель, другая длина); Errorf — когда остальные
// проверки ещё дадут информацию.

// checkEqual — тестовый хелпер. t.Helper() помечает его служебным:
// в сообщении об ошибке будет строка ВЫЗОВА хелпера, а не строка внутри
// него. Без t.Helper() все ошибки «родом» из одной строки — искать
// упавший случай придётся по тексту.
func checkEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("получили %d, ожидалось %d", got, want)
	}
}

func TestHelperPointsAtCaller(t *testing.T) {
	checkEqual(t, 2+2, 4)
	checkEqual(t, len(SplitWords("a b c")), 3)
}

// Отдельный подтест можно запустить по имени:
//
//	go test -run 'TestSplitWords/лишние' ./seminar-06-testing/examples/01-table-tests
//
// -run принимает регексп — пробелы в имени заменяются на подчёркивания.
