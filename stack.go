package main

import (
	"errors"
)

type Stack struct {
	data []int
}

func (myStr *Stack) Push(item int) {
	myStr.data = append(myStr.data, item)
}

func (myStr *Stack) Pop() (int, error) {
	if len(myStr.data) == 0 {
		return 0, errors.New("Stack is empty")
	}

	myStr.data = myStr.data[:len(myStr.data)-1] //Срез через data int
}

func (myStr *Stack) IsEmpty() bool {
	return len(myStr.data) == 0
}
