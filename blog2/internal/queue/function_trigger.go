package queue

import "context"

type FunctionTrigger interface {
	Put(
		ctx context.Context,
		msg any,
	) error
}
