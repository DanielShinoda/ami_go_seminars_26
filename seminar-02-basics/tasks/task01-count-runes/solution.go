package countrunes

// CountRunes возвращает количество СИМВОЛОВ (рун) в строке.
// Для "Привет" это 6, хотя len("Привет") == 12.
//
// Требование: не использовать len([]rune(s)) — решить циклом range,
// чтобы увидеть, как строка разбирается на руны.
func CountRunes(s string) int {
	panic("TODO: реализуйте CountRunes")
}

// CountOccurrences возвращает, сколько раз руна target встречается в строке.
// Сравнение регистрозависимое: 'а' и 'А' — разные руны.
func CountOccurrences(s string, target rune) int {
	panic("TODO: реализуйте CountOccurrences")
}
