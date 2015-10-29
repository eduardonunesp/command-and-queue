package gommons

import (
	"container/heap"
	"runtime"
	"strconv"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go test -goblin.timeout=15s

func TestQueue(t *testing.T) {
	g := Goblin(t)
	g.Describe("Basic Queue test", func() {
		g.It("Should push and pop data ", func() {
			queue := &Queue{}
			heap.Init(queue)
			heap.Push(queue, &Event{data: 1})
			heap.Push(queue, &Event{data: 2})
			heap.Push(queue, &Event{data: 3})

			data := heap.Pop(queue).(*Event)
			g.Assert(data.data).Equal(1)

			data = heap.Pop(queue).(*Event)
			g.Assert(data.data).Equal(2)

			data = heap.Pop(queue).(*Event)
			g.Assert(data.data).Equal(3)

			d := heap.Pop(queue)
			g.Assert(d == nil).IsTrue()
		})
	})
}

func TestQueueConcurrency(t *testing.T) {
	g := Goblin(t)
	g.Describe("Concurrency test", func() {
		g.It("Should test concurrency with 1 CPU", func() {
			runtime.GOMAXPROCS(1)
			queue := &Queue{}
			heap.Init(queue)

			buffer := []int{}

			for i := 0; i < 1000; i++ {
				go func(v int) {
					heap.Push(queue, &Event{data: v})
				}(i)
			}

			for {
				go func() {
					data := heap.Pop(queue)
					if data != nil {
						buffer = append(buffer, data.(*Event).data)
					}
				}()

				time.Sleep(time.Millisecond)
				if len(buffer) >= 1000 {
					break
				}
			}

			g.Assert(len(buffer) >= 1000).IsTrue()
		})

		strCpus := strconv.Itoa(runtime.NumCPU())

		g.It("Should test concurrency with more than 1 CPU ("+strCpus+")", func() {
			cpus := runtime.NumCPU()
			if cpus == 1 {
				cpus = 2
			}

			runtime.GOMAXPROCS(cpus)
			queue := &Queue{}
			heap.Init(queue)

			buffer := []int{}

			for i := 0; i < 1000; i++ {
				go func(v int) {
					heap.Push(queue, &Event{data: v})
				}(i)
			}

			for {
				go func() {
					data := heap.Pop(queue)
					if data != nil {
						buffer = append(buffer, data.(*Event).data)
					}
				}()

				time.Sleep(time.Millisecond)
				if len(buffer) >= 1000 {
					break
				}
			}

			g.Assert(len(buffer) >= 1000).IsTrue()
		})
	})
}
