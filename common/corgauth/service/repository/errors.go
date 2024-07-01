package repository

import (
	"fmt"
)

type EntityType string

const (
	EntityTypeOrganization EntityType = "Organization"
	EntityTypePrincipal    EntityType = "Principal"
)

type EntityNotFoundError struct {
	EntityType EntityType
	ID         string
}

func (t *EntityNotFoundError) Error() string {
	return fmt.Sprintf("not found entity %s.%s", t.EntityType, t.ID)
}

type EntityAlreadyExistsError struct {
	EntityType EntityType
	ID         string
}

func (t *EntityAlreadyExistsError) Error() string {
	return fmt.Sprintf("already exists entity %s.%s", t.EntityType, t.ID)
}
