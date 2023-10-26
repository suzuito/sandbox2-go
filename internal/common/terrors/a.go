package terrors

import "fmt"

type TraceInfo struct {
	Filename string
	Line     int
}

type TraceableError interface {
	Error() string
	Unwrap() error
	StackTrace() []TraceInfo
}

func Wrap(err error) TraceableError {
	if err == nil {
		return nil
	}
	return newTraceableErrorImpl(2, err)
}

func Wrapf(format string, a ...any) TraceableError {
	return newTraceableErrorImpl(2, fmt.Errorf(format, a...))
}
