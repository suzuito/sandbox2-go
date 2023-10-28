package terrors

import (
	"errors"
	"runtime"
)

func newTraceableErrorImpl(runtimeCallerSkip int, orig error) *traceableErrorImpl {
	err := traceableErrorImpl{
		OriginalErr: orig,
	}
	_, filename, line, ok := runtime.Caller(runtimeCallerSkip)
	if ok {
		err.TraceInfo.Filename = filename
		err.TraceInfo.Line = line
	}
	return &err
}

type traceableErrorImpl struct {
	TraceInfo   TraceInfo
	OriginalErr error
}

func (t *traceableErrorImpl) Error() string {
	if t.OriginalErr != nil {
		return t.OriginalErr.Error()
	}
	return ""
}

func (t *traceableErrorImpl) Unwrap() error {
	return t.OriginalErr
}

func (t *traceableErrorImpl) StackTrace() []TraceInfo {
	traceInfos := []TraceInfo{
		t.TraceInfo,
	}
	var err error = t
	for {
		err = errors.Unwrap(err)
		if err == nil {
			break
		}
		terr, ok := err.(*traceableErrorImpl)
		if ok {
			traceInfos = append(traceInfos, terr.TraceInfo)
		}
	}
	return traceInfos
}
