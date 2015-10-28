package command_and_queue

import "sync"

type Queue struct {
	sync.Mutex
	push   chan interface{}
	buffer []interface{}
}

func NewQueue() *Queue {
	q := &Queue{
		push: make(chan interface{}, 1),
	}

	go q.run()
	return q
}

func (q *Queue) Push(data interface{}) {
	q.push <- data
}

func (q *Queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	if len(q.buffer) > 0 {
		var data interface{}
		data, q.buffer = q.buffer[0], q.buffer[1:]
		return data
	} else {
		return nil
	}
}

func (q *Queue) run() {
	for data := range q.push {
		q.buffer = append(q.buffer, data)
	}
}
