package entity

import (
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/pkg/versia"
	versiautils "github.com/lysand-org/versia-go/pkg/versia/utils"
)

type Follow struct {
	*ent.Follow

	URI         *versiautils.URL
	FollowerURI *versiautils.URL
	FolloweeURI *versiautils.URL
}

func NewFollow(dbFollow *ent.Follow) (*Follow, error) {
	f := &Follow{Follow: dbFollow}

	var err error

	f.URI, err = versiautils.ParseURL(dbFollow.URI)
	if err != nil {
		return nil, err
	}

	f.FollowerURI, err = versiautils.ParseURL(dbFollow.Edges.Follower.URI)
	if err != nil {
		return nil, err
	}

	f.FolloweeURI, err = versiautils.ParseURL(dbFollow.Edges.Followee.URI)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f Follow) ToLysand() *versia.Follow {
	return &versia.Follow{
		Entity: versia.Entity{
			ID:         f.ID,
			URI:        f.URI,
			CreatedAt:  versiautils.Time(f.CreatedAt),
			Extensions: f.Extensions,
		},
		Author:   f.FollowerURI,
		Followee: f.FolloweeURI,
	}
}

func (f Follow) ToLysandAccept() *versia.FollowAccept {
	return &versia.FollowAccept{
		Entity: versia.Entity{
			ID:         f.ID,
			URI:        f.URI,
			CreatedAt:  versiautils.Time(f.CreatedAt),
			Extensions: f.Extensions,
		},
		Author:   f.FolloweeURI,
		Follower: f.FollowerURI,
	}
}

func (f Follow) ToLysandReject() *versia.FollowReject {
	return &versia.FollowReject{
		Entity: versia.Entity{
			ID:         f.ID,
			URI:        f.URI,
			CreatedAt:  versiautils.Time(f.CreatedAt),
			Extensions: f.Extensions,
		},
		Author:   f.FolloweeURI,
		Follower: f.FollowerURI,
	}
}
