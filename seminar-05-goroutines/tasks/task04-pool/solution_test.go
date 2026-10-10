package pool

import (
	"reflect"
	"sync/atomic"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	square := func(n int) int { return n * n }

	testCases := []struct {
		name    string
		jobs    []int
		workers int
		want    []int
	}{
		{
			name:    "несколько задач, один воркер",
			jobs:    []int{1, 2, 3},
			workers: 1,
			want:    []int{1, 4, 9},
		},
		{
			name:    "несколько задач, несколько воркеров",
			jobs:    []int{1, 2, 3, 4, 5},
			workers: 3,
			want:    []int{1, 4, 9, 16, 25},
		},
		{
			name:    "воркеров больше, чем задач",
			jobs:    []int{2, 4},
			workers: 100,
			want:    []int{4, 16},
		},
		{
			name:    "workers = 0 трактуется как 1",
			jobs:    []int{3},
			workers: 0,
			want:    []int{9},
		},
		{
			name:    "отрицательный workers трактуется как 1",
			jobs:    []int{3},
			workers: -2,
			want:    []int{9},
		},
		{name: "nil-слайс", jobs: nil, workers: 4, want: []int{}},
		{name: "пустой слайс", jobs: []int{}, workers: 4, want: []int{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Process(tc.jobs, square, tc.workers)

			if got == nil {
				t.Fatalf("Process = nil, ожидался не-nil слайс")
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Process(%v, square, %d) = %v, ожидалось %v", tc.jobs, tc.workers, got, tc.want)
			}
		})
	}
}

// Порядок результатов — порядок задач, даже если медленные задачи идут
// первыми.
func TestProcessPreservesOrder(t *testing.T) {
	jobs := []int{3, 1, 2}
	// Чем больше вход, тем дольше работа.
	slowDouble := func(n int) int {
		time.Sleep(time.Duration(n) * 20 * time.Millisecond)
		return n * 2
	}

	got := Process(jobs, slowDouble, 3)
	want := []int{6, 2, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Process = %v, ожидалось %v (порядок задач, а не завершения)", got, want)
	}
}

// Параллелизм ограничен workers: замеряем максимум одновременно
// выполняющихся f.
func TestProcessLimitsConcurrency(t *testing.T) {
	const workers = 3

	var current atomic.Int32
	var maxSeen atomic.Int32

	jobs := make([]int, 12)
	for i := range jobs {
		jobs[i] = i
	}

	f := func(n int) int {
		cur := current.Add(1)
		defer current.Add(-1)
		for {
			old := maxSeen.Load()
			if cur <= old || maxSeen.CompareAndSwap(old, cur) {
				break
			}
		}
		time.Sleep(10 * time.Millisecond)
		return n
	}

	Process(jobs, f, workers)

	if got := maxSeen.Load(); got > workers {
		t.Errorf("одновременно работали %d горутин, ожидалось не более %d", got, workers)
	}
	if got := maxSeen.Load(); got < 2 {
		t.Errorf("одновременно работали максимум %d горутин — похоже на последовательное выполнение", got)
	}
}

// Каждая задача обрабатывается ровно один раз.
func TestProcessRunsEachJobOnce(t *testing.T) {
	const n = 50
	counts := make([]atomic.Int32, n)

	jobs := make([]int, n)
	for i := range jobs {
		jobs[i] = i
	}

	Process(jobs, func(x int) int {
		counts[x].Add(1)
		return x
	}, 5)

	for i := range n {
		if c := counts[i].Load(); c != 1 {
			t.Errorf("задача %d обработана %d раз, ожидалось 1", i, c)
		}
	}
}
