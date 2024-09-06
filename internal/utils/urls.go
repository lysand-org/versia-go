package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/internal/config"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
	"net/url"
)

var dicebearURL = &url.URL{
	Scheme: "https",
	Host:   "api.dicebear.com",
	Path:   "9.x/adventurer/svg",
}

func UserAPIURL(uuid uuid.UUID) *versiautils.URL {
	newPath := &url.URL{Path: fmt.Sprintf("/api/users/%s/", uuid.String())}
	return versiautils.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func DefaultAvatarURL(uuid uuid.UUID) *versiautils.URL {
	u := &url.URL{}
	q := u.Query()
	q.Set("seed", uuid.String())
	u.RawQuery = q.Encode()

	return versiautils.URLFromStd(dicebearURL.ResolveReference(u))
}

func UserInboxAPIURL(uuid uuid.UUID) *versiautils.URL {
	newPath := &url.URL{Path: "./inbox"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserOutboxAPIURL(uuid uuid.UUID) *versiautils.URL {
	newPath := &url.URL{Path: "./outbox"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserFollowersAPIURL(uuid uuid.UUID) *versiautils.URL {
	newPath := &url.URL{Path: "./followers"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserFollowingAPIURL(uuid uuid.UUID) *versiautils.URL {
	newPath := &url.URL{Path: "./following"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserFeaturedAPIURL(uuid uuid.UUID) *versiautils.URL {
	newPath := &url.URL{Path: "./featured"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserLikesAPIURL(uuid uuid.UUID) *versiautils.URL {
	newPath := &url.URL{Path: "./likes"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func UserDislikesAPIURL(uuid uuid.UUID) *versiautils.URL {
	newPath := &url.URL{Path: "./dislikes"}
	return UserAPIURL(uuid).ResolveReference(newPath)
}

func FollowAPIURL(uuid uuid.UUID) *versiautils.URL {
	newPath := &url.URL{Path: fmt.Sprintf("/api/follows/%s/", uuid.String())}
	return versiautils.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func NoteAPIURL(uuid uuid.UUID) *versiautils.URL {
	newPath := &url.URL{Path: fmt.Sprintf("/api/notes/%s/", uuid.String())}
	return versiautils.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func InstanceMetadataAPIURL() *versiautils.URL {
	newPath := &url.URL{Path: "/.well-known/versia/"}
	return versiautils.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func InstanceMetadataAdminsAPIURL() *versiautils.URL {
	newPath := &url.URL{Path: "/.well-known/versia/admins/"}
	return versiautils.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func InstanceMetadataModeratorsAPIURL() *versiautils.URL {
	newPath := &url.URL{Path: "/.well-known/versia/moderators/"}
	return versiautils.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}

func SharedInboxAPIURL() *versiautils.URL {
	newPath := &url.URL{Path: "/api/inbox/"}
	return versiautils.URLFromStd(config.C.PublicAddress.ResolveReference(newPath))
}
