// Строки, байты, руны. Здесь ломается интуиция, принесённая из Python.
//
// Запуск:
//
//	go run ./seminar-02-basics/examples/02-strings-runes
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Привет, Go!"

	// ЛОВУШКА №1: len — это длина в БАЙТАХ, а не в символах.
	// Кириллица в UTF-8 занимает по 2 байта.
	fmt.Printf("%q: len=%d байт\n", s, len(s))
	fmt.Printf("рун (символов): %d\n", len([]rune(s)))

	// ЛОВУШКА №2: индексация даёт БАЙТ (тип byte = uint8), а не символ.
	fmt.Printf("s[0] = %d (%T), выводится как %q\n", s[0], s[0], s[0])

	// Правильный обход по символам — range: он декодирует UTF-8.
	// i — индекс НАЧАЛА руны в байтах, поэтому он «прыгает» через 2.
	fmt.Println("range по строке:")
	for i, r := range "Го!" {
		fmt.Printf("  байтовый индекс=%d руна=%q код=%d\n", i, r, r)
	}

	// rune — это псевдоним int32, код символа Unicode.
	var r rune = 'П'
	fmt.Printf("руна %q — это число %d\n", r, r)

	// ЛОВУШКА №3: строки НЕИЗМЕНЯЕМЫ. s[0] = 'X' не компилируется.
	// Чтобы менять — переводим в []rune, правим, собираем обратно.
	runes := []rune(s)
	runes[0] = 'п'
	fmt.Println(string(runes))

	// Склейка в цикле через += создаёт новую строку на каждой итерации:
	// O(n^2) по памяти. Для сборки строк есть strings.Builder.
	var b strings.Builder
	for i := range 5 {
		fmt.Fprintf(&b, "%d,", i)
	}
	fmt.Println(b.String())

	// Полезное из пакета strings — читать документацию, а не изобретать.
	fmt.Println(strings.ToUpper("go"), strings.Contains(s, "Go"), strings.Split("a,b,c", ","))
	fmt.Println(strings.Repeat("=", 10))
}
