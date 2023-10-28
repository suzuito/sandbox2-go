package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryError(t *testing.T) {
	var err error
	var rerr *RepositoryError

	err = &RepositoryError{
		EntityURL: "foo",
		Message:   "bar",
		Code:      RepositoryErrorCodeNotFound,
	}
	assert.Equal(t, "Not found foo", err.Error())
	assert.True(t, errors.As(err, &rerr))

	err = &RepositoryError{
		EntityURL: "foo",
		Message:   "bar",
		Code:      RepositoryErrorCode("99"),
	}
	assert.Equal(t, "Error on 'foo' bar", err.Error())
	assert.True(t, errors.As(err, &rerr))
}
