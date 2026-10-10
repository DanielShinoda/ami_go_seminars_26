// Общее состояние: гонка данных, мьютекс и атомики.
//
//	go run ./seminar-05-goroutines/examples/05-shared-state
//
// Гонку ловит детектор: go run -race ./seminar-05-goroutines/examples/05-shared-state
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Гонка данных (data race) — две горутины трогают одну память, хотя бы
// одна пишет, и нет синхронизации. Результат — непредсказуемое значение,
// а не ошибка.

func main() {
	const goroutines = 100
	const incs = 1000

	// Вариант 1: без защиты. Запустите с флагом -race — детектор покажет
	// гонку. Без флага результат почти наверняка будет меньше ожидаемого:
	// инкремент — это три операции (прочитать, прибавить, записать),
	// и они переплетаются между горутинами.
	unsafe := 0
	var wg sync.WaitGroup
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range incs {
				unsafe++ // ГОНКА: чтение и запись общего unsafe
			}
		}()
	}
	wg.Wait()
	fmt.Println("без защиты:     ", unsafe, "(ожидалось", goroutines*incs, ")")

	// Вариант 2: мьютекс. Только одна горутина удерживает мьютекс
	// в каждый момент — критическая секция исполняется последовательно.
	var mu sync.Mutex
	guarded := 0
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range incs {
				mu.Lock()
				guarded++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("с мьютексом:    ", guarded)

	// Вариант 3: атомик — для простых операций над числом. Дешевле
	// мьютекса, но умеет только примитивные операции.
	var counter atomic.Int64
	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range incs {
				counter.Add(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("с атомиком:     ", counter.Load())

	// Как выбирать: одна переменная-счётчик — atomic; инвариант на
	// нескольких полях — мьютекс; передача владения данными — канал.
	// Девиз Go: «не общайтесь через память, общайтесь через каналы» —
	// но мьютекс к счётчику это не отменяет.
}
