package tasks

import "context"

type FederateFollowData struct {
	FollowID string `json:"followID"`
}

func (t *Handler) FederateFollow(ctx context.Context, data FederateNoteData) error {
	return nil
}
