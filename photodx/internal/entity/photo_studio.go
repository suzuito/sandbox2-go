package entity

type PhotoStudioID string
type PhotoStudio struct {
	ID     PhotoStudioID `json:"id"`
	Name   string        `json:"name"`
	Active bool          `json:"active"`
}

func (*PhotoStudio) Validate() error {
	// TODO validate name
	// TODO validate id
	return nil
}
