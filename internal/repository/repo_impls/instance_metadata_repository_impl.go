package repo_impls

import (
	"context"
	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/ldez/mimetype"
	"github.com/versia-pub/versia-go/ent"
	"github.com/versia-pub/versia-go/ent/instancemetadata"
	"github.com/versia-pub/versia-go/internal/entity"
	"github.com/versia-pub/versia-go/internal/repository"
	"github.com/versia-pub/versia-go/internal/service"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
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

func (i *InstanceMetadataRepositoryImpl) Resolve(ctx context.Context, host string) (*entity.InstanceMetadata, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/InstanceMetadataRepositoryImpl.Resolve").
		AddAttribute("host", host)
	defer s.End()
	ctx = s.Context()

	metadata, err := i.federationService.DiscoverInstance(ctx, host)
	if err != nil {
		return nil, err
	}

	logoURL, logoType := preferredImage(metadata.Logo)
	bannerURL, bannerType := preferredImage(metadata.Banner)

	meta, err := i.db.InstanceMetadata.Create().
		SetName(metadata.Name).
		SetNillableDescription(metadata.Description).
		SetHost(metadata.Host).
		SetPublicKey(metadata.PublicKey.Key.ToKey().([]byte)).
		SetPublicKeyAlgorithm(metadata.PublicKey.Key.Algorithm).
		SetSoftwareName(metadata.Software.Name).
		SetSoftwareVersion(metadata.Software.Version).
		SetNillableSharedInboxURI(urlToStrPtr(metadata.SharedInbox)).
		SetNillableModeratorsURI(urlToStrPtr(metadata.Moderators)).
		SetNillableAdminsURI(urlToStrPtr(metadata.Admins)).
		SetNillableLogoEndpoint(logoURL).
		SetNillableLogoMimeType(logoType).
		SetNillableBannerEndpoint(bannerURL).
		SetNillableBannerMimeType(bannerType).
		SetSupportedVersions(metadata.Compatibility.Versions).
		SetSupportedExtensions(metadata.Compatibility.Extensions).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity.NewInstanceMetadata(meta)
}

func urlToStrPtr(u *versiautils.URL) *string {
	if u == nil {
		return nil
	}

	v := u.String()

	return &v
}

var preferredImageMIMETypes = []string{
	mimetype.ImageWebp,
	mimetype.ImageJxl,
	mimetype.ImagePng,
	mimetype.ImageJpeg,
	mimetype.ImageGif,
	mimetype.ImageBmp,
}

func preferredImage(i *versiautils.ImageContentMap) (*string, *string) {
	if i == nil {
		return nil, nil
	}

	m := i.Map()

	for _, type_ := range preferredImageMIMETypes {
		if v, ok := m[type_]; !ok {
			return urlToStrPtr(v.Content), &type_
		}
	}

	return nil, nil
}
