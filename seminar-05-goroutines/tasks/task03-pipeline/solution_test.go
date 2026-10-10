package pipeline

import (
	"reflect"
	"testing"
	"time"
)

// source превращает слайс в закрывающийся канал.
func source(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

// collect читает канал до закрытия с защитой от дедлока.
func collect(t *testing.T, ch <-chan int) []int {
	t.Helper()
	done := make(chan []int, 1)
	go func() {
		var out []int
		for v := range ch {
			out = append(out, v)
		}
		done <- out
	}()

	select {
	case out := <-done:
		return out
	case <-time.After(2 * time.Second):
		t.Fatal("превышен таймаут 2 с: выходной канал не закрыт (или этап не читает вход до конца)")
		return nil
	}
}

func TestMap(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		f    func(int) int
		want []int
	}{
		{name: "удвоение", nums: []int{1, 2, 3}, f: func(n int) int { return n * 2 }, want: []int{2, 4, 6}},
		{name: "порядок сохраняется", nums: []int{3, 1, 2}, f: func(n int) int { return n * 10 }, want: []int{30, 10, 20}},
		{name: "отрицательные", nums: []int{-2, 0, 2}, f: func(n int) int { return -n }, want: []int{2, 0, -2}},
		{name: "пустой вход", nums: nil, f: func(n int) int { return n }, want: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := collect(t, Map(source(tc.nums...), tc.f))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Map(%v) = %v, ожидалось %v", tc.nums, got, tc.want)
			}
		})
	}
}

// Map обязан вернуться сразу: вызов без читателя выходного канала
// не блокируется.
func TestMapReturnsImmediately(t *testing.T) {
	done := make(chan struct{})
	go func() {
		Map(source(1, 2, 3), func(n int) int { return n })
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Map не вернулся без читателя: обработку нужно делать в отдельной горутине")
	}
}

func TestFilter(t *testing.T) {
	even := func(n int) bool { return n%2 == 0 }

	testCases := []struct {
		name string
		nums []int
		keep func(int) bool
		want []int
	}{
		{name: "чётные", nums: []int{1, 2, 3, 4, 5, 6}, keep: even, want: []int{2, 4, 6}},
		{name: "ничего не проходит", nums: []int{1, 3, 5}, keep: even, want: nil},
		{name: "проходит всё", nums: []int{7, 8}, keep: func(int) bool { return true }, want: []int{7, 8}},
		{name: "пустой вход", nums: nil, keep: even, want: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := collect(t, Filter(source(tc.nums...), tc.keep))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Filter(%v) = %v, ожидалось %v", tc.nums, got, tc.want)
			}
		})
	}
}

// Этапы соединяются в конвейер: close входа «проезжает» по всей цепочке.
func TestChainedPipeline(t *testing.T) {
	double := func(n int) int { return n * 2 }
	even := func(n int) bool { return n%2 == 0 }

	got := collect(t, Filter(Map(source(1, 2, 3, 4, 5), double), even))

	want := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Filter(Map(1..5)) = %v, ожидалось %v", got, want)
	}
}
