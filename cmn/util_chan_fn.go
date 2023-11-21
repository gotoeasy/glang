package cmn

type Func func(args ...any) any

type SyncExecutor struct {
	chanExec chan *funcModel // 控制顺序执行用通道
}

type funcModel struct {
	Fn     Func      // 函数
	Args   any       // 参数
	Result any       // 返回结果
	Chan   chan bool // 等待直到取出返回结果用通道
}

// 创建线程安全的执行器
func NewSyncExecutor() *SyncExecutor {
	se := &SyncExecutor{
		chanExec: make(chan *funcModel, 32),
	}
	go func() {
		for {
			m := <-se.chanExec
			m.Result = m.Fn(m.Args.([]any)...)
			m.Chan <- true
		}
	}()
	return se
}

// 传入指定函数及其所需参数，返回该函数的执行结果，线程安全
func (s *SyncExecutor) Exec(fn Func, args ...any) any {
	m := &funcModel{
		Fn:   fn,
		Args: args,
		Chan: make(chan bool, 1),
	}
	s.chanExec <- m // 控制逐个执行，达到线程安全效果

	<-m.Chan // 等待返回结果
	return m.Result
}
