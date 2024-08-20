package utils

import (
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/config"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

var dicebearURL = &url.URL{
	Scheme: "https",
	Host:   "api.dicebear.com",
	Path:   "9.x/adventurer/svg",
}

func UserAPIURL(uuid uuid.UUID) *lysand.URL {
	newPath := &url.URL{Path: fmt.Sprintf("/api/users/%s/", uuid.String())}
	return lysand.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func DefaultAvatarURL(uuid uuid.UUID) *lysand.URL {
	u := &url.URL{}
	q := u.Query()
	q.Set("seed", uuid.String())
	u.RawQuery = q.Encode()

	return lysand.URLFromStd(dicebearURL.ResolveReference(u))
}

func UserInboxAPIURL(uuid uuid.UUID) *lysand.URL {
	newPath := &url.URL{Path: "./inbox"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserOutboxAPIURL(uuid uuid.UUID) *lysand.URL {
	newPath := &url.URL{Path: "./outbox"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserFollowersAPIURL(uuid uuid.UUID) *lysand.URL {
	newPath := &url.URL{Path: "./followers"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserFollowingAPIURL(uuid uuid.UUID) *lysand.URL {
	newPath := &url.URL{Path: "./following"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserFeaturedAPIURL(uuid uuid.UUID) *lysand.URL {
	newPath := &url.URL{Path: "./featured"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserLikesAPIURL(uuid uuid.UUID) *lysand.URL {
	newPath := &url.URL{Path: "./likes"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserDislikesAPIURL(uuid uuid.UUID) *lysand.URL {
	newPath := &url.URL{Path: "./dislikes"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func FollowAPIURL(uuid uuid.UUID) *lysand.URL {
	newPath := &url.URL{Path: fmt.Sprintf("/api/follows/%s/", uuid.String())}
	return lysand.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func NoteAPIURL(uuid uuid.UUID) *lysand.URL {
	newPath := &url.URL{Path: fmt.Sprintf("/api/notes/%s/", uuid.String())}
	return lysand.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func InstanceMetadataAPIURL() *lysand.URL {
	newPath := &url.URL{Path: "/.well-known/versia/"}
	return lysand.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func InstanceMetadataAdminsAPIURL() *lysand.URL {
	newPath := &url.URL{Path: "/.well-known/versia/admins/"}
	return lysand.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func InstanceMetadataModeratorsAPIURL() *lysand.URL {
	newPath := &url.URL{Path: "/.well-known/versia/moderators/"}
	return lysand.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func SharedInboxAPIURL() *lysand.URL {
	newPath := &url.URL{Path: "/api/inbox/"}
	return lysand.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}
