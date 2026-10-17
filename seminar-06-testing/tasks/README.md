# Задачи семинара 6

Этот семинар — наоборот: реализация уже написана, ваше дело — написать
для неё тесты. В каждой задаче рядом лежат:

- `<имя>.go` — эталонная реализация; её контракт (что гарантируется,
  какие граничные случаи) описан в doc-комментариях — это и есть условие;
- `<имя>_test.go` — файл, в который вы пишете тесты (условие — в шапке);
- `variants/` — сломанные реализации с типичными дефектами;
- `harness_test.go` — приёмка, он не для правки.

Задача принята, когда ваши тесты **проходят на эталонной реализации** и
**падают на КАЖДОЙ сломанной** из `variants/`. Прогон по всем вариантам
делает харнесс: достаточно `go test .` в папке задачи или
`make task T=seminar-06-testing/tasks/task01-assert`. Все задачи семинара
— `make test-06`. Подробно про цикл работы, команды и сдачу —
[`docs/how-to-solve.md`](../../docs/how-to-solve.md).

Хороший набор тестов — это не «много тестов», а такой, который отличает
эталон от каждой сломанной копии. Если харнесс говорит «слабые тесты» —
какой-то дефект из `variants/` ваши тесты не заметили.

Прогоняйте тесты с детектором гонок — конкурентные задачи без него могут
выглядеть решёнными, не будучи ими:

```bash
go test -race ./seminar-06-testing/tasks/...
```

| # | Задача | Что тестируем | О чём |
|---|---|---|---|
| 1 | [`task01-assert`](task01-assert) | `Equal`, `True`, `NoError` | Хелперы проверяют фейком интерфейса `TestingT`: порядок вызовов `Helper`/`Errorf`, содержимое сообщений |
| 2 | [`task02-eventually`](task02-eventually) | `Eventually` | Время в тестах: немедленный первый опрос, дедлайн, отсутствие busy-spin, ровно один опрос при `timeout <= 0` |
| 3 | [`task03-cache`](task03-cache) | `New`, `Get`, `Set`, `Len` | Стресс-тест и `-race`: потокобезопасность, перезапись ключа, независимость экземпляров |
| 4 | [`task04-retry`](task04-retry) | `Do` | Подмена зависимости: паузы через свой `Sleep` вместо настоящих ожиданий — тесты детерминированные и быстрые |
| 5 | [`task05-collect`](task05-collect) | `Collect` | Таймауты и закрытие каналов: частичный результат, таймер не сбрасывается на каждое значение |

Пятая задача — дополнительная.

## Сигнатуры

```go
// task01-assert
func Equal(t TestingT, got, want any)
func True(t TestingT, cond bool)
func NoError(t TestingT, err error)

// task02-eventually
func Eventually(cond func() bool, timeout time.Duration) bool

// task03-cache
func New() *Cache
func (c *Cache) Get(key string) (string, bool)
func (c *Cache) Set(key, value string)
func (c *Cache) Len() int

// task04-retry
func Do(op func() error, opts Options) error

// task05-collect
func Collect(ch <-chan int, timeout time.Duration) (vals []int, closed bool)
```
