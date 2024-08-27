package entity

import (
	"github.com/versia-pub/versia-go/ent"
	"github.com/versia-pub/versia-go/pkg/versia"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
)

type Follow struct {
	*ent.Follow

	URI      *versiautils.URL
	Follower *User
	Followee *User
}

func NewFollow(dbData *ent.Follow) (*Follow, error) {
	f := &Follow{Follow: dbData}

	var err error

	if f.URI, err = versiautils.ParseURL(dbData.URI); err != nil {
		return nil, err
	}

	if f.Follower, err = NewUser(dbData.Edges.Follower); err != nil {
		return nil, err
	}

	if f.Followee, err = NewUser(dbData.Edges.Followee); err != nil {
		return nil, err
	}

	return f, nil
}

func (f Follow) ToVersia() *versia.Follow {
	return &versia.Follow{
		Entity: versia.Entity{
			ID:         f.ID,
			URI:        f.URI,
			CreatedAt:  versiautils.Time(f.CreatedAt),
			Extensions: f.Extensions,
		},
		Author:   f.Follower.URI,
		Followee: f.Followee.URI,
	}
}

func (f Follow) ToVersiaAccept() *versia.FollowAccept {
	return &versia.FollowAccept{
		Entity: versia.Entity{
			ID:         f.ID,
			URI:        f.URI,
			CreatedAt:  versiautils.Time(f.CreatedAt),
			Extensions: f.Extensions,
		},
		Author:   f.Followee.URI,
		Follower: f.Follower.URI,
	}
}

func (f Follow) ToVersiaReject() *versia.FollowReject {
	return &versia.FollowReject{
		Entity: versia.Entity{
			ID:         f.ID,
			URI:        f.URI,
			CreatedAt:  versiautils.Time(f.CreatedAt),
			Extensions: f.Extensions,
		},
		Author:   f.Followee.URI,
		Follower: f.Follower.URI,
	}
}
