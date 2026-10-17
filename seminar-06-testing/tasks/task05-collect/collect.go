package collect

import "time"

// Collect читает значения из ch, пока канал не закроется или не истечёт
// timeout.
//
// Возвращает все прочитанные значения (в порядке чтения) и true, если
// канал закрылся до истечения таймаута. Если сработал таймаут, возвращает
// то, что успело прочитаться, и false.
//
// Это тестовый хелпер: им проверяют, что конвейер из семинара 5 дошёл
// до конца и закрыл выходной канал, а не завис на половине.
func Collect(ch <-chan int, timeout time.Duration) (vals []int, closed bool) {
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return vals, true
			}
			vals = append(vals, v)
		case <-deadline.C:
			return vals, false
		}
	}
}
