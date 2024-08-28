package task_impls

import (
	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/versia-pub/versia-go/internal/task"
)

var _ task.Manager = (*Manager)(nil)

type Manager struct {
	notes *NoteHandler

	telemetry *unitel.Telemetry
	log       logr.Logger
}

func NewManager(notes *NoteHandler, telemetry *unitel.Telemetry, log logr.Logger) *Manager {
	return &Manager{
		notes: notes,

		telemetry: telemetry,
		log:       log,
	}
}

func (m *Manager) Notes() task.NoteHandler {
	return m.notes
}
