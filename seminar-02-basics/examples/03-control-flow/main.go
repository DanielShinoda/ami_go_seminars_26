// Управляющие конструкции: if, for, switch.
// Их всего три — в Go нет while, do-while и тернарного оператора.
//
// Запуск:
//
//	go run ./seminar-02-basics/examples/03-control-flow
package main

import "fmt"

func main() {
	ifExamples()
	forExamples()
	switchExamples()
}

func ifExamples() {
	fmt.Println("--- if ---")

	x := 7

	// Скобок вокруг условия нет, фигурные скобки обязательны всегда.
	if x > 5 {
		fmt.Println("больше пяти")
	} else if x == 5 {
		fmt.Println("ровно пять")
	} else {
		fmt.Println("меньше пяти")
	}

	// if с инициализатором: переменная видна только внутри if/else.
	// Это основная идиома Go — так сужают область видимости.
	if half := x / 2; half > 3 {
		fmt.Println("половина больше трёх:", half)
	} else {
		fmt.Println("половина:", half)
	}
	// здесь half уже не существует
}

func forExamples() {
	fmt.Println("--- for ---")

	// 1. Классический счётчик, как в C++.
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2. for как while: одно условие.
	n := 3
	for n > 0 {
		fmt.Print(n, " ")
		n--
	}
	fmt.Println()

	// 3. Бесконечный цикл + break. Так пишут while(true).
	count := 0
	for {
		count++
		if count == 3 {
			break
		}
	}
	fmt.Println("count =", count)

	// 4. range по числу (Go 1.22+): i пробегает 0..n-1.
	for i := range 3 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 5. range по коллекции: индекс и КОПИЯ значения.
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Printf("[%d]=%d ", i, v)
	}
	fmt.Println()

	// Изменение v внутри цикла не меняет слайс — v это копия.
	for _, v := range nums {
		v *= 2
	}
	fmt.Println("после range с v *= 2:", nums) // [10 20 30] — не изменился!

	// Чтобы изменить — обращаемся по индексу.
	for i := range nums {
		nums[i] *= 2
	}
	fmt.Println("после nums[i] *= 2:", nums)

	// continue и метки. Метка нужна, чтобы выйти из ВНЕШНЕГО цикла.
outer:
	for i := range 3 {
		for j := range 3 {
			if j > i {
				continue outer
			}
			if i == 2 {
				break outer
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()
}

func switchExamples() {
	fmt.Println("--- switch ---")

	day := 6

	// В Go НЕТ проваливания: break в конце case писать не нужно.
	// Это ровно наоборот по сравнению с C++.
	switch day {
	case 1, 2, 3, 4, 5: // несколько значений в одном case
		fmt.Println("будни")
	case 6, 7:
		fmt.Println("выходные")
	default:
		fmt.Println("такого дня нет")
	}

	// switch без выражения = лестница if/else if.
	// Это и есть замена отсутствующему тернарному оператору для длинных цепочек.
	score := 87
	switch {
	case score >= 90:
		fmt.Println("отлично")
	case score >= 75:
		fmt.Println("хорошо")
	case score >= 60:
		fmt.Println("удовлетворительно")
	default:
		fmt.Println("неудовлетворительно")
	}

	// switch с инициализатором.
	switch mod := score % 2; mod {
	case 0:
		fmt.Println("чётный балл")
	default:
		fmt.Println("нечётный балл")
	}

	// fallthrough — явно провалиться в следующий case. Нужен редко.
	switch 1 {
	case 1:
		fmt.Println("один")
		fallthrough
	case 2:
		fmt.Println("и два (сюда попали через fallthrough)")
	case 3:
		fmt.Println("три — не выполнится")
	}
}
