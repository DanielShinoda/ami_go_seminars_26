# Команды для работы с материалами курса. `make help` — список.
#
# Коротко:
#   make doctor  — «у меня всё установлено?»
#   make test    — «мои задачи решены?»   (нерешённые падают — это норма)
#   make check   — «я ничего не сломал?»  (должно быть зелёным всегда)
#
# Список семинаров нигде не записан: новая папка seminar-NN-тема
# подхватывается сама, править этот файл не нужно.

.DEFAULT_GOAL := help
.PHONY: help doctor test fmt lint check

SEMINARS := $(sort $(wildcard seminar-*))
EXAMPLES := $(patsubst %,./%/examples/...,$(SEMINARS))
TASKS    := $(patsubst %,./%/tasks/...,$(SEMINARS))

help: ## Показать список команд
	@grep -E '^[a-zA-Z0-9_%-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-9s\033[0m %s\n", $$1, $$2}'

doctor: ## Проверить, что окружение настроено
	@go version
	@go test $(EXAMPLES) > /dev/null && echo "Окружение готово."

test: ## Тесты всех задач (нерешённые падают — это норма)
	go test $(TASKS)

test-%: ## Тесты задач одного семинара: make test-01
	@d=$$(ls -d seminar-$*-* 2>/dev/null | head -1); \
		test -n "$$d" || { echo "Нет семинара с номером $*"; exit 1; }; \
		go test ./$$d/tasks/...

task: ## Одна задача подробно: make task T=seminar-02-basics/tasks/task02-reverse
	@test -n "$(T)" || { echo "Укажите задачу: make task T=<путь>"; exit 1; }
	go test -v ./$(T)

fmt: ## Отформатировать код
	gofmt -w .

lint: ## Форматирование и go vet — перед сдачей должно быть пусто
	@out=$$(gofmt -l .); test -z "$$out" || { echo "Не отформатировано:"; echo "$$out"; exit 1; }
	go vet ./...

check: lint ## Материалы целы: форматирование, vet, примеры компилируются и проходят тесты
	go build ./...
	go test $(EXAMPLES)
