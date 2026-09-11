// Переменные, константы, iota.
//
// Запуск:
//
//	go run ./seminar-02-basics/examples/01-vars-consts
package main

import "fmt"

// Константы уровня пакета. Константа в Go вычисляется на этапе компиляции
// и не имеет адреса: взять &MaxRetries нельзя.
const (
	MaxRetries = 3
	AppName    = "seminar"
)

// iota — счётчик внутри блока const: 0 в первой строке, +1 на каждой следующей.
// Это штатный способ сделать перечисление; отдельного enum в Go нет.
type Weekday int

const (
	Monday    Weekday = iota // 0
	Tuesday                  // 1 — выражение справа повторяется автоматически
	Wednesday                // 2
	Thursday                 // 3
	Friday                   // 4
)

// Тот же приём для «степеней двойки»: выражение с iota тоже повторяется.
const (
	KB = 1 << (10 * (iota + 1)) // 1 << 10
	MB                          // 1 << 20
	GB                          // 1 << 30
)

func main() {
	// 1. var с явным типом — когда тип важнее значения.
	var attempts int = 0

	// 2. var без типа — тип выводится из значения.
	var host = "localhost"

	// 3. Короткое объявление := — только ВНУТРИ функции.
	port := 8080

	fmt.Println(attempts, host, port)

	// Блок var: группируем связанные объявления.
	var (
		enabled bool    // false
		ratio   float64 // 0
		label   string  // ""
	)
	fmt.Printf("%v %v %q\n", enabled, ratio, label)

	// Множественное присваивание. Обмен значений без временной переменной:
	a, b := 1, 2
	a, b = b, a
	fmt.Println(a, b)

	// В := достаточно, чтобы ХОТЯ БЫ ОДНА переменная слева была новой.
	// Здесь port переиспользуется, а scheme объявляется.
	port, scheme := 9090, "http"
	fmt.Println(port, scheme)

	// Пустой идентификатор _ — «значение мне не нужно».
	// Так глушат ненужный результат, не нарушая правило «объявил — используй».
	_, second := divmod(17, 5)
	fmt.Println(second)

	fmt.Println(Monday, Friday) // 0 4 — под капотом это int
	fmt.Println(KB, MB, GB)
}

func divmod(a, b int) (int, int) { return a / b, a % b }
