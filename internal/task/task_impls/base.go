package task_impls

import "git.devminer.xyz/devminer/unitel"

type baseHandler struct {
	telemetry *unitel.Telemetry
}

func newBaseHandler() *baseHandler {
	return &baseHandler{}
}
