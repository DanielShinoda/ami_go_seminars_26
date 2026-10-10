package parallel

import (
	"reflect"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name string
		fns  []func() int
		want []int
	}{
		{
			name: "несколько функций",
			fns: []func() int{
				func() int { return 1 },
				func() int { return 2 },
				func() int { return 3 },
			},
			want: []int{1, 2, 3},
		},
		{
			name: "порядок результатов совпадает с порядком функций, а не с порядком завершения",
			fns: []func() int{
				func() int { time.Sleep(40 * time.Millisecond); return 1 },
				func() int { time.Sleep(20 * time.Millisecond); return 2 },
				func() int { return 3 },
			},
			want: []int{1, 2, 3},
		},
		{
			name: "одна функция",
			fns:  []func() int{func() int { return 42 }},
			want: []int{42},
		},
		{name: "nil-слайс", fns: nil, want: []int{}},
		{name: "пустой слайс", fns: []func() int{}, want: []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Run(tc.fns)

			if got == nil {
				t.Fatalf("Run = nil, ожидался не-nil слайс")
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Run = %v, ожидалось %v", got, tc.want)
			}
		})
	}
}

// Run обязан выполнять функции конкурентно: десять функций по 50 мс
// укладываются заметно меньше, чем в 500 мс.
func TestRunIsConcurrent(t *testing.T) {
	const n = 10
	const sleep = 50 * time.Millisecond

	fns := make([]func() int, n)
	for i := range fns {
		fns[i] = func() int {
			time.Sleep(sleep)
			return i
		}
	}

	start := time.Now()
	got := Run(fns)
	elapsed := time.Since(start)

	if len(got) != n {
		t.Fatalf("Run вернул %d результатов, ожидалось %d", len(got), n)
	}
	for i, v := range got {
		if v != i {
			t.Errorf("Run[%d] = %d, ожидалось %d", i, v, i)
		}
	}

	limit := n * sleep / 2
	if elapsed > limit {
		t.Errorf("Run занял %v — похоже на последовательное выполнение (лимит %v)", elapsed, limit)
	}
}

// Каждая функция вызывается ровно один раз.
func TestRunCallsEachFuncOnce(t *testing.T) {
	const n = 8
	calls := make([]int, n)

	fns := make([]func() int, n)
	for i := range fns {
		fns[i] = func() int {
			calls[i]++
			return 0
		}
	}

	Run(fns)

	for i, c := range calls {
		if c != 1 {
			t.Errorf("функция %d вызвана %d раз, ожидалось 1", i, c)
		}
	}
}
