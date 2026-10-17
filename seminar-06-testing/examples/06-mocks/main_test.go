// Моки через mockgen: тестируем код с внешними зависимостями.
//
//	go test -v ./seminar-06-testing/examples/06-mocks
package main

import (
	"errors"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

// Этот пример связывает семинары 4 и 5: код принимает интерфейс, тест
// подставляет сгенерированный мок — и получает полный контроль над
// «внешним миром»: что вызвали, сколько раз и с какими аргументами.

func TestRetrySucceedsOnThirdAttempt(t *testing.T) {
	// Контроллер следит за ожиданиями всех моков теста и в конце сам
	// проверяет, что ожидаемые вызовы случились (через t.Cleanup).
	ctrl := gomock.NewController(t)
	sleeper := NewMockSleeper(ctrl)

	// Ожидание: две паузы — 10 мс после первой неудачи, 20 мс после
	// второй. После УСПЕШНОЙ попытки ждать не должны — лишний вызов
	// провалит тест.
	gomock.InOrder(
		sleeper.EXPECT().Sleep(10*time.Millisecond),
		sleeper.EXPECT().Sleep(20*time.Millisecond),
	)

	failures := 2
	op := func() error {
		if failures > 0 {
			failures--
			return errors.New("не сейчас")
		}
		return nil
	}

	if err := Retry(5, sleeper, op); err != nil {
		t.Fatalf("Retry вернул ошибку %v, ожидалась nil", err)
	}
}

func TestRetryGivesUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	sleeper := NewMockSleeper(ctrl)

	// Три попытки — две паузы между ними. Порядок аргументов проверяем
	// через gomock.InOrder: сначала 10 мс, потом 20 мс.
	gomock.InOrder(
		sleeper.EXPECT().Sleep(10*time.Millisecond),
		sleeper.EXPECT().Sleep(20*time.Millisecond),
	)

	calls := 0
	op := func() error {
		calls++
		return errors.New("всегда падаю")
	}

	err := Retry(3, sleeper, op)
	if err == nil {
		t.Fatal("Retry вернул nil, ожидалась последняя ошибка")
	}
	if calls != 3 {
		t.Errorf("op вызвана %d раз, ожидалось 3", calls)
	}
}

// Мок умеет не только принимать вызовы, но и отвечать: DoAndReturn
// подменяет поведение. Здесь — имитация «часы идут», без реального сна.
func TestRetrySleepsAreVirtual(t *testing.T) {
	ctrl := gomock.NewController(t)
	sleeper := NewMockSleeper(ctrl)

	var virtual time.Duration
	sleeper.EXPECT().Sleep(gomock.Any()).AnyTimes().DoAndReturn(
		func(d time.Duration) { virtual += d },
	)

	op := func() error { return errors.New("падаю") }
	_ = Retry(4, sleeper, op)

	// 10 + 20 + 30 мс — паузы между четырьмя попытками.
	if want := 60 * time.Millisecond; virtual != want {
		t.Errorf("суммарная виртуальная пауза = %v, ожидалось %v", virtual, want)
	}
}

// Чем мок отличается от ручного фейка (семинар 4 давал интерфейсы,
// пример выше — ту же задачу):
//   - ожидания объявляются ДО вызова: какие методы, с какими аргументами,
//     сколько раз — и проверяются автоматически в конце теста;
//   - неожиданный вызов проваливает тест сразу, а не молча записывается;
//   - код мока генерируется, а не пишется руками: go generate ./...
//
// Ручной фейк всё ещё уместен, когда поведение сложнее, чем «вызвали —
// ответили»: накопление состояния, эмуляция протокола. Но стартовать
// стоит с mockgen.
