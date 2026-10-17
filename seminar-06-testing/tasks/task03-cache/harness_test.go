// Этот файл менять не нужно — это приёмка ваших тестов.
//
// Харнесс собирает тестовый пакет из реализации (эталонной cache.go
// или одной из сломанных variants/*.go.txt) и вашего cache_test.go и
// прогоняет по нему go test в временном каталоге. Задача принята, когда
// тесты проходят на эталоне и падают на КАЖДОЙ сломанной реализации.

package cache

import (
	"context"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"
)

//go:embed variants/*.go.txt
var variants embed.FS

const (
	implFile = "cache.go"      // эталон, лежит рядом
	testFile = "cache_test.go" // ваши тесты

	// Предел времени на один прогон. Зависший тест (например, сломанная
	// реализация не выходит по таймауту) считается провалом: дефект всё
	// равно обнаружен, просто сообщением станет "прогон завис".
	runTimeout = 30 * time.Second
)

var testFuncRE = regexp.MustCompile(`(?m)^func\s+Test\w*\s*\(`)

func TestYourTestsRejectBrokenImplementations(t *testing.T) {
	student, err := os.ReadFile(testFile)
	if err != nil || !testFuncRE.Match(student) {
		t.Fatalf("в %s нет ни одного теста (func Test...). Условие — в шапке файла.", testFile)
	}

	impl, err := os.ReadFile(implFile)
	if err != nil {
		t.Fatalf("не читается %s: %v", implFile, err)
	}

	type variant struct {
		name   string
		src    []byte
		note   string
		broken bool
	}

	cases := []variant{{
		name: implFile + " (эталон)",
		src:  impl,
		note: "правильная реализация — ваши тесты обязаны проходить",
	}}

	entries, err := variants.ReadDir("variants")
	if err != nil {
		t.Fatalf("не читается variants/: %v", err)
	}
	for _, e := range entries {
		src, err := variants.ReadFile("variants/" + e.Name())
		if err != nil {
			t.Fatalf("не читается variants/%s: %v", e.Name(), err)
		}
		cases = append(cases, variant{
			name:   "variants/" + e.Name(),
			src:    src,
			note:   firstCommentLine(src),
			broken: true,
		})
	}

	var missed []string
	for _, tc := range cases {
		passed, output := runStudentTests(t, student, tc.src)
		switch {
		case !tc.broken && !passed:
			t.Errorf("тесты НЕ ПРОХОДЯТ на эталонной реализации — сначала добейтесь зелёного прогона по условию.\nвывод прогона:\n%s", trimOutput(output))
		case tc.broken && passed:
			missed = append(missed, tc.name+" — "+tc.note)
		}
	}

	if len(missed) > 0 {
		t.Errorf("слабые тесты: эти сломанные реализации прошли ваш прогон незамеченными:\n  - %s\nДопишите тесты, которые их отклоняют.", strings.Join(missed, "\n  - "))
	}
}

// runStudentTests собирает временный пакет из реализации impl и вашего
// файла тестов и прогоняет go test. Возвращает, прошёл ли прогон, и вывод.
func runStudentTests(t *testing.T, student, impl []byte) (passed bool, output string) {
	t.Helper()
	dir := t.TempDir()

	files := map[string][]byte{
		"impl.go":         impl,
		"student_test.go": student,
	}
	if types, err := os.ReadFile("types.go"); err == nil {
		files["types.go"] = types
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), content, 0o644); err != nil {
			t.Fatalf("не пишется временный файл: %v", err)
		}
	}
	goMod := []byte("module variantcheck\n\ngo 1.22\n")
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), goMod, 0o644); err != nil {
		t.Fatalf("не пишется временный файл: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), runTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, goTool(), "test", "-count=1", ".")
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOTOOLCHAIN=local", "GOPROXY=off")
	out, runErr := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		return false, fmt.Sprintf("прогон завис и был остановлен по таймауту %s", runTimeout)
	}
	return runErr == nil, strings.TrimSpace(string(out))
}

func goTool() string {
	if p, err := exec.LookPath("go"); err == nil {
		return p
	}
	return filepath.Join(runtime.GOROOT(), "bin", "go")
}

func firstCommentLine(src []byte) string {
	for _, line := range strings.Split(string(src), "\n") {
		if strings.HasPrefix(line, "//") {
			return strings.TrimSpace(strings.TrimPrefix(line, "//"))
		}
	}
	return ""
}

func trimOutput(output string) string {
	const limit = 2000
	if len(output) <= limit {
		return output
	}
	return output[:limit/2] + "\n...\n" + output[len(output)-limit/2:]
}
