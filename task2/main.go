package main

import (
	"fmt"
	"time"
)

type myWrappError struct {
	text      string
	timeEvent time.Time
}

func NewMyWrappError(text string, time time.Time) error {
	return &myWrappError{text, time}
}

func (e *myWrappError) Error() string {
	return fmt.Sprintf("Ошибка: %s, получена в %s", e.text, e.timeEvent)
}

// Пустая ошибка для проброса между функцией, получающая панику и main
var myErr error

func main() {
	err := funcWrappPanic()

	if err != nil {
		fmt.Println(err)
	}
}

func funcWrappPanic() error {
	funcWithPanic()

	if myErr != nil {
		return myErr
	}

	return nil
}

func funcWithPanic() {
	defer func() {
		if v := recover(); v != nil {
			myErr = NewMyWrappError(fmt.Sprintf("Перехвачена ошибка (%s)", v), time.Now())
		}
	}()

	var a int = 0
	_ = 1 / a
}
