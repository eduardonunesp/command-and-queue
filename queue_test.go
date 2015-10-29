package gommons

import (
	"runtime"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

// go test -goblin.timeout=15s

func TestQueue(t *testing.T) {
	g := Goblin(t)
	g.Describe("Basic Queue test", func() {
		g.It("Should create queue ", func() {
			q := NewQueue()
			g.Assert(q != nil).IsTrue()
		})

		g.It("Should push data ", func() {
			q := NewQueue()
			q.Push(10)
		})

		g.It("Should pop data ", func() {
			q := NewQueue()
			q.Push(10)
			time.Sleep(time.Millisecond * 100)
			data := q.Pop()
			g.Assert(data != nil).IsTrue()
		})

		g.It("Should pop many data ", func() {
			q := NewQueue()
			q.Push(10)
			q.Push(11)
			q.Push(12)

			time.Sleep(time.Millisecond * 100)

			data := q.Pop()
			g.Assert(data == 10).IsTrue()

			data = q.Pop()
			g.Assert(data == 11).IsTrue()

			data = q.Pop()
			g.Assert(data == 12).IsTrue()
		})

	})
}

func TestQueueConcurrency(t *testing.T) {
	g := Goblin(t)
	g.Describe("Concurrency test", func() {
		g.It("Should test concurrency with 1 CPU", func() {
			runtime.GOMAXPROCS(1)
			q := NewQueue()
			g.Assert(q != nil).IsTrue()

			buffer := []int{}

			for i := 0; i < 1000; i++ {
				go func() {
					q.Push(i)
				}()
			}

			for {
				go func() {
					data := q.Pop()
					if data != nil {
						buffer = append(buffer, data.(int))
					}
				}()

				time.Sleep(time.Millisecond)
				if len(buffer) >= 1000 {
					break
				}
			}

			g.Assert(len(buffer) >= 1000).IsTrue()
		})

		g.It("Should test concurrency with 2 >= CPUs", func() {
			cpus := runtime.NumCPU()
			if cpus == 1 {
				cpus = 2
			}

			runtime.GOMAXPROCS(cpus)
			q := NewQueue()
			g.Assert(q != nil).IsTrue()

			buffer := []int{}

			for i := 0; i < 1000; i++ {
				go func() {
					q.Push(i)
				}()
			}

			for {
				go func() {
					data := q.Pop()
					if data != nil {
						buffer = append(buffer, data.(int))
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
