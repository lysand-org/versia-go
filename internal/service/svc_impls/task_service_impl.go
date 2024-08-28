package svc_impls

import (
	"context"

	"github.com/versia-pub/versia-go/internal/service"
	"github.com/versia-pub/versia-go/internal/task"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/versia-pub/versia-go/pkg/taskqueue"
)

var _ service.TaskService = (*TaskServiceImpl)(nil)

type TaskServiceImpl struct {
	manager task.Manager

	telemetry *unitel.Telemetry
	log       logr.Logger
}

func NewTaskServiceImpl(manager task.Manager, telemetry *unitel.Telemetry, log logr.Logger) *TaskServiceImpl {
	return &TaskServiceImpl{
		manager: manager,

		telemetry: telemetry,
		log:       log,
	}
}

func (i TaskServiceImpl) ScheduleNoteTask(ctx context.Context, type_ string, data any) error {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/TaskServiceImpl.ScheduleTask")
	defer s.End()
	ctx = s.Context()

	t, err := taskqueue.NewTask(type_, data)
	if err != nil {
		i.log.Error(err, "Failed to create task", "type", type_)
		return err
	}

	if err := i.manager.Notes().Submit(ctx, t); err != nil {
		i.log.Error(err, "Failed to schedule task", "type", type_, "taskID", t.ID)
		return err
	}

	i.log.V(2).Info("Scheduled task", "type", type_, "taskID", t.ID)

	return nil
}
