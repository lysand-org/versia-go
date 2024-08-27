package svc_impls

import (
	"context"
	"github.com/versia-pub/versia-go/internal/service"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/versia-pub/versia-go/pkg/taskqueue"
)

var _ service.TaskService = (*TaskServiceImpl)(nil)

type TaskServiceImpl struct {
	client *taskqueue.Client

	telemetry *unitel.Telemetry
	log       logr.Logger
}

func NewTaskServiceImpl(client *taskqueue.Client, telemetry *unitel.Telemetry, log logr.Logger) *TaskServiceImpl {
	return &TaskServiceImpl{
		client: client,

		telemetry: telemetry,
		log:       log,
	}
}

func (i TaskServiceImpl) ScheduleTask(ctx context.Context, type_ string, data any) error {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/TaskServiceImpl.ScheduleTask")
	defer s.End()
	ctx = s.Context()

	t, err := taskqueue.NewTask(type_, data)
	if err != nil {
		i.log.Error(err, "Failed to create task", "type", type_)
		return err
	}

	if err := i.client.Submit(ctx, t); err != nil {
		i.log.Error(err, "Failed to schedule task", "type", type_, "taskID", t.ID)
		return err
	}

	i.log.V(2).Info("Scheduled task", "type", type_, "taskID", t.ID)

	return nil
}
