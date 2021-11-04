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

func main() {
	err := funcWithPanic()
	if err != nil {
		fmt.Println(err)
	}
}

func funcWithPanic() (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = NewMyWrappError(fmt.Sprintf("Перехвачена ошибка (%s)", v), time.Now())
		}
	}()

	var a int = 0
	_ = 1 / a

	return nil
}
