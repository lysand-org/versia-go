package versia

import (
	"encoding/json"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
)

// Follow defines a follow relationship between two users. For more information, see the [Spec].
//
// Once a follow relationship is established, the followee's instance should send all new notes from the followee to
// the follower's inbox.
//
// [Spec]: https://versia.pub/entities/follow
type Follow struct {
	Entity

	// Author is the URL to the user that triggered the follow
	Author *versiautils.URL `json:"author"`

	// Followee is the URL to the user that is being followed
	Followee *versiautils.URL `json:"followee"`
}

func (f Follow) MarshalJSON() ([]byte, error) {
	type follow Follow
	f2 := follow(f)
	f2.Type = "Follow"
	return json.Marshal(f2)
}

// FollowAccept accepts a Follow request, which will form the follow relationship between the two parties.
// For more information, see the [Spec].
//
// This can only be sent by the Followee.
//
// [Spec]: https://versia.pub/entities/follow-accept
type FollowAccept struct {
	Entity

	// Author is the URL to the user that accepted the follow
	Author *versiautils.URL `json:"author"`

	// Follower is the URL to the user that is now following the followee
	Follower *versiautils.URL `json:"follower"`
}

func (f FollowAccept) MarshalJSON() ([]byte, error) {
	type followAccept FollowAccept
	f2 := followAccept(f)
	f2.Type = "FollowAccept"
	return json.Marshal(f2)
}

// FollowReject rejects a Follow request, which will dismiss the follow relationship between the two parties.
// For more information, see the [Spec].
//
// This can only be sent by the Followee and should not be confused with Unfollow, which can only be sent by the Follower.
// FollowReject can still be sent after the relationship has been formed.
//
// [Spec]: https://versia.pub/entities/follow-reject
type FollowReject struct {
	Entity

	// Author is the URL to the user that rejected the follow
	Author *versiautils.URL `json:"author"`

	// Follower is the URL to the user that is no longer following the followee
	Follower *versiautils.URL `json:"follower"`
}

func (f FollowReject) MarshalJSON() ([]byte, error) {
	type followReject FollowReject
	f2 := followReject(f)
	f2.Type = "FollowReject"
	return json.Marshal(f2)
}

// Unfollow disbands request, which will disband the follow relationship between the two parties.
// For more information, see the [Spec].
//
// This can only be sent by the Follower and should not be confused with FollowReject, which can only be sent by the Followee.
//
// [Spec]: https://versia.pub/entities/unfollow
type Unfollow struct {
	Entity

	// Author is the URL to the user that unfollowed the followee
	Author *versiautils.URL `json:"author"`

	// Followee is the URL to the user that has been followed
	Followee *versiautils.URL `json:"follower"`
}

func (f Unfollow) MarshalJSON() ([]byte, error) {
	type a Unfollow
	u := a(f)
	u.Type = "Unfollow"
	return json.Marshal(u)
}
