#!/usr/bin/env bash
#
# Проверка эталонных решений преподавателя.
#
# Для каждого файла solutions/<семинар>/<задача>/solution.go скрипт собирает
# во временном каталоге пакет из ЭТАЛОННОГО решения и ОТКРЫТОГО теста студента
# (<семинар>/tasks/<задача>/solution_test.go) и запускает go test.
#
# Так тест хранится ровно в одном месте и не может разойтись с решением.
#
# Использование: ./scripts/check_solutions.sh
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

total=0
failed=0

while IFS= read -r sol; do
	rel="${sol#"$ROOT"/solutions/}"       # seminar-01-intro/task01-greet/solution.go
	dir="$(dirname "$rel")"               # seminar-01-intro/task01-greet
	seminar="${dir%%/*}"                  # seminar-01-intro
	task="${dir#*/}"                      # task01-greet
	test_file="$ROOT/$seminar/tasks/$task/solution_test.go"

	total=$((total + 1))

	if [[ ! -f "$test_file" ]]; then
		echo "SKIP  $dir — нет открытого теста ($seminar/tasks/$task/solution_test.go)"
		continue
	fi

	work="$TMP/$seminar/$task"
	mkdir -p "$work"
	cp "$sol" "$work/solution.go"
	cp "$test_file" "$work/solution_test.go"
	cat > "$work/go.mod" <<EOF
module check/$task

go 1.24
EOF

	if (cd "$work" && go test ./... >"$work/out.txt" 2>&1); then
		echo "PASS  $dir"
	else
		echo "FAIL  $dir"
		sed 's/^/      /' "$work/out.txt"
		failed=$((failed + 1))
	fi
done < <(find "$ROOT/solutions" -name 'solution.go' | sort)

echo
echo "Итого: $total задач(и), провалено: $failed"
[[ "$failed" -eq 0 ]]
