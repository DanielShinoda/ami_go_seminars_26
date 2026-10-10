// Небуферизованные каналы: рандеву, направления, close и range.
//
//	go run ./seminar-05-goroutines/examples/02-channels
package main

import (
	"fmt"
	"time"
)

// Канал — типизированная труба между горутинами. Небуферизованный канал
// синхронизирует: отправка блокируется, пока кто-то не получит значение,
// и наоборот. Это «рандеву» — передача из рук в руки.

// generator ТОЛЬКО отправляет: тип chan<- int запрещает чтение из канала
// внутри функции — это проверяется компилятором.
func generator(out chan<- int, count int) {
	for i := range count {
		out <- i
	}
	// Кто отправляет, тот и закрывает. close — это сигнал «значений
	// больше не будет». Отправка в закрытый канал — паника.
	close(out)
}

// sum ТОЛЬКО получает: <-chan int запрещает отправку и закрытие.
// range по каналу читает значения, пока канал не закроют.
func sum(in <-chan int) int {
	total := 0
	for v := range in {
		total += v
	}
	return total
}

func main() {
	ch := make(chan int) // небуферизованный: ёмкость 0

	go generator(ch, 5)

	// sum блокируется в range и завершится, когда generator закроет канал.
	fmt.Println("сумма 0..4 =", sum(ch))

	// Чтение из закрытого канала не паникует: возвращает нулевое значение
	// и ok == false. Именно так отличают «конец потока» от «прислали 0».
	v, ok := <-ch
	fmt.Printf("после close: v=%d ok=%t\n", v, ok)

	// Рандеву наглядно: отправитель и получатель встречаются.
	sync := make(chan string)
	go func() {
		time.Sleep(30 * time.Millisecond)
		sync <- "привет из горутины"
	}()
	fmt.Println("main ждёт сообщение…")
	fmt.Println("получено:", <-sync) // main заблокировался и дождался
}
