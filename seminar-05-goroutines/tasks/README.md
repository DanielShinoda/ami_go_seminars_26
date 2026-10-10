# Задачи семинара 5

Правите только `solution.go`. `solution_test.go` и `types.go` не трогаете:
тесты — это и есть условие, а типы зафиксированы, чтобы ваше решение
собиралось с чужим тестом. Прогнать задачи семинара — `make test-05`.
Подробно про цикл работы, команды и сдачу —
[`docs/how-to-solve.md`](../../docs/how-to-solve.md).

Условие каждой задачи целиком лежит в doc-комментариях над функциями в
`solution.go`. Читайте его внимательно: там оговорены граничные случаи, и
тесты проверяют каждый.

Прогоняйте тесты с детектором гонок — в задачах на конкурентность это
половина проверки:

```bash
go test -race ./seminar-05-goroutines/tasks/...
```

| # | Задача | Функции и методы | О чём |
|---|---|---|---|
| 1 | [`task01-parallel`](task01-parallel) | `Run` | Запуск горутин, `sync.WaitGroup`, сбор результатов в порядке входа |
| 2 | [`task02-chansum`](task02-chansum) | `Generator`, `Sum` | Каналы: отправка, чтение до закрытия, `close` |
| 3 | [`task03-pipeline`](task03-pipeline) | `Map`, `Filter` | Этапы конвейера: горутина + входной и выходной каналы |
| 4 | [`task04-pool`](task04-pool) | `Process` | Пул воркеров с ограничением параллелизма |
| 5 | [`task05-counter`](task05-counter) | `Inc`, `Add`, `Value` | Конкурентно безопасное состояние, `sync.Mutex` |

Пятая задача — дополнительная.

## Сигнатуры

```go
// task01-parallel
func Run(fns []func() int) []int

// task02-chansum
func Generator(nums ...int) <-chan int
func Sum(ch <-chan int) int

// task03-pipeline
func Map(in <-chan int, f func(int) int) <-chan int
func Filter(in <-chan int, keep func(int) bool) <-chan int

// task04-pool
func Process(jobs []int, f func(int) int, workers int) []int

// task05-counter
func (c *Counter) Inc()
func (c *Counter) Add(delta int)
func (c *Counter) Value() int
```
