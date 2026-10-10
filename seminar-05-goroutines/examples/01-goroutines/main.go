// Горутины: запуск, завершение программы и sync.WaitGroup.
//
//	go run ./seminar-05-goroutines/examples/01-goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

// Горутина — это легковесный поток, которым управляет рантайм Go, а не ОС.
// Горутин можно запустить сотни тысяч: каждая начинается со стека в
// несколько килобайт, который растёт по мере надобности.

func worker(name string, d time.Duration) {
	time.Sleep(d)
	fmt.Println(name, "готов")
}

func main() {
	// Главная ловушка: функция main завершается — программа умирает
	// целиком, не дожидаясь остальных горутин.
	go worker("без ожидания", 50*time.Millisecond)

	// sync.WaitGroup — счётчик «сколько горутин осталось дождаться».
	// Add ДО запуска горутины, Done — внутри неё (через defer), Wait — там,
	// где нужно дождаться всех.
	var wg sync.WaitGroup
	for i := range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(fmt.Sprintf("воркер %d", i), time.Duration(i+1)*30*time.Millisecond)
		}()
	}

	fmt.Println("main: все запущены, жду")
	wg.Wait()
	fmt.Println("main: все закончили")

	// Горутины выполняются КОНКУРЕНТНО, порядок вывода не гарантирован.
	// Общее время — время самой медленной, а не сумма: убедитесь,
	// подставив time.Now() и измерив.
}
