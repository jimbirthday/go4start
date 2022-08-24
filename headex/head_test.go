package headex

import (
	"fmt"
	"testing"
	"time"
)

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

//线程安全优先级队列
func TestHead(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	items := make(PriorityQueue, 3)
	i := 0
	for k, v := range m {
		items[i] = &Item{
			value:    k,
			priority: v,
			index:    0,
		}
		i++
	}
	syncHeap := Init(&items)
	for j := 0; j < 100; j++ {
		go func(k int) {
			syncHeap.Push(&Item{
				value:    fmt.Sprintf("%d%s", k, "some"),
				priority: k,
				index:    k,
			})
		}(j)
	}

	syncHeap.Swap(0, 1)

	time.Sleep(time.Second * 3)

	pop := syncHeap.Pop()
	item := pop.(*Item)
	fmt.Println(item.value)
	fmt.Println(item.priority)
	fmt.Println(item.index)
	fmt.Println("done")
}
