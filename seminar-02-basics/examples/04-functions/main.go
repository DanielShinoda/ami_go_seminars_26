// Функции: множественный возврат, именованные результаты,
// переменное число аргументов, функции как значения, замыкания.
//
// Запуск:
//
//	go run ./seminar-02-basics/examples/04-functions
package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// Множественный возврат — основная идиома Go.
	q, r := divmod(17, 5)
	fmt.Println(q, r)

	// Второе значение по соглашению — ошибка. Ошибка в Go это ЗНАЧЕНИЕ,
	// а не исключение: её нельзя «не заметить», её обрабатывают сразу.
	if v, err := parsePositive("42"); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("разобрали:", v)
	}
	if _, err := parsePositive("-1"); err != nil {
		fmt.Println("ошибка:", err) // сюда и попадём
	}

	fmt.Println(sum(), sum(1, 2, 3))

	// Слайс можно «раскрыть» в переменное число аргументов через ...
	nums := []int{4, 5, 6}
	fmt.Println(sum(nums...))

	// Функция — обычное значение: её можно положить в переменную,
	// передать аргументом и вернуть из другой функции.
	var op func(int, int) int = add
	fmt.Println(op(2, 3))
	fmt.Println(applyAll([]int{1, 2, 3}, double))

	// Анонимная функция прямо на месте.
	fmt.Println(applyAll([]int{1, 2, 3}, func(x int) int { return x * x }))

	// Замыкание: функция помнит переменные из окружения.
	next := counter()
	fmt.Println(next(), next(), next()) // 1 2 3

	// Два счётчика независимы — у каждого своя захваченная переменная.
	other := counter()
	fmt.Println(other())
}

// Параметры одного типа перечисляются через запятую: (a, b int).
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func add(a, b int) int { return a + b }

func double(x int) int { return x * 2 }

// Именованные результаты. Полезны как документация, когда неочевидно,
// что есть что. Голый return возвращает их текущие значения —
// но злоупотреблять им не стоит: читаемость падает.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// parsePositive показывает связку «значение + ошибка».
// errors.New создаёт ошибку, fmt.Errorf — форматирует.
func parsePositive(s string) (int, error) {
	if s == "" {
		return 0, errors.New("пустая строка")
	}
	if strings.HasPrefix(s, "-") {
		return 0, fmt.Errorf("%q: ожидалось положительное число", s)
	}
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0, fmt.Errorf("%q: не цифра в позиции символа %q", s, ch)
		}
		n = n*10 + int(ch-'0')
	}
	return n, nil
}

// Переменное число аргументов: внутри функции nums — обычный слайс []int.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Функция принимает функцию — обычное дело.
func applyAll(nums []int, f func(int) int) []int {
	out := make([]int, 0, len(nums))
	for _, n := range nums {
		out = append(out, f(n))
	}
	return out
}

// counter возвращает функцию, захватившую переменную n.
// n живёт столько, сколько живёт возвращённая функция —
// компилятор сам решает, разместить её на стеке или в куче.
func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

var _ = split // чтобы пример компилировался: split нигде не вызывается
