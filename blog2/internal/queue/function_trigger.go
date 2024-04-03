package queue

import "context"

type FunctionTrigger interface {
	Publish(
		ctx context.Context,
		msg any,
	) error
}
