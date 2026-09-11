.PHONY: help test test-01 test-02 fmt fmt-fix vet check solutions

help: ## Показать список команд
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

test: ## Прогнать все тесты (нерешённые задачи падают — это норма)
	go test ./...

test-01: ## Тесты задач семинара 1
	go test ./seminar-01-intro/tasks/...

test-02: ## Тесты задач семинара 2
	go test ./seminar-02-basics/tasks/...

fmt: ## Показать неотформатированные файлы
	@gofmt -l .

fmt-fix: ## Отформатировать все файлы
	gofmt -w .

vet: ## Статический анализ
	go vet ./...

check: fmt vet test ## Форматирование + анализ + тесты

solutions: ## Проверить эталонные решения преподавателя (папка solutions/)
	./scripts/check_solutions.sh
