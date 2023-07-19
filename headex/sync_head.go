package headex

import (
	"container/heap"
	"sync"
)

type SyncHeap struct {
	LockRW sync.RWMutex
	heap   heap.Interface
}

func (h *SyncHeap) Len() int {
	h.LockRW.RLock()
	defer h.LockRW.RUnlock()
	return h.heap.Len()
}

func (h *SyncHeap) Less(i, j int) bool {
	h.LockRW.RLock()
	defer h.LockRW.RUnlock()
	return h.heap.Less(i, j)
}

func (h *SyncHeap) Swap(i, j int) {
	h.LockRW.Lock()
	defer h.LockRW.Unlock()
	h.heap.Swap(i, j)
}

func (h *SyncHeap) Push(x interface{}) {
	h.LockRW.Lock()
	defer h.LockRW.Unlock()
	h.heap.Push(x)
}

func (h *SyncHeap) Pop() interface{} {
	h.LockRW.Lock()
	defer h.LockRW.Unlock()
	return h.heap.Pop()
}

func Init(hp heap.Interface) *SyncHeap {
	sh := &SyncHeap{heap: hp}
	heap.Init(sh.heap)
	return sh
}
