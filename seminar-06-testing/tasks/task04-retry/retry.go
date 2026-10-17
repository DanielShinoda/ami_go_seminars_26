package retry

import "time"

// Do вызывает op, пока та не вернёт nil или не кончатся попытки
// (их число — opts.Attempts, см. types.go).
//
// После неудачной попытки с номером i (нумерация с 1), если она не
// последняя, вызывается opts.Sleep с паузой i * 10 миллисекунд:
// 10 мс после первой, 20 мс после второй и так далее.
//
// При первом же успехе возвращается nil. Если все попытки неудачны,
// возвращается ошибка последней попытки.
func Do(op func() error, opts Options) error {
	if opts.Attempts < 1 {
		opts.Attempts = 1
	}

	var err error
	for i := 1; i <= opts.Attempts; i++ {
		if err = op(); err == nil {
			return nil
		}
		if i < opts.Attempts && opts.Sleep != nil {
			opts.Sleep(time.Duration(i) * 10 * time.Millisecond)
		}
	}
	return err
}
