package yolog

import (
	"bytes"
	"runtime"
	"strings"
	"time"
)

// Entry
// 将日志输出到支持的输出中
type Entry struct {
	logger *logger
	Buffer *bytes.Buffer
	Map    map[string]interface{} // json序列化format
	Level  Level                  // 写入的日志级别
	Time   time.Time              // 日志写入时间
	File   string                 // 反射拿到记录日志的文件名
	Line   int                    // 反射拿到记录日志的行号
	Func   string                 // 反射拿到记录日志的函数名
	Format string                 // 日志的格式化样式
	Args   []interface{}          // 写入日志的具体内容
}

func entry(logger *logger) *Entry {
	return &Entry{
		logger: logger,
		Buffer: new(bytes.Buffer),
		Map:    make(map[string]interface{}, 5),
	}
}

func (e *Entry) write(level Level, format string, args ...interface{}) {
	if e.logger.opt.level > level {
		// 低于指定的level 不进行写入
		return
	}
	e.Time = time.Now()
	e.Level = level
	e.Format = format
	if !e.logger.opt.disableCaller {
		if pc, file, line, ok := runtime.Caller(2); !ok {
			e.File = "???"
			e.Func = "???"
		} else {
			e.File, e.Line, e.Func = file, line, runtime.FuncForPC(pc).Name()
			e.Func = e.Func[strings.LastIndex(e.Func, "/")+1:]
		}
	}
	e.format()
	e.writer()
	e.release()
}

func (e *Entry) format() {
	_ = e.logger.opt.formatter.Format(e)
}

func (e *Entry) writer() {
	e.logger.mu.Lock()
	_, _ = e.logger.opt.output.Write(e.Buffer.Bytes())
	e.logger.mu.Unlock()
}

func (e *Entry) release() {
	e.Args, e.Line, e.File, e.Format, e.Func = nil, 0, "", "", ""
	e.Buffer.Reset()
	e.logger.entryPool.Put(e)
}
