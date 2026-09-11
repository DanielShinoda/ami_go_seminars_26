// Что на самом деле делает `go build`.
//
// Демонстрация на семинаре:
//
//	go run ./seminar-01-intro/examples/02-build-and-run
//	go build -o /tmp/whereami ./seminar-01-intro/examples/02-build-and-run
//	ls -lh /tmp/whereami        # ~2 МБ: рантайм и GC внутри бинаря
//	file /tmp/whereami
//	/tmp/whereami               # запускается без установленного Go
//
// Кросс-компиляция — без единого дополнительного инструмента:
//
//	GOOS=linux   GOARCH=amd64 go build -o /tmp/whereami-linux ./seminar-01-intro/examples/02-build-and-run
//	GOOS=windows GOARCH=amd64 go build -o /tmp/whereami.exe   ./seminar-01-intro/examples/02-build-and-run
//	file /tmp/whereami-linux /tmp/whereami.exe
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("версия Go:    %s\n", runtime.Version())
	fmt.Printf("ОС:           %s\n", runtime.GOOS)
	fmt.Printf("архитектура:  %s\n", runtime.GOARCH)
	fmt.Printf("ядер CPU:     %d\n", runtime.NumCPU())
}
