package test_helper

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertError(t *testing.T, expectedErrorMessageRegExp string, actual error) {
	if expectedErrorMessageRegExp != "" {
		if actual == nil {
			assert.Failf(
				t,
				"assertError is failed",
				`expected "%s" but actual nil`,
				expectedErrorMessageRegExp,
			)
			return
		}
		assert.Regexp(t, expectedErrorMessageRegExp, actual.Error())
		return
	}
	if actual != nil {
		assert.Failf(
			t,
			"assertError is failed",
			`expected nil but actual "%s"`,
			actual.Error(),
		)
	}
}

func AssertErrorAs(t *testing.T, expectedAs any, actual error) {
	if actual == nil {
		return
	}
	typ := reflect.TypeOf(expectedAs)
	assert.Truef(
		t,
		errors.As(actual, &expectedAs),
		"Expected error is %s but not",
		typ.Name(),
	)
}
