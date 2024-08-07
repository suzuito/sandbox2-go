package proc

import (
	"math/rand"
	"strconv"
)

type UserCreationCodeGeneratorImpl struct {
}

func (t *UserCreationCodeGeneratorImpl) Gen() (string, error) {
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code, nil
}
