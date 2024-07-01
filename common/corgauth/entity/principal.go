package entity

type PrincipalID string

type Principal struct {
	ID     PrincipalID `json:"id"`
	Email  string      `json:"email"`
	Active bool        `json:"active"`
}

func (t *Principal) Validate() error {
	return nil
}
