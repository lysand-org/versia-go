package utils

import (
	"context"
	"encoding/json"

	"github.com/versia-pub/versia-go/pkg/taskqueue"
)

func ParseTask[T any](handler func(context.Context, T) error) func(context.Context, taskqueue.Task) error {
	return func(ctx context.Context, task taskqueue.Task) error {
		var data T
		if err := json.Unmarshal(task.Payload, &data); err != nil {
			return err
		}

		return handler(ctx, data)
	}
}
