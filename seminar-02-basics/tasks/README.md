# Задачи семинара 2

Правила решения и сдачи — [`docs/how-to-solve.md`](../../docs/how-to-solve.md).

```bash
go test ./seminar-02-basics/tasks/...
go test -v ./seminar-02-basics/tasks/task02-reverse
```

| # | Задача | Что тренируем |
|---|---|---|
| 1 | [`task01-count-runes`](task01-count-runes) | `range` по строке, руны против байтов |
| 2 | [`task02-reverse`](task02-reverse) | `[]rune`, обмен значений, два индекса |
| 3 | [`task03-grade`](task03-grade) | Обе формы `switch` |
| 4 | [`task04-sum-variadic`](task04-sum-variadic) | Переменное число аргументов |
| 5 | [`task05-counter`](task05-counter) | Замыкания |
| 6 | [`task06-collatz`](task06-collatz) | Цикл с условием, граничные случаи |

## 1. `CountRunes`, `CountOccurrences`

```go
func CountRunes(s string) int
func CountOccurrences(s string, target rune) int
```

Количество символов в строке и число вхождений заданной руны.

**Не используйте `len([]rune(s))`** — решайте циклом `range`, чтобы своими
глазами увидеть, как строка разбирается на руны. Проверьте себя на `"Привет"`:
`len` вернёт 12, а ответ — 6.

## 2. `Reverse`, `IsPalindrome`

```go
func Reverse(s string) string
func IsPalindrome(s string) bool
```

Разворот строки и проверка на палиндром. **По рунам, а не по байтам** —
наивный разворот байтов превратит кириллицу в мусор.

Отдельный тест проверяет свойство: `Reverse(Reverse(s)) == s` для любой строки.
Такие тесты ловят случаи, о которых вы не подумали, составляя таблицу вручную.

## 3. `Grade`, `Season`

```go
func Grade(score int) string
func Season(month int) string
```

Оценка по баллу и время года по номеру месяца.

Требования к решению (в этом весь смысл задачи):
- `Grade` — через `switch` **без выражения**: `switch { case ... }`;
- `Season` — через `switch` с перечислением значений: `case 12, 1, 2:`.

## 4. `Sum`, `MaxOf`, `Average`

```go
func Sum(nums ...int) int
func MaxOf(first int, rest ...int) int
func Average(nums ...int) (float64, bool)
```

Обратите внимание на сигнатуру `MaxOf`: обязательный первый аргумент делает
вызов без данных **невозможным на этапе компиляции**, поэтому и признак ошибки
возвращать не нужно. Это частый приём в реальном Go-коде.

В `Average` не забудьте про явное приведение: `int / int` — целочисленное
деление.

## 5. `New`, `Accumulator`

```go
func New() func() int
func Accumulator(start int) func(int) int
```

Счётчик и накопитель на замыканиях.

Ключевая проверка — тест `TestNewIndependent`: два счётчика, полученные
разными вызовами `New()`, обязаны быть независимы. Если вы использовали
переменную уровня пакета вместо локальной — тест это поймает.

## 6. `Steps`, `MaxValue`

```go
func Steps(n int) int
func MaxValue(n int) int
```

Последовательность Коллатца: чётное делим на 2, нечётное умножаем на 3 и
прибавляем 1. Считаем число шагов до единицы и максимальное достигнутое
значение. Для `n <= 0` — `-1`.

Проверьте на `n = 27`: 111 шагов и максимум 9232 — последовательность
забирается неожиданно высоко, хотя стартует с маленького числа.
