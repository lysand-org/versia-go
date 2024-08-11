package lysand

import "encoding/json"

type Follow struct {
	Entity

	// Author is the URL to the user that triggered the follow
	Author *URL `json:"author"`
	// Followee is the URL to the user that is being followed
	Followee *URL `json:"followee"`
}

func (f Follow) MarshalJSON() ([]byte, error) {
	type follow Follow
	f2 := follow(f)
	f2.Type = "Follow"
	return json.Marshal(f2)
}

type FollowAccept struct {
	Entity

	// Author is the URL to the user that accepted the follow
	Author *URL `json:"author"`
	// Follower is the URL to the user that is now following the followee
	Follower *URL `json:"follower"`
}

func (f FollowAccept) MarshalJSON() ([]byte, error) {
	type followAccept FollowAccept
	f2 := followAccept(f)
	f2.Type = "FollowAccept"
	return json.Marshal(f2)
}

type FollowReject struct {
	Entity

	// Author is the URL to the user that rejected the follow
	Author *URL `json:"author"`
	// Follower is the URL to the user that is no longer following the followee
	Follower *URL `json:"follower"`
}

func (f FollowReject) MarshalJSON() ([]byte, error) {
	type followReject FollowReject
	f2 := followReject(f)
	f2.Type = "FollowReject"
	return json.Marshal(f2)
}
