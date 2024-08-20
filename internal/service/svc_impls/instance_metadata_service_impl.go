package svc_impls

import (
	"context"
	"github.com/lysand-org/versia-go/config"
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/service"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/lysand-org/versia-go/internal/entity"
)

var _ service.InstanceMetadataService = (*InstanceMetadataServiceImpl)(nil)

type InstanceMetadataServiceImpl struct {
	federationService service.FederationService

	repositories repository.Manager

	telemetry *unitel.Telemetry
	log       logr.Logger
}

func NewInstanceMetadataServiceImpl(federationService service.FederationService, repositories repository.Manager, telemetry *unitel.Telemetry, log logr.Logger) *InstanceMetadataServiceImpl {
	return &InstanceMetadataServiceImpl{
		federationService: federationService,

		repositories: repositories,

		telemetry: telemetry,
		log:       log,
	}
}

func (i InstanceMetadataServiceImpl) Ours(ctx context.Context) (*entity.InstanceMetadata, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/InstanceMetadataServiceImpl.Ours")
	defer s.End()
	ctx = s.Context()

	m, err := i.repositories.InstanceMetadata().GetByHost(ctx, config.C.Host)
	if err != nil {
		if ent.IsNotFound(err) {
			panic("could not find our own instance metadata")
		}

		return nil, err
	}

	return m, nil
}
