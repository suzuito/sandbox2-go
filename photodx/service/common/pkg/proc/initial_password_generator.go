package proc

import (
	"github.com/sethvargo/go-password/password"
)

type InitialPasswordGeneratorImpl struct {
}

func (t *InitialPasswordGeneratorImpl) Gen() (string, error) {
	return password.Generate(10, 5, 0, false, false)
}
