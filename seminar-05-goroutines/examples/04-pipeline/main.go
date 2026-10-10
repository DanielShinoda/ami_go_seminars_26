// Конвейер (pipeline): каналы, соединённые в цепочку обработки.
//
//	go run ./seminar-05-goroutines/examples/04-pipeline
package main

import (
	"fmt"
	"strings"
)

// Конвейер — это функции-этапы, каждый из которых читает входной канал,
// что-то делает и пишет в выходной. Этапы работают параллельно: пока один
// обрабатывает третий элемент, другой уже читает четвёртый.

// generator превращает список строк в канал: отправляет и закрывает.
func generator(lines ...string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, line := range lines {
			out <- line
		}
	}()
	return out
}

// filter — этап конвейера: пропускает только строки, содержащие needle.
func filter(in <-chan string, needle string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for line := range in {
			if strings.Contains(line, needle) {
				out <- line
			}
		}
	}()
	return out
}

// upper — этап конвейера: переводит строки в верхний регистр.
func upper(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for line := range in {
			out <- strings.ToUpper(line)
		}
	}()
	return out
}

func main() {
	// Собираем конвейер: каждый этап — горутина, соединены каналами.
	lines := generator(
		"go — простой язык",
		"rust — сложный язык",
		"go быстро собирается",
		"python медленный",
	)
	pipe := upper(filter(lines, "go"))

	// Читаем до закрытия: когда generator закроет свой выход, close
	// «проедет» по всей цепочке через defer close(out) каждого этапа.
	for line := range pipe {
		fmt.Println(line)
	}

	// Правило жизни конвейера: этап обязан ДОЧИТАТЬ вход до конца
	// и ЗАКРЫТЬ выход. Если бросить чтение на середине, пишущий этап
	// зависнет на отправке навсегда — это утечка горутины.
}
