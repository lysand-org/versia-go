// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/ent/attachment"
	"github.com/lysand-org/versia-go/ent/user"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

// Attachment is the model entity for the Attachment schema.
type Attachment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// IsRemote holds the value of the "isRemote" field.
	IsRemote bool `json:"isRemote,omitempty"`
	// URI holds the value of the "uri" field.
	URI string `json:"uri,omitempty"`
	// Extensions holds the value of the "extensions" field.
	Extensions lysand.Extensions `json:"extensions,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Sha256 holds the value of the "sha256" field.
	Sha256 []byte `json:"sha256,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Blurhash holds the value of the "blurhash" field.
	Blurhash *string `json:"blurhash,omitempty"`
	// Height holds the value of the "height" field.
	Height *int `json:"height,omitempty"`
	// Width holds the value of the "width" field.
	Width *int `json:"width,omitempty"`
	// Fps holds the value of the "fps" field.
	Fps *int `json:"fps,omitempty"`
	// MimeType holds the value of the "mimeType" field.
	MimeType string `json:"mimeType,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttachmentQuery when eager-loading is set.
	Edges             AttachmentEdges `json:"edges"`
	attachment_author *uuid.UUID
	note_attachments  *uuid.UUID
	selectValues      sql.SelectValues
}

// AttachmentEdges holds the relations/edges for other nodes in the graph.
type AttachmentEdges struct {
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttachmentEdges) AuthorOrErr() (*User, error) {
	if e.Author != nil {
		return e.Author, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "author"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attachment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attachment.FieldExtensions, attachment.FieldSha256:
			values[i] = new([]byte)
		case attachment.FieldIsRemote:
			values[i] = new(sql.NullBool)
		case attachment.FieldSize, attachment.FieldHeight, attachment.FieldWidth, attachment.FieldFps:
			values[i] = new(sql.NullInt64)
		case attachment.FieldURI, attachment.FieldDescription, attachment.FieldBlurhash, attachment.FieldMimeType:
			values[i] = new(sql.NullString)
		case attachment.FieldCreatedAt, attachment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case attachment.FieldID:
			values[i] = new(uuid.UUID)
		case attachment.ForeignKeys[0]: // attachment_author
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case attachment.ForeignKeys[1]: // note_attachments
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attachment fields.
func (a *Attachment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attachment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case attachment.FieldIsRemote:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isRemote", values[i])
			} else if value.Valid {
				a.IsRemote = value.Bool
			}
		case attachment.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				a.URI = value.String
			}
		case attachment.FieldExtensions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field extensions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Extensions); err != nil {
					return fmt.Errorf("unmarshal field extensions: %w", err)
				}
			}
		case attachment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case attachment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case attachment.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				a.Description = value.String
			}
		case attachment.FieldSha256:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field sha256", values[i])
			} else if value != nil {
				a.Sha256 = *value
			}
		case attachment.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				a.Size = int(value.Int64)
			}
		case attachment.FieldBlurhash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field blurhash", values[i])
			} else if value.Valid {
				a.Blurhash = new(string)
				*a.Blurhash = value.String
			}
		case attachment.FieldHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				a.Height = new(int)
				*a.Height = int(value.Int64)
			}
		case attachment.FieldWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				a.Width = new(int)
				*a.Width = int(value.Int64)
			}
		case attachment.FieldFps:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fps", values[i])
			} else if value.Valid {
				a.Fps = new(int)
				*a.Fps = int(value.Int64)
			}
		case attachment.FieldMimeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mimeType", values[i])
			} else if value.Valid {
				a.MimeType = value.String
			}
		case attachment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field attachment_author", values[i])
			} else if value.Valid {
				a.attachment_author = new(uuid.UUID)
				*a.attachment_author = *value.S.(*uuid.UUID)
			}
		case attachment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field note_attachments", values[i])
			} else if value.Valid {
				a.note_attachments = new(uuid.UUID)
				*a.note_attachments = *value.S.(*uuid.UUID)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Attachment.
// This includes values selected through modifiers, order, etc.
func (a *Attachment) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryAuthor queries the "author" edge of the Attachment entity.
func (a *Attachment) QueryAuthor() *UserQuery {
	return NewAttachmentClient(a.config).QueryAuthor(a)
}

// Update returns a builder for updating this Attachment.
// Note that you need to call Attachment.Unwrap() before calling this method if this Attachment
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attachment) Update() *AttachmentUpdateOne {
	return NewAttachmentClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Attachment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attachment) Unwrap() *Attachment {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attachment is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attachment) String() string {
	var builder strings.Builder
	builder.WriteString("Attachment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("isRemote=")
	builder.WriteString(fmt.Sprintf("%v", a.IsRemote))
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(a.URI)
	builder.WriteString(", ")
	builder.WriteString("extensions=")
	builder.WriteString(fmt.Sprintf("%v", a.Extensions))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(a.Description)
	builder.WriteString(", ")
	builder.WriteString("sha256=")
	builder.WriteString(fmt.Sprintf("%v", a.Sha256))
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", a.Size))
	builder.WriteString(", ")
	if v := a.Blurhash; v != nil {
		builder.WriteString("blurhash=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := a.Height; v != nil {
		builder.WriteString("height=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := a.Width; v != nil {
		builder.WriteString("width=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := a.Fps; v != nil {
		builder.WriteString("fps=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("mimeType=")
	builder.WriteString(a.MimeType)
	builder.WriteByte(')')
	return builder.String()
}

// Attachments is a parsable slice of Attachment.
type Attachments []*Attachment
