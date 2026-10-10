package counter

import (
	"sync"
	"testing"
)

func TestCounterSequential(t *testing.T) {
	testCases := []struct {
		name string
		incs int
		adds []int
		want int
	}{
		{name: "нулевое значение готово к работе", incs: 0, adds: nil, want: 0},
		{name: "только Inc", incs: 3, adds: nil, want: 3},
		{name: "только Add", incs: 0, adds: []int{5, -2, 10}, want: 13},
		{name: "Inc и Add вместе", incs: 2, adds: []int{-5}, want: -3},
		{name: "Add нуля", incs: 0, adds: []int{0}, want: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var c Counter
			for range tc.incs {
				c.Inc()
			}
			for _, d := range tc.adds {
				c.Add(d)
			}
			if got := c.Value(); got != tc.want {
				t.Errorf("Value() = %d, ожидалось %d", got, tc.want)
			}
		})
	}
}

// Главное свойство: счётчик не теряет обновления при конкурентном
// доступе. Этот тест нужно прогонять с детектором гонок:
//
//	go test -race ./seminar-05-goroutines/tasks/task05-counter
func TestCounterConcurrent(t *testing.T) {
	const goroutines = 20
	const incs = 500

	var c Counter
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

	want := goroutines * incs
	if got := c.Value(); got != want {
		t.Errorf("Value() = %d, ожидалось %d — обновления теряются, нужна синхронизация", got, want)
	}
}

// Add и Value тоже конкурентно безопасны.
func TestCounterConcurrentAddAndValue(t *testing.T) {
	const goroutines = 10
	const adds = 200

	var c Counter
	var wg sync.WaitGroup

	for range goroutines {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for range adds {
				c.Add(2)
			}
		}()
		go func() {
			defer wg.Done()
			for range adds {
				_ = c.Value() // чтение не должно мешать записям и наоборот
			}
		}()
	}
	wg.Wait()

	want := goroutines * adds * 2
	if got := c.Value(); got != want {
		t.Errorf("Value() = %d, ожидалось %d", got, want)
	}
}
