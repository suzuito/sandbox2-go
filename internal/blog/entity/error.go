package entity

type ValidationError struct {
	err error
}

func (t *ValidationError) Error() string {
	return t.err.Error()
}
