package cmn

import "sync"

// 队列
type Queue struct {
	items []any // 队列元素
	mu    sync.Mutex
}

// 新建队列（线程安全）
func NewQueue() *Queue {
	return &Queue{
		items: make([]any, 0),
	}
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items) == 0
}

// 添加元素
func (q *Queue) Push(item any) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// 取出元素
func (q *Queue) Pop() any {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) != 0 {
		item := q.items[0]
		q.items = q.items[1:]
		return item
	}
	return nil
}

// 取元素，不出队
func (q *Queue) Peek() any {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

// 当前数据的切片副本
func (q *Queue) Copied() []any {
	q.mu.Lock()
	defer q.mu.Unlock()
	copiedSlice := make([]any, len(q.items))
	copy(copiedSlice, q.items)
	return copiedSlice
}
