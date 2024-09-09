package queue

import (
	"fmt"
)

type Queue struct {
	data []interface{}
}

func (q *Queue) Enqueue(value interface{}) {
	q.data = append(q.data, value)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("queue is empty")
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value, nil
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("queue is empty")
	}
	return q.data[0], nil
}

func (q *Queue) Clear() {
	q.data = make([]interface{}, 0)
}
