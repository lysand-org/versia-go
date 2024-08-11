package entity

import (
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

type Note struct {
	*ent.Note
	URI         *lysand.URL
	Content     lysand.TextContentTypeMap
	Author      *User
	Mentions    []User
	MentionURIs []lysand.URL
}

func NewNote(dbNote *ent.Note) (*Note, error) {
	n := &Note{
		Note: dbNote,
		Content: lysand.TextContentTypeMap{
			"text/plain": lysand.TextContent{Content: dbNote.Content},
		},
		Mentions:    make([]User, 0, len(dbNote.Edges.Mentions)),
		MentionURIs: make([]lysand.URL, 0, len(dbNote.Edges.Mentions)),
	}

	var err error
	if n.URI, err = lysand.ParseURL(dbNote.URI); err != nil {
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

func (n Note) ToLysand() lysand.Note {
	return lysand.Note{
		Entity: lysand.Entity{
			ID:         n.ID,
			URI:        n.URI,
			CreatedAt:  lysand.TimeFromStd(n.CreatedAt),
			Extensions: n.Extensions,
		},
		Author:      n.Author.URI,
		Content:     n.Content,
		Category:    nil,
		Device:      nil,
		Previews:    nil,
		Group:       nil,
		Attachments: nil,
		RepliesTo:   nil,
		Quoting:     nil,
		Mentions:    n.MentionURIs,
		Subject:     n.Subject,
		IsSensitive: &n.IsSensitive,
		Visibility:  lysand.PublicationVisibility(n.Visibility),
	}
}
