package view

import "context"

type View interface {
	Run(ctx context.Context) error
}
