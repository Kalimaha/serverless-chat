package main

import "errors"

type Result struct {
	value float32
	err   error
}

func Divide(a float32, b float32) (result Result) {
	if b == 0 {
		return Result{err: errors.New("division by 0")}
	}
	return Result{value: a / b}
}
