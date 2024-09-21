package entity

import (
	"github.com/versia-pub/versia-go/ent"
	"github.com/versia-pub/versia-go/pkg/versia"
	versiacrypto "github.com/versia-pub/versia-go/pkg/versia/crypto"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
)

type InstanceMetadata struct {
	*ent.InstanceMetadata

	URI *versiautils.URL

	Moderators           []User
	ModeratorsCollection *versiautils.URL

	Admins           []User
	AdminsCollection *versiautils.URL

	SharedInbox *versiautils.URL

	PublicKey *versiacrypto.SPKIPublicKey

	Logo   *versiautils.ImageContentMap
	Banner *versiautils.ImageContentMap
}

func NewInstanceMetadata(dbData *ent.InstanceMetadata) (*InstanceMetadata, error) {
	n := &InstanceMetadata{
		InstanceMetadata: dbData,
		PublicKey:        &versiacrypto.SPKIPublicKey{},
	}

	var err error
	if n.PublicKey.Key, err = versiacrypto.ToTypedKey(dbData.PublicKeyAlgorithm, dbData.PublicKey); err != nil {
		return nil, err
	}

	if n.URI, err = versiautils.ParseURL(dbData.URI); err != nil {
		return nil, err
	}
	if dbData.SharedInboxURI != nil {
		if n.SharedInbox, err = versiautils.ParseURL(*dbData.SharedInboxURI); err != nil {
			return nil, err
		}
	}
	if dbData.ModeratorsURI != nil {
		if n.ModeratorsCollection, err = versiautils.ParseURL(*dbData.ModeratorsURI); err != nil {
			return nil, err
		}
	}
	if dbData.AdminsURI != nil {
		if n.AdminsCollection, err = versiautils.ParseURL(*dbData.AdminsURI); err != nil {
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

func (m InstanceMetadata) ToVersia() versia.InstanceMetadata {
	return versia.InstanceMetadata{
		CreatedAt:   versiautils.TimeFromStd(m.CreatedAt),
		Extensions:  m.Extensions,
		Name:        m.Name,
		Description: m.Description,
		Host:        m.Host,
		SharedInbox: m.SharedInbox,
		Moderators:  m.ModeratorsCollection,
		Admins:      m.AdminsCollection,
		Logo:        m.Logo,
		Banner:      m.Banner,
		PublicKey: versia.InstancePublicKey{
			Algorithm: m.PublicKeyAlgorithm,
			Key:       m.PublicKey,
		},
		Software: versia.InstanceSoftware{
			Name:    m.SoftwareName,
			Version: m.SoftwareVersion,
		},
		Compatibility: versia.InstanceCompatibility{
			Versions:   m.SupportedVersions,
			Extensions: m.SupportedExtensions,
		},
	}
}
