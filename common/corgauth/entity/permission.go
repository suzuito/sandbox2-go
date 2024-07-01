package entity

type Permission struct {
	Resource Resource
	Action   Action
}

func (t *Permission) Validate() error {
	return nil
}
