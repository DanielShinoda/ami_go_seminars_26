// Время в тестах: почему sleep — это флаки, и как ждать правильно.
//
//	go test -v ./seminar-06-testing/examples/04-time-and-flakes
package main

import (
	"sync/atomic"
	"testing"
	"time"
)

// Флаки (flaky test) — тест, который падает и проходит случайно.
// Главный источник флаки в конкурентных тестах — time.Sleep как способ
// «дождаться» другую горутину: на загруженной машине CI сна может не
// хватить, на быстрой — тест даром жжёт время.

func backgroundWorker(done *atomic.Bool) {
	go func() {
		time.Sleep(20 * time.Millisecond) // работа
		done.Store(true)
	}()
}

// ПЛОХОЙ тест выглядел бы так (не пишите так):
//
//	func TestWorkerBad(t *testing.T) {
//		var done atomic.Bool
//		backgroundWorker(&done)
//		time.Sleep(100 * time.Millisecond) // «наверняка хватит» — нет
//		if !done.Load() { t.Fatal(...) }
//	}
//
// На перегруженном раннере воркер может не успеть за 100 мс — тест упал
// ни за что. А на нормальной машине 80 мс выброшены впустую.

// eventually — хелпер-опрос: проверяет условие сразу, потом с коротким
// интервалом, пока не выйдет бюджет времени. Вместо «ждём 100 мс наугад»
// — «ждём ДО 2 с, но выходим, как только готово». Быстрый случай быстр,
// медленный — надёжен. Именно такой хелпер вы напишете в задаче 2.
func eventually(cond func() bool, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for {
		if cond() {
			return true
		}
		if time.Now().After(deadline) {
			return false
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func TestWorkerGood(t *testing.T) {
	var done atomic.Bool
	backgroundWorker(&done)

	if !eventually(done.Load, 2*time.Second) {
		t.Fatal("воркер не завершился за 2 с")
	}
}

// Второй приём — ждать СИГНАЛ, а не время. Если код может сообщить
// о завершении (канал, WaitGroup), тест обязан пользоваться именно этим:
// никакого опроса и таймаута вообще не нужно.
func TestWorkerWithSignal(t *testing.T) {
	finished := make(chan struct{})
	go func() {
		defer close(finished)
		time.Sleep(20 * time.Millisecond) // работа
	}()

	// select + time.After — таймаут только как страховка от зависшего
	// теста, а не как механизм ожидания.
	select {
	case <-finished:
	case <-time.After(2 * time.Second):
		t.Fatal("воркер завис: нет сигнала за 2 с")
	}
}

// Иерархия ожидания в тестах: сигнал (канал/WaitGroup) > опрос с
// дедлайном (eventually) > sleep. Sleep допустим только ВНУТРИ цикла
// опроса и в подделках времени.
