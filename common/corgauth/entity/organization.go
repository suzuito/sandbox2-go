package entity

type OrganizationID string

type Organization struct {
	ID     OrganizationID
	Active bool
}

func (t *Organization) Validate() error {
	return nil
}
