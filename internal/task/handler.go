package task

import (
	"context"

	"github.com/versia-pub/versia-go/pkg/taskqueue"
)

type Manager interface {
	Notes() NoteHandler
}

type Handler interface {
	Register(*taskqueue.Set)
	Submit(context.Context, taskqueue.Task) error
}

type NoteHandler interface {
	Submit(context.Context, taskqueue.Task) error
}
