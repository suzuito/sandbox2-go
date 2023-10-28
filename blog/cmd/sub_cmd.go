package cmd

import "context"

type SubCmd interface {
	Run(
		ctx context.Context,
		args []string,
	) error
}
