package proc

import "github.com/google/uuid"

type IDGeneratorImpl struct {
}

func (t *IDGeneratorImpl) Gen() (string, error) {
	uid := uuid.Must(uuid.NewV7())
	return uid.String(), nil
}
