package terrors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var exact error
	exact = Wrap(nil)
	assert.Nil(t, exact)
	exact = Wrap(fmt.Errorf("dummy error"))
	assert.Regexp(t, "^dummy error", exact.Error())
	assert.Regexp(t, ".+/impl_test.go", exact.(*traceableErrorImpl).TraceInfo.Filename)
	assert.Equal(t, 14, exact.(*traceableErrorImpl).TraceInfo.Line)
}

func TestStackTrace(t *testing.T) {
	err1 := fmt.Errorf("err1")
	err2 := Wrap(err1)
	err3 := fmt.Errorf("%w", err2)
	err4 := Wrap(err3)
	traceInfos := err2.(*traceableErrorImpl).StackTrace()
	assert.Equal(t, 1, len(traceInfos))
	assert.Regexp(t, ".+/impl_test.go", traceInfos[0].Filename)
	assert.Equal(t, 22, traceInfos[0].Line)
	traceInfos = err4.(*traceableErrorImpl).StackTrace()
	assert.Equal(t, 2, len(traceInfos))
	assert.Regexp(t, ".+/impl_test.go", traceInfos[0].Filename)
	assert.Equal(t, 24, traceInfos[0].Line)
	assert.Regexp(t, ".+/impl_test.go", traceInfos[1].Filename)
	assert.Equal(t, 22, traceInfos[1].Line)
}
