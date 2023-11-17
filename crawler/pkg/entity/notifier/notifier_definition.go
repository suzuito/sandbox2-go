package notifier

import "github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"

type NotifierDefinition struct {
	ID                       NotifierID
	DocPathFirestoreMatchers []string
	Argument                 argument.ArgumentDefinition
}
