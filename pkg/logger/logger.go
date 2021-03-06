package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
)

type Level int8
type Fields map[string]interface{}
const(
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

type Logger struct{
	newLogger *log.Logger
	ctx context.Context
	fileds Fields
	callers []string
}


func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

func NewLogger(w io.Writer,prefix string,flag int)*Logger{
	l := log.New(w,prefix,flag)
	return &Logger{newLogger:l}
}
func (l *Logger)clone()*Logger{
	nl := *l
	return &nl
}
func (l *Logger)WithFields(f Fields)*Logger{
	ll := l.clone()
	if ll.fileds == nil {
		ll.fileds = make(Fields)
	}
	for k,v := range  f{
		ll.fileds[k] = v
	}
	return ll
}
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}
func (l *Logger)WtihCaller(skip int)*Logger{
	ll := l.clone()
	pc,file,line,ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s",file,line,f.Name())}
	}
	return ll
}
func (l *Logger)WithCallersFrames() *Logger{
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr,maxCallerDepth)
	depth := runtime.Callers(minCallerDepth,pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next();more; frame,more = frames.Next(){
		callers = append(callers,fmt.Sprint("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}