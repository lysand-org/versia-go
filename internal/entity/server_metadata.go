package entity

import (
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/pkg/lysand"
	versiacrypto "github.com/lysand-org/versia-go/pkg/lysand/crypto"
)

type InstanceMetadata struct {
	*ent.InstanceMetadata

	Moderators           []User
	ModeratorsCollection *lysand.URL

	Admins           []User
	AdminsCollection *lysand.URL

	SharedInbox *lysand.URL

	PublicKey *lysand.SPKIPublicKey

	Logo   *lysand.ImageContentTypeMap
	Banner *lysand.ImageContentTypeMap
}

func NewInstanceMetadata(dbData *ent.InstanceMetadata) (*InstanceMetadata, error) {
	n := &InstanceMetadata{
		InstanceMetadata: dbData,
		PublicKey:        &lysand.SPKIPublicKey{},
	}

	var err error
	if n.PublicKey.Key, err = versiacrypto.ToTypedKey(dbData.PublicKeyAlgorithm, dbData.PublicKey); err != nil {
		return nil, err
	}

	if n.SharedInbox, err = lysand.ParseURL(dbData.SharedInboxURI); err != nil {
		return nil, err
	}
	if dbData.ModeratorsURI != nil {
		if n.ModeratorsCollection, err = lysand.ParseURL(*dbData.ModeratorsURI); err != nil {
			return nil, err
		}
	}
	if dbData.AdminsURI != nil {
		if n.AdminsCollection, err = lysand.ParseURL(*dbData.AdminsURI); err != nil {
			return nil, err
		}
	}

	for _, r := range dbData.Edges.Moderators {
		u, err := NewUser(r)
		if err != nil {
			return nil, err
		}

		n.Moderators = append(n.Moderators, *u)
	}

	for _, r := range dbData.Edges.Admins {
		u, err := NewUser(r)
		if err != nil {
			return nil, err
		}

		n.Admins = append(n.Admins, *u)
	}

	return n, nil
}

func (m InstanceMetadata) ToLysand() lysand.InstanceMetadata {
	return lysand.InstanceMetadata{
		Extensions:  m.Extensions,
		Name:        m.Name,
		Description: m.Description,
		Host:        m.Host,
		SharedInbox: m.SharedInbox,
		Moderators:  m.ModeratorsCollection,
		Admins:      m.AdminsCollection,
		Logo:        m.Logo,
		Banner:      m.Banner,
		PublicKey: lysand.InstancePublicKey{
			Algorithm: m.PublicKeyAlgorithm,
			Key:       m.PublicKey,
		},
		Software: lysand.InstanceSoftware{
			Name:    m.SoftwareName,
			Version: m.SoftwareVersion,
		},
		Compatibility: lysand.InstanceCompatibility{
			Versions:   m.SupportedVersions,
			Extensions: m.SupportedExtensions,
		},
	}
}
