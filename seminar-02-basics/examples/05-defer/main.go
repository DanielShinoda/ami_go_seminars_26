// defer: отложенный вызов.
//
// Запуск:
//
//	go run ./seminar-02-basics/examples/05-defer
//
// Аналогия для тех, кто знает C++: это RAII, только явный и на уровне функции.
// Аналогия для тех, кто знает Python: это with / finally.
package main

import (
	"fmt"
	"os"
)

func main() {
	basics()
	fmt.Println("аргумент вычисляется сразу:")
	argEvaluation()
	fmt.Println("результат namedResult():", namedResult())
	writeFileDemo()
}

func basics() {
	fmt.Println("--- basics ---")
	// Отложенные вызовы складываются в стек и выполняются в порядке LIFO
	// при ВЫХОДЕ ИЗ ФУНКЦИИ (не из блока!) — в том числе при панике.
	defer fmt.Println("выполнится третьим")
	defer fmt.Println("выполнится вторым")
	fmt.Println("выполнится первым")
}

func argEvaluation() {
	i := 0
	// ЛОВУШКА: аргументы defer вычисляются В МОМЕНТ ОБЪЯВЛЕНИЯ,
	// а сам вызов происходит потом. Здесь напечатается 0, а не 10.
	defer fmt.Println("  defer видит i =", i)
	i = 10
	fmt.Println("  в конце функции i =", i)
}

// defer может изменить ИМЕНОВАННЫЙ результат функции —
// это единственный способ поправить возвращаемое значение перед выходом.
func namedResult() (result int) {
	defer func() {
		result *= 2
	}()
	return 21 // вернётся 42
}

func writeFileDemo() {
	fmt.Println("--- работа с файлом ---")
	f, err := os.CreateTemp("", "seminar-*.txt")
	if err != nil {
		fmt.Println("не удалось создать файл:", err)
		return
	}
	// Закрытие и удаление ставим сразу после успешного открытия.
	// Дальше по коду можно спокойно делать return на любой ветке.
	defer os.Remove(f.Name())
	defer f.Close()

	if _, err := f.WriteString("привет из defer-примера\n"); err != nil {
		fmt.Println("ошибка записи:", err)
		return
	}

	data, err := os.ReadFile(f.Name())
	if err != nil {
		fmt.Println("ошибка чтения:", err)
		return
	}
	fmt.Printf("  прочитано из %s: %q\n", f.Name(), string(data))
}
