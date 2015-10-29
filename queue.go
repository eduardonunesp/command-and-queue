package gommons

import "sync"

type node struct {
	data interface{}
	next *node
}

//  A go-routine safe FIFO (first in first out) data stucture.
type Queue struct {
	head  *node
	tail  *node
	count int
	sync.RWMutex
}

//  Creates a new pointer to a new queue.
func NewQueue() *Queue {
	return &Queue{}
}

//  Returns the number of elements in the queue (i.e. size/length)
//  go-routine safe.
func (q *Queue) Len() int {
	q.RLock()
	defer q.RUnlock()
	return q.count
}

//  Pushes/inserts a value at the end/tail of the queue.
//  Note: this function does mutate the queue.
//  go-routine safe.
func (q *Queue) Push(item interface{}) {
	q.Lock()
	defer q.Unlock()

	n := &node{data: item}

	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = n
	}

	q.count++
}

//  Returns the value at the front of the queue.
//  i.e. the oldest value in the queue.
//  Note: this function does mutate the queue.
//  go-routine safe.
func (q *Queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	if q.head == nil {
		return nil
	}

	n := q.head
	q.head = n.next

	if q.head == nil {
		q.tail = nil
	}

	q.count--
	return n.data
}

//  Returns a read value at the front of the queue.
//  i.e. the oldest value in the queue.
//  Note: this function does NOT mutate the queue.
//  go-routine safe.
func (q *Queue) Peek() interface{} {
	q.Lock()
	defer q.Unlock()

	n := q.head
	if n == nil || n.data == nil {
		return nil
	}

	return n.data
}
