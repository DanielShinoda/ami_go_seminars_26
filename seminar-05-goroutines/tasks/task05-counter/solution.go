package counter

// Counter — счётчик, безопасный для конкурентного использования.
//
// Нулевое значение — готовый к работе счётчик со значением 0:
//
//	var c counter.Counter
//	c.Inc()
type Counter struct {
	// Объявите здесь поля, которые нужны вашей реализации.
}

// Inc увеличивает счётчик на единицу.
func (c *Counter) Inc() {
	panic("TODO: реализуйте Counter.Inc")
}

// Add увеличивает счётчик на delta (delta может быть отрицательной).
func (c *Counter) Add(delta int) {
	panic("TODO: реализуйте Counter.Add")
}

// Value возвращает текущее значение счётчика.
func (c *Counter) Value() int {
	panic("TODO: реализуйте Counter.Value")
}
