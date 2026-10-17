package eventually

import "time"

// Eventually опрашивает cond, пока та не вернёт true или не истечёт
// timeout, и сообщает, дождались ли.
//
// Первую проверку cond выполняет немедленно, при вызове. Если cond ложна,
// проверка повторяется с небольшим интервалом, пока не истечёт timeout.
// Возвращает true, если cond стала истинной в пределах таймаута.
//
// При timeout <= 0 cond проверяется ровно один раз, без ожидания.
func Eventually(cond func() bool, timeout time.Duration) bool {
	const poll = 5 * time.Millisecond
	deadline := time.Now().Add(timeout)
	for {
		if cond() {
			return true
		}
		if !time.Now().Before(deadline) {
			return false
		}
		time.Sleep(poll)
	}
}
