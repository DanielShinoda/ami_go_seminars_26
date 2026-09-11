// Go-компилятор строг там, где C++ и Python лишь пожимают плечами.
//
// Запуск:
//
//	go run ./seminar-01-intro/examples/03-strict-compiler
//
// Живая демонстрация: по очереди раскомментируйте блоки ниже и покажите,
// что программа ПЕРЕСТАЁТ КОМПИЛИРОВАТЬСЯ. Это не предупреждение —
// это ошибка сборки. Мёртвый код физически не может попасть в репозиторий.
package main

import (
	"fmt"
	// "os" // (1) неиспользуемый импорт → "os" imported and not used
)

func main() {
	used := 42
	fmt.Println(used)

	// (2) неиспользуемая переменная → declared and not used: unused
	// unused := 1

	// (3) нет неявных приведений типов: int и float64 — разные типы
	// var i int = 1
	// var f float64 = 2.5
	// fmt.Println(i + f) // invalid operation: mismatched types int and float64

	// (4) нет неявного приведения к bool: if 1 { } не скомпилируется
	// if used {
	// 	fmt.Println("!")
	// } // non-boolean condition in if statement

	// (5) пропущенная точка с запятой ставится компилятором автоматически,
	// поэтому открывающая скобка ОБЯЗАНА стоять на той же строке:
	// func main()
	// {            // syntax error: unexpected semicolon or newline before {
}
