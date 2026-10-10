// Буферизованные каналы и select: ожидание нескольких каналов сразу.
//
//	go run ./seminar-05-goroutines/examples/03-buffered-select
package main

import (
	"fmt"
	"time"
)

func fetch(name string, d time.Duration, out chan<- string) {
	time.Sleep(d) // имитация сетевого запроса
	out <- name + " готов"
}

func main() {
	// Буферизованный канал: отправка не блокируется, пока буфер не полон.
	buf := make(chan int, 3)
	buf <- 1
	buf <- 2
	buf <- 3
	fmt.Println("записали 3 значения без получателя; ёмкость:", cap(buf), "занято:", len(buf))
	// buf <- 4 здесь заблокировал бы main навсегда — дедлок в одной
	// горутине Go замечает и паникует. Не раскомментируйте без получателя.

	// select ждёт, пока ОДИН из case станет выполнимым. Если готовы
	// несколько — выбирает случайный.
	a := make(chan string)
	b := make(chan string)
	go fetch("быстрый", 20*time.Millisecond, a)
	go fetch("медленный", 60*time.Millisecond, b)

	for range 2 {
		select {
		case msg := <-a:
			fmt.Println("из a:", msg)
		case msg := <-b:
			fmt.Println("из b:", msg)
		}
	}

	// select + time.After — стандартный способ поставить таймаут на
	// операцию с каналом.
	slow := make(chan string)
	go fetch("очень медленный", time.Second, slow)

	select {
	case msg := <-slow:
		fmt.Println(msg)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("не дождались: таймаут 50 мс")
	}

	// select с default — неблокирующее чтение: «если есть — забери».
	select {
	case v := <-buf:
		fmt.Println("неблокирующе забрали из буфера:", v)
	default:
		fmt.Println("буфер пуст")
	}
}
