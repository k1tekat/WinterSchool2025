package queue

import "errors"

type Queue struct {
	Data []int
	//head, tail int
}

func (myQue *Queue) Enqueue(x int) {
	myQue.Data = append(myQue.Data, x)
}

func (myQue *Queue) Dequeue() (int, error) {
	if len(myQue.Data) == 0 {
		return 0, errors.New("Queue is empty")
	}

	item := myQue.Data[0]
	myQue.Data = myQue.Data[1:]

	return item, nil
}
