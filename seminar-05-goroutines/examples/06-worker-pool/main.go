// Пул воркеров: фиксированное число горутин разбирает очередь задач.
//
//	go run ./seminar-05-goroutines/examples/06-worker-pool
package main

import (
	"fmt"
	"sync"
	"time"
)

// Пул воркеров — самый ходовой паттерн конкурентности в Go: вместо
// горутины на каждую задачу запускаем фиксированные N горутин, которые
// разбирают общий канал задач. Так нагрузка ограничена, а задач может
// быть хоть миллион.

func main() {
	const (
		jobCount = 10
		workers  = 3
	)

	jobs := make(chan int)
	results := make(chan string)

	// Запускаем воркеров. Каждый читает jobs, пока канал не закроют.
	var wg sync.WaitGroup
	for w := range workers {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				time.Sleep(30 * time.Millisecond) // работа
				results <- fmt.Sprintf("задача %2d -> воркер %d", job, id)
			}
		}(w)
	}

	// Закрываем results, когда ВСЕ воркеры закончили: отдельная горутина.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Отправляем задачи и сразу идём читать результаты.
	go func() {
		for j := range jobCount {
			jobs <- j
		}
		close(jobs) // сигнал воркерам: задач больше не будет
	}()

	for res := range results {
		fmt.Println(res)
	}

	// Ключевые ходы паттерна:
	//  1. воркеры живут в range по каналу задач — завершаются по close;
	//  2. результаты закрывает тот, кто дождался WaitGroup всех воркеров;
	//  3. отправитель задач и читатель результатов не блокируют друг друга.
}
