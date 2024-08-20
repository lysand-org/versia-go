package repo_impls

import (
	"context"
	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/ent/instancemetadata"
	"github.com/lysand-org/versia-go/internal/entity"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/service"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

var _ repository.InstanceMetadataRepository = (*InstanceMetadataRepositoryImpl)(nil)

type InstanceMetadataRepositoryImpl struct {
	federationService service.FederationService

	db        *ent.Client
	log       logr.Logger
	telemetry *unitel.Telemetry
}

func NewInstanceMetadataRepositoryImpl(federationService service.FederationService, db *ent.Client, log logr.Logger, telemetry *unitel.Telemetry) repository.InstanceMetadataRepository {
	return &InstanceMetadataRepositoryImpl{
		federationService: federationService,

		db:        db,
		log:       log,
		telemetry: telemetry,
	}
}

func (i *InstanceMetadataRepositoryImpl) GetByHost(ctx context.Context, host string) (*entity.InstanceMetadata, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/InstanceMetadataRepositoryImpl.GetByHost").
		AddAttribute("host", host)
	defer s.End()
	ctx = s.Context()

	m, err := i.db.InstanceMetadata.Query().
		Where(instancemetadata.Host(host)).
		WithAdmins().
		WithModerators().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity.NewInstanceMetadata(m)
}

func (i *InstanceMetadataRepositoryImpl) ImportFromLysandByURI(ctx context.Context, uri *lysand.URL) (*entity.InstanceMetadata, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/InstanceMetadataRepositoryImpl.ImportFromLysandByURI").
		AddAttribute("uri", uri.String())
	defer s.End()
	ctx = s.Context()

	//i.federationService.

	return nil, nil
}
