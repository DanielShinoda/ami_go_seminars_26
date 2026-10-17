// Окружение теста: t.TempDir, t.Cleanup, testdata, t.Parallel.
//
//	go test -v ./seminar-06-testing/examples/02-test-env
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// CountLines считает строки в файле — функция, которую тестируем.
func CountLines(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	if len(data) == 0 {
		return 0, nil
	}
	lines := strings.Count(string(data), "\n")
	if !strings.HasSuffix(string(data), "\n") {
		lines++ // последняя строка без завершающего перевода строки
	}
	return lines, nil
}

// t.TempDir создаёт временную директорию, которую тестовый раннер
// УДАЛИТ САМ после теста. Никаких os.MkdirTemp + defer RemoveAll.
func TestCountLinesFromTempFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "input.txt")

	if err := os.WriteFile(path, []byte("one\ntwo\nthree\n"), 0o644); err != nil {
		t.Fatalf("не записать файл: %v", err)
	}

	got, err := CountLines(path)
	if err != nil {
		t.Fatalf("CountLines(%q) ошибка = %v", path, err)
	}
	if got != 3 {
		t.Errorf("CountLines(%q) = %d, ожидалось 3", path, got)
	}
}

// testdata — директория с фикстурами рядом с тестом. Инструменты Go
// игнорируют её при сборке: это «данные для тестов» по соглашению.
// Подходит для больших входов, которые неудобно зашивать в код.
func TestCountLinesFromTestdata(t *testing.T) {
	got, err := CountLines(filepath.Join("testdata", "greeting.txt"))
	if err != nil {
		t.Fatalf("CountLines ошибка = %v", err)
	}
	if got != 2 {
		t.Errorf("CountLines(testdata/greeting.txt) = %d, ожидалось 2", got)
	}
}

// t.Cleanup регистрирует функцию, которая выполнится ПОСЛЕ теста —
// даже если тест упал по Fatalf. Это defer для всего теста, включая
// подтесты и хелперы.
func TestCleanupRunsAfterTest(t *testing.T) {
	var cleaned bool
	t.Cleanup(func() { cleaned = true })

	if cleaned {
		t.Fatal("cleanup не должен был выполниться до конца теста")
	}
	// После выхода из теста cleaned станет true — проверить это может
	// только соседний тест, что само по себе хрупко, поэтому просто
	// помните: Cleanup всегда выполняется, порядок — обратный (LIFO).
}

// t.Parallel разрешает тесту выполняться параллельно с другими
// параллельными тестами пакета. Запуск: go test -parallel 8 ./...
// Ставится первой строкой. Тесты с t.Parallel обязаны быть
// независимыми: своя TempDir, свои данные, никаких общих переменных.
func TestParallelIndependent(t *testing.T) {
	t.Parallel()

	for range 3 {
		t.Run("независимый подтест", func(t *testing.T) {
			t.Parallel()
			dir := t.TempDir() // у каждого подтеста своя директория
			if dir == "" {
				t.Error("пустая временная директория")
			}
		})
	}
}
