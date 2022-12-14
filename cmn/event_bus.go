package cmn

import (
	"reflect"
	"sync"
)

// 事件总线结构体
type EventBus struct {
	mapHandle map[string]([]EventHandler)
	mu        sync.Mutex // 锁
}

// 事件处理器
type EventHandler func(params ...any)

var _eventBus *EventBus

// 创建事件总线（单例）
func NewEventBus() *EventBus {
	if _eventBus == nil {
		_eventBus = &EventBus{
			mapHandle: make(map[string][]EventHandler),
		}
	}
	return _eventBus
}

// 注册事件
func (e *EventBus) On(event string, handle EventHandler) *EventBus {
	e.mu.Lock()
	defer e.mu.Unlock()

	name := ToLower(Trim(event))
	handles := e.mapHandle[name]
	if handles == nil {
		handles = []EventHandler{}
	}

	f := false
	for i := 0; i < len(handles); i++ {
		v := reflect.ValueOf(handle)
		if reflect.ValueOf(handles[i]) == v {
			f = true
			break
		}
	}

	if !f {
		handles = append(handles, handle)
		e.mapHandle[name] = handles
	}

	return e
}

// 注销事件
func (e *EventBus) Off(event string, delHandles ...EventHandler) *EventBus {
	e.mu.Lock()
	defer e.mu.Unlock()

	if delHandles == nil || len(delHandles) < 1 {
		return e
	}

	name := ToLower(Trim(event))
	handles := e.mapHandle[name]
	if handles == nil || len(handles) < 1 {
		return e
	}

	newHandles := []EventHandler{}
	for i := 0; i < len(handles); i++ {
		f := false
		v := reflect.ValueOf(handles[i])
		for j := 0; j < len(delHandles); j++ {
			if v == reflect.ValueOf(delHandles[j]) {
				f = true
				break
			}
		}
		if !f {
			newHandles = append(newHandles, handles[i])
		}
	}

	e.mapHandle[name] = newHandles
	return e
}

// 注销事件
func (e *EventBus) Del(event string) *EventBus {
	e.mu.Lock()
	defer e.mu.Unlock()

	name := ToLower(Trim(event))
	handles := e.mapHandle[name]
	if handles == nil || len(handles) < 1 {
		return e
	}

	e.mapHandle[name] = []EventHandler{}
	return e
}

// 重置
func (e *EventBus) Reset() *EventBus {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.mapHandle = make(map[string][]EventHandler)
	return e
}

// 触发事件
func (e *EventBus) At(event string, params ...any) *EventBus {
	e.mu.Lock()
	defer e.mu.Unlock()

	name := ToLower(Trim(event))
	handles := e.mapHandle[name]
	if handles == nil || len(handles) < 1 {
		return e
	}

	for i := 0; i < len(handles); i++ {
		go execEvent(name, handles[i], params...) // 异步执行，出错时打印错误，不中断循环
	}

	return e
}

func execEvent(name string, handle EventHandler, params ...any) {
	defer func() {
		if err := recover(); err != nil {
			Error("事件执行发生异常，事件名：", name, "，参数：", params, "，异常：", err)
		}
	}()
	handle(params...)
}
