package chansum

import (
	"testing"
	"time"
)

// withTimeout запускает f и возвращает результат; если f не завершилась
// за 2 секунды, тест падает — чаще всего это «забыли закрыть канал».
func withTimeout(t *testing.T, f func() int) int {
	t.Helper()
	done := make(chan int, 1)
	go func() { done <- f() }()

	select {
	case v := <-done:
		return v
	case <-time.After(2 * time.Second):
		t.Fatal("превышен таймаут 2 с: похоже на дедлок (канал не закрыт или чтение/запись не в горутине)")
		return 0
	}
}

func TestGenerator(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "несколько чисел", nums: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "отрицательные и ноль", nums: []int{-5, 0, 5}, want: []int{-5, 0, 5}},
		{name: "одно число", nums: []int{7}, want: []int{7}},
		{name: "без аргументов", nums: nil, want: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ch := Generator(tc.nums...)

			got := withTimeout(t, func() int {
				count := 0
				i := 0
				for v := range ch {
					if i >= len(tc.want) {
						t.Errorf("канал отдал лишнее значение %d", v)
						return -1
					}
					if v != tc.want[i] {
						t.Errorf("элемент %d = %d, ожидалось %d", i, v, tc.want[i])
					}
					i++
					count++
				}
				return count
			})

			if got != len(tc.want) {
				t.Errorf("канал закрылся после %d значений, ожидалось %d", got, len(tc.want))
			}
		})
	}
}

// Generator обязан вернуться сразу: вызов без читателя не блокируется.
func TestGeneratorReturnsImmediately(t *testing.T) {
	done := make(chan struct{})
	go func() {
		Generator(1, 2, 3)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Generator не вернулся без читателя: отправку нужно делать в отдельной горутине")
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{name: "несколько чисел", nums: []int{1, 2, 3, 4}, want: 10},
		{name: "отрицательные", nums: []int{-10, 5, -5}, want: -10},
		{name: "одно число", nums: []int{42}, want: 42},
		{name: "пустой канал", nums: nil, want: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ch := make(chan int)
			go func() {
				defer close(ch)
				for _, n := range tc.nums {
					ch <- n
				}
			}()

			got := withTimeout(t, func() int { return Sum(ch) })
			if got != tc.want {
				t.Errorf("Sum(%v) = %d, ожидалось %d", tc.nums, got, tc.want)
			}
		})
	}
}

// Generator и Sum соединяются напрямую: сумма потока.
func TestGeneratorAndSumTogether(t *testing.T) {
	got := withTimeout(t, func() int { return Sum(Generator(1, 2, 3, 4, 5)) })
	if got != 15 {
		t.Errorf("Sum(Generator(1, 2, 3, 4, 5)) = %d, ожидалось 15", got)
	}
}
