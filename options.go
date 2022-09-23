package yolog

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	FmtEmptySeparate = ""
)

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

// Level
// 日志级别
type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var LevelNameMapping = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug", "DEBUG":
		*l = DebugLevel // todo 这里为什么要 *
	case "info", "INFO":
		*l = InfoLevel
	case "warn", "WARN":
		*l = WarnLevel
	case "error", "ERROR":
		*l = ErrorLevel
	case "panic", "PANIC":
		*l = PanicLevel
	case "fatal", "Fatal":
		*l = FatalLevel
	default:
		return false
	}
	return true
}

func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

// 日志选项，可配置项
type options struct {
	output        io.Writer // 输出位置 标准输出或者文件输出
	level         Level     // 输出级别
	stdLevel      Level     //
	formatter     Formatter // 设置输出级别 需要一个Formatter接口，先建一个空接口占位
	disableCaller bool      // 设置是否打印文件名和行号
}

type Option func(*options)

// 用 Option 来 初始化 option
func initOptions(opts ...Option) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o) // 这里传的指针，所以 func的执行变更会带出来
	}

	if o.output == nil {
		o.output = os.Stderr
	}
	if o.formatter == nil {
		o.formatter = &TextFormatter{} //设置默认输出格式，普通文本模式
	}
	return
}

func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

// WithLevel
// 返回带level的Option
func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithStdLevel(level Level) Option {
	return func(o *options) {
		o.stdLevel = level
	}
}

func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		o.formatter = formatter
	}
}

func WithDisableCaller(caller bool) Option {
	return func(o *options) {
		o.disableCaller = caller
	}
}
