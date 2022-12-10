package cmn

// 事件总线结构体
type EventBus struct {
	mapHandle map[string]([]EventHandler)
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
	name := ToLower(Trim(event))
	handles := e.mapHandle[name]
	if handles == nil {
		handles = []EventHandler{}
	}
	handles = append(handles, handle)
	e.mapHandle[name] = handles
	return e
}

// 触发事件
func (e *EventBus) At(event string, params ...any) *EventBus {
	name := ToLower(Trim(event))
	handles := e.mapHandle[name]
	if handles == nil || len(handles) < 1 {
		return e
	}

	for i := 0; i < len(handles); i++ {
		defer func() {
			if err := recover(); err != nil {
				Error("事件执行发生异常，事件名：", name, "，参数：", params, "，异常：", err)
			}
		}()
		handles[i](params...)
	}
	return e
}
