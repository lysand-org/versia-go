package entity

import (
	"github.com/versia-pub/versia-go/ent"
	"github.com/versia-pub/versia-go/pkg/versia"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
)

type Note struct {
	*ent.Note
	URI         *versiautils.URL
	Content     versiautils.TextContentTypeMap
	Author      *User
	Mentions    []User
	MentionURIs []versiautils.URL
}

func NewNote(dbNote *ent.Note) (*Note, error) {
	n := &Note{
		Note: dbNote,
		Content: versiautils.TextContentTypeMap{
			"text/plain": versiautils.TextContent{Content: dbNote.Content},
		},
		Mentions:    make([]User, 0, len(dbNote.Edges.Mentions)),
		MentionURIs: make([]versiautils.URL, 0, len(dbNote.Edges.Mentions)),
	}

	var err error
	if n.URI, err = versiautils.ParseURL(dbNote.URI); err != nil {
		return nil, err
	}
	if n.Author, err = NewUser(dbNote.Edges.Author); err != nil {
		return nil, err
	}

	for _, m := range dbNote.Edges.Mentions {
		u, err := NewUser(m)
		if err != nil {
			return nil, err
		}

		n.Mentions = append(n.Mentions, *u)
		n.MentionURIs = append(n.MentionURIs, *u.URI)
	}

	return n, nil
}

func (n Note) ToVersia() versia.Note {
	return versia.Note{
		Entity: versia.Entity{
			ID:         n.ID,
			URI:        n.URI,
			CreatedAt:  versiautils.Time(n.CreatedAt),
			Extensions: n.Extensions,
		},
		Author:   n.Author.URI,
		Content:  n.Content,
		Category: nil,
		Device:   nil,
		Previews: nil,
		// TODO: Get from database
		Group:       nil,
		Attachments: nil,
		RepliesTo:   nil,
		Quotes:      nil,
		Mentions:    n.MentionURIs,
		Subject:     n.Subject,
		IsSensitive: &n.IsSensitive,
	}
}
