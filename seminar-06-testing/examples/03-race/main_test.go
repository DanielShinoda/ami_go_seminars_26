// Тестирование конкурентного кода и детектор гонок.
//
//	go test ./seminar-06-testing/examples/03-race
//	go test -race ./seminar-06-testing/examples/03-race
package main

import (
	"sync"
	"testing"
)

// SafeCounter — потокобезопасный счётчик (см. семинар 5, задача 5).
type SafeCounter struct {
	mu sync.Mutex
	n  int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// Как тестировать конкурентный код: не «один раз дёрнуть», а СТРЕСС —
// много горутин, много операций, проверка инварианта в конце.
// Чем больше итераций, тем выше шанс поймать гонку даже без -race.
func TestSafeCounterStress(t *testing.T) {
	const goroutines = 20
	const incs = 1000

	var c SafeCounter
	var wg sync.WaitGroup
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range incs {
				c.Inc()
			}
		}()
	}
	wg.Wait()

	if got, want := c.Value(), goroutines*incs; got != want {
		t.Errorf("Value() = %d, ожидалось %d — обновления теряются", got, want)
	}
}

// Важно понимать: этот тест может пройти и с гонкой — если планировщику
// «повезло». Поэтому конкурентный код ВСЕГДА прогоняют с детектором:
//
//	go test -race ./...
//
// Детектор следит за доступами к памяти и находит гонку, даже если
// результат численно правильный. Цена — замедление в 5–10 раз, поэтому
// -race включают в CI, а не при каждом локальном прогоне.

// Так выглядит гонка, которую поймает -race. Раскомментируйте и
// прогоните с флагом — детектор напечатает оба стека доступа:
//
//	func TestRacyCounter(t *testing.T) {
//		n := 0
//		var wg sync.WaitGroup
//		for range 10 {
//			wg.Add(1)
//			go func() {
//				defer wg.Done()
//				for range 1000 {
//					n++ // DATA RACE: WARNING в выводе go test -race
//				}
//			}()
//		}
//		wg.Wait()
//	}
//
// Полезные флаги рядом с -race:
//
//	go test -count=5 ./...   прогнать тесты 5 раз (ловит «мигающие»)
//	go test -shuffle=on ./... перемешать порядок тестов (ловит скрытые
//	                          зависимости между тестами)
