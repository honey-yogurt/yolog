package yolog

import (
	"bytes"
	"time"
)

// Entry
// 将日志输出到支持的输出中
type Entry struct {
	logger *logger
	Buffer *bytes.Buffer
	Map    map[string]interface{}
	Level  Level
	Time   time.Time
	File   string
	Line   int
	Func   string
	Format string
	Args   []interface{}
}
