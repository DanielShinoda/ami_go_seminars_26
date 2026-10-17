package assert

import "reflect"

// Equal проверяет, что got равно want (глубокое сравнение).
//
// При несовпадении вызывает t.Errorf с сообщением, содержащим оба
// значения. При совпадении не делает ничего. В любом случае сначала
// вызывает t.Helper(), чтобы ошибка указывала на строку вызова Equal.
func Equal(t TestingT, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("получили %v, ожидалось %v", got, want)
	}
}

// True проверяет, что cond истинно.
//
// При ложном cond вызывает t.Errorf. При истинном не делает ничего.
// В любом случае сначала вызывает t.Helper().
func True(t TestingT, cond bool) {
	t.Helper()
	if !cond {
		t.Errorf("условие ложно, ожидалось true")
	}
}

// NoError проверяет, что err равен nil.
//
// При не-nil ошибке вызывает t.Errorf с сообщением, содержащим текст
// ошибки. При nil не делает ничего. В любом случае сначала вызывает
// t.Helper().
func NoError(t TestingT, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("неожиданная ошибка: %v", err)
	}
}
