package gommons

import "sync"

type Event struct {
	index int
	data  int
}

type Queue struct {
	queue []*Event
	sync.RWMutex
}

func (q Queue) Len() int {
	q.RLock()
	defer q.RUnlock()

	return len(q.queue)
}

func (q Queue) Less(i, j int) bool {
	q.RLock()
	defer q.RUnlock()

	return q.queue[i].index < q.queue[j].index
}

func (q Queue) Swap(i, j int) {
	q.Lock()
	defer q.Unlock()

	if len(q.queue) == 0 {
		return
	}

	q.queue[i], q.queue[j] = q.queue[j], q.queue[i]
	q.queue[i].index = j
	q.queue[j].index = i
}

func (q *Queue) Push(x interface{}) {
	q.Lock()
	defer q.Unlock()

	n := len(q.queue)
	item := x.(*Event)
	item.index = n
	q.queue = append(q.queue, item)
}

func (q *Queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	old := q.queue
	n := len(old)

	if n == 0 {
		return nil
	}

	item := old[n-1]
	item.index = -1
	q.queue = old[:n-1]
	return item
}
