package yolog

import "sync"

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

func New(opts ...Option) *logger {
	// Option 定制化 options，然后填充到 logger
	logger := &logger{opt: initOptions(opts...)} //这里为什么需要三个点 本来不就是切片类型吗？
	//	todo 这里的entry暂时不处理，需要再处理
	return logger

}
