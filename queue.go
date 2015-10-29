package gommons

import "sync"

type Event struct {
	data int
}

type Queue struct {
	queue []*Event
	sync.Mutex
}

func (q Queue) Len() int           { return len(q.queue) }
func (q Queue) Less(i, j int) bool { return true }
func (q Queue) Swap(i, j int)      {}

func (q *Queue) Push(x interface{}) {
	q.Lock()
	defer q.Unlock()

	item := x.(*Event)
	q.queue = append(q.queue, item)
}

func (q *Queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	old := q
	n := len(old.queue)

	if len(old.queue) < 1 {
		return nil
	}

	item := old.queue[n-1]
	q.queue = old.queue[0 : n-1]
	return item
}
