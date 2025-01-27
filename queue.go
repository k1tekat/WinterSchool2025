package main

import "errors"

type Queue struct {
	data       []int
	head, tail int
}

func (myQue *Queue) Enqueue(x int) {
	myQue.data = append(myQue.data, x)
}

func (myQue *Queue) Dequeue() (int, error) {
	if len(myQue.data) == 0 {
		return 0, errors.New("Queue is empty")
	}

	item := myQue.data[0]
	myQue.data = myQue.data[1:]

	return item, nil
}
