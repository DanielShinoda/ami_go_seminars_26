// Моки через mockgen: тестируем код с внешними зависимостями.
//
// Этот пример живёт в тестах — запускайте его так:
//
//	go test -v ./seminar-06-testing/examples/06-mocks
package main

import "time"

// Sleeper — зависимость нашего кода: умение ждать.
// В проде это time.Sleep, в тесте — сгенерированный мок.
type Sleeper interface {
	Sleep(d time.Duration)
}

// realSleeper — продовая реализация.
type realSleeper struct{}

func (realSleeper) Sleep(d time.Duration) { time.Sleep(d) }

// Retry вызывает op до успеха, но не больше attempts раз, с растущей
// паузой между попытками: после i-й неудачной (нумерация с 1) — 10*i мс.
func Retry(attempts int, s Sleeper, op func() error) error {
	var err error
	for i := 1; i <= attempts; i++ {
		if err = op(); err == nil {
			return nil
		}
		if i < attempts {
			s.Sleep(time.Duration(i) * 10 * time.Millisecond)
		}
	}
	return err
}

// Мок для Sleeper сгенерирован командой (стоит в go:generate ниже):
//
//	mockgen -source=retry.go -destination=mock_sleeper_test.go -package=main Sleeper
//
// Мок лежит в _test.go файле: в продовую сборку он не попадает.
//
//go:generate go run go.uber.org/mock/mockgen -source=retry.go -destination=mock_sleeper_test.go -package=main Sleeper
func main() {}
