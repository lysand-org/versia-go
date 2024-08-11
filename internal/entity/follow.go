package entity

import (
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

type Follow struct {
	*ent.Follow

	URI         *lysand.URL
	FollowerURI *lysand.URL
	FolloweeURI *lysand.URL
}

func NewFollow(dbFollow *ent.Follow) (*Follow, error) {
	f := &Follow{Follow: dbFollow}

	var err error

	f.URI, err = lysand.ParseURL(dbFollow.URI)
	if err != nil {
		return nil, err
	}

	f.FollowerURI, err = lysand.ParseURL(dbFollow.Edges.Follower.URI)
	if err != nil {
		return nil, err
	}

	f.FolloweeURI, err = lysand.ParseURL(dbFollow.Edges.Followee.URI)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f Follow) ToLysand() *lysand.Follow {
	return &lysand.Follow{
		Entity: lysand.Entity{
			ID:         f.ID,
			URI:        f.URI,
			CreatedAt:  lysand.TimeFromStd(f.CreatedAt),
			Extensions: f.Extensions,
		},
		Author:   f.FollowerURI,
		Followee: f.FolloweeURI,
	}
}

func (f Follow) ToLysandAccept() *lysand.FollowAccept {
	return &lysand.FollowAccept{
		Entity: lysand.Entity{
			ID:         f.ID,
			URI:        f.URI,
			CreatedAt:  lysand.TimeFromStd(f.CreatedAt),
			Extensions: f.Extensions,
		},
		Author:   f.FolloweeURI,
		Follower: f.FollowerURI,
	}
}

func (f Follow) ToLysandReject() *lysand.FollowReject {
	return &lysand.FollowReject{
		Entity: lysand.Entity{
			ID:         f.ID,
			URI:        f.URI,
			CreatedAt:  lysand.TimeFromStd(f.CreatedAt),
			Extensions: f.Extensions,
		},
		Author:   f.FolloweeURI,
		Follower: f.FollowerURI,
	}
}
