package argument

import (
	"fmt"
	"reflect"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type ArgumentDefinition map[string]any

type NoKeyInAgumentDefinitionError struct {
	Key string
}

func (t *NoKeyInAgumentDefinitionError) Error() string {
	return fmt.Sprintf("Key '%s' is not found in AgumentDefinition", t.Key)
}

type CannotCastValueInAgumentDefinitionError struct {
	Key string
	V   any
}

func (t *CannotCastValueInAgumentDefinitionError) Error() string {
	return fmt.Sprintf("Value of '%s' cannot be cast into type '%s' in AgumentDefinition", t.Key, reflect.TypeOf(t.V).Name())
}

func GetFromArgumentDefinition[V any](m ArgumentDefinition, key string) (V, error) {
	var zero V
	v, exists := m[key]
	if !exists {
		return zero, terrors.Wrap(&NoKeyInAgumentDefinitionError{Key: key})
	}
	vv, ok := v.(V)
	if !ok {
		return zero, terrors.Wrap(&CannotCastValueInAgumentDefinitionError{Key: key, V: v})
	}
	return vv, nil
}

func DefaultGetFromArgumentDefinition[V any](m ArgumentDefinition, key string, dflt V) V {
	v, err := GetFromArgumentDefinition[V](m, key)
	if err != nil {
		return dflt
	}
	return v
}
