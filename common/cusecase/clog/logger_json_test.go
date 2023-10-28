package clog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

var _ Logger = NewLoggerJSON()

func hasMap(m map[string]any, key string) bool {
	_, ok := m[key]
	return ok
}

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		execute func(ctx context.Context, logger *LoggerJSON)
		assert  func(t *testing.T, exact map[string]any)
	}{
		{
			desc: "Infof test",
			execute: func(ctx context.Context, logger *LoggerJSON) {
				logger.Infof(ctx, "hoge")
			},
			assert: func(t *testing.T, exact map[string]any) {
				assert.Equal(t, "INFO", exact["severity"])
				assert.Equal(t, "hoge", exact["message"])
			},
		},
		{
			desc: "Output context data into log. How to use AddKey method",
			execute: func(ctx context.Context, logger *LoggerJSON) {
				ctx = context.WithValue(ctx, "foo", "value of foo")
				ctx = context.WithValue(ctx, "bar", "value of bar")

				logger.AddKey("foo")
				logger.Infof(ctx, "hoge")
			},
			assert: func(t *testing.T, exact map[string]any) {
				assert.Equal(t, "INFO", exact["severity"])
				assert.Equal(t, "hoge", exact["message"])
				assert.Equal(t, "value of foo", exact["foo"])
				assert.False(t, hasMap(exact, "bar"))
			},
		},
		{
			desc: "If TraceableError is contained in arguments, output traceInfo into log.",
			execute: func(ctx context.Context, logger *LoggerJSON) {
				err1 := fmt.Errorf("err1")
				err2 := terrors.Wrap(err1)
				err3 := fmt.Errorf("%w", err2)
				err4 := terrors.Wrap(err3)
				logger.Infof(ctx, "hoge %+v", err4)
			},
			assert: func(t *testing.T, exact map[string]any) {
				assert.Equal(t, "INFO", exact["severity"])
				assert.Equal(t, "hoge err1", exact["message"])
				traceInfos := exact["traceInfos"].([]any)
				assert.Equal(t, 2, len(traceInfos))
				traceInfo := traceInfos[0].(map[string]any)
				assert.True(t, hasMap(traceInfo, "Filename"))
				assert.True(t, hasMap(traceInfo, "Line"))
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctx := context.Background()
			logger := NewLoggerJSON()
			dummyWriter := bytes.NewBufferString("")
			logger.log.SetOutput(dummyWriter)
			tC.execute(ctx, logger)
			exact := map[string]any{}
			err := json.Unmarshal(dummyWriter.Bytes(), &exact)
			if err != nil {
				t.Errorf("%+v", err)
				t.Fail()
			}
			assert.True(t, hasMap(exact, "file"))
			assert.True(t, hasMap(exact, "line"))
			tC.assert(t, exact)
		})
	}
}

/*
func Test(t *testing.T) {
	ctx := context.Background()
	logger := NewLoggerJSON()
	dummyWriter := bytes.NewBufferString("")
	logger.log.SetOutput(dummyWriter)
	logger.Infof(ctx, "a")
	exact := map[string]any{}
	err := json.Unmarshal(dummyWriter.Bytes(), &exact)
	if err != nil {
		t.Errorf("%+v", err)
		t.Fail()
	}
	assert.Equal(t, "a", exact["line"])
}
*/
