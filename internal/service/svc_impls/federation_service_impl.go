package svc_impls

import (
	"context"
	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/lysand-org/versia-go/internal/entity"
	"github.com/lysand-org/versia-go/internal/service"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

var _ service.FederationService = (*FederationServiceImpl)(nil)

type FederationServiceImpl struct {
	federationClient *lysand.FederationClient

	telemetry *unitel.Telemetry

	log logr.Logger
}

func NewFederationServiceImpl(federationClient *lysand.FederationClient, telemetry *unitel.Telemetry, log logr.Logger) *FederationServiceImpl {
	return &FederationServiceImpl{
		federationClient: federationClient,
		telemetry:        telemetry,
		log:              log,
	}
}

func (i FederationServiceImpl) SendToInbox(ctx context.Context, author *entity.User, target *entity.User, object any) ([]byte, error) {
	s := i.telemetry.StartSpan(ctx, "function", "service/svc_impls.FederationServiceImpl.SendToInbox")
	defer s.End()
	ctx = s.Context()

	response, err := i.federationClient.SendToInbox(ctx, author.Signer, target.ToLysand(), object)
	if err != nil {
		i.log.Error(err, "Failed to send to inbox", "author", author.ID, "target", target.ID)
		return response, err
	}

	return response, nil
}

func (i FederationServiceImpl) GetUser(ctx context.Context, uri *lysand.URL) (*lysand.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "service/svc_impls.FederationServiceImpl.GetUser")
	defer s.End()
	ctx = s.Context()

	u, err := i.federationClient.GetUser(ctx, uri.ToStd())
	if err != nil {
		i.log.Error(err, "Failed to fetch remote user", "uri", uri)
		return nil, err
	}

	return u, nil
}
