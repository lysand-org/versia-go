// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsRemote holds the string denoting the isremote field in the database.
	FieldIsRemote = "is_remote"
	// FieldURI holds the string denoting the uri field in the database.
	FieldURI = "uri"
	// FieldExtensions holds the string denoting the extensions field in the database.
	FieldExtensions = "extensions"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPasswordHash holds the string denoting the passwordhash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldDisplayName holds the string denoting the displayname field in the database.
	FieldDisplayName = "display_name"
	// FieldBiography holds the string denoting the biography field in the database.
	FieldBiography = "biography"
	// FieldPublicKey holds the string denoting the publickey field in the database.
	FieldPublicKey = "public_key"
	// FieldPublicKeyActor holds the string denoting the publickeyactor field in the database.
	FieldPublicKeyActor = "public_key_actor"
	// FieldPublicKeyAlgorithm holds the string denoting the publickeyalgorithm field in the database.
	FieldPublicKeyAlgorithm = "public_key_algorithm"
	// FieldPrivateKey holds the string denoting the privatekey field in the database.
	FieldPrivateKey = "private_key"
	// FieldIndexable holds the string denoting the indexable field in the database.
	FieldIndexable = "indexable"
	// FieldPrivacyLevel holds the string denoting the privacylevel field in the database.
	FieldPrivacyLevel = "privacy_level"
	// FieldFields holds the string denoting the fields field in the database.
	FieldFields = "fields"
	// FieldInbox holds the string denoting the inbox field in the database.
	FieldInbox = "inbox"
	// FieldFeatured holds the string denoting the featured field in the database.
	FieldFeatured = "featured"
	// FieldFollowers holds the string denoting the followers field in the database.
	FieldFollowers = "followers"
	// FieldFollowing holds the string denoting the following field in the database.
	FieldFollowing = "following"
	// FieldOutbox holds the string denoting the outbox field in the database.
	FieldOutbox = "outbox"
	// EdgeAvatarImage holds the string denoting the avatarimage edge name in mutations.
	EdgeAvatarImage = "avatarImage"
	// EdgeHeaderImage holds the string denoting the headerimage edge name in mutations.
	EdgeHeaderImage = "headerImage"
	// EdgeAuthoredNotes holds the string denoting the authorednotes edge name in mutations.
	EdgeAuthoredNotes = "authoredNotes"
	// EdgeMentionedNotes holds the string denoting the mentionednotes edge name in mutations.
	EdgeMentionedNotes = "mentionedNotes"
	// EdgeServers holds the string denoting the servers edge name in mutations.
	EdgeServers = "servers"
	// EdgeModeratedServers holds the string denoting the moderatedservers edge name in mutations.
	EdgeModeratedServers = "moderatedServers"
	// EdgeAdministeredServers holds the string denoting the administeredservers edge name in mutations.
	EdgeAdministeredServers = "administeredServers"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AvatarImageTable is the table that holds the avatarImage relation/edge.
	AvatarImageTable = "users"
	// AvatarImageInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	AvatarImageInverseTable = "images"
	// AvatarImageColumn is the table column denoting the avatarImage relation/edge.
	AvatarImageColumn = "user_avatar_image"
	// HeaderImageTable is the table that holds the headerImage relation/edge.
	HeaderImageTable = "users"
	// HeaderImageInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	HeaderImageInverseTable = "images"
	// HeaderImageColumn is the table column denoting the headerImage relation/edge.
	HeaderImageColumn = "user_header_image"
	// AuthoredNotesTable is the table that holds the authoredNotes relation/edge.
	AuthoredNotesTable = "notes"
	// AuthoredNotesInverseTable is the table name for the Note entity.
	// It exists in this package in order to avoid circular dependency with the "note" package.
	AuthoredNotesInverseTable = "notes"
	// AuthoredNotesColumn is the table column denoting the authoredNotes relation/edge.
	AuthoredNotesColumn = "note_author"
	// MentionedNotesTable is the table that holds the mentionedNotes relation/edge. The primary key declared below.
	MentionedNotesTable = "note_mentions"
	// MentionedNotesInverseTable is the table name for the Note entity.
	// It exists in this package in order to avoid circular dependency with the "note" package.
	MentionedNotesInverseTable = "notes"
	// ServersTable is the table that holds the servers relation/edge. The primary key declared below.
	ServersTable = "instance_metadata_users"
	// ServersInverseTable is the table name for the InstanceMetadata entity.
	// It exists in this package in order to avoid circular dependency with the "instancemetadata" package.
	ServersInverseTable = "instance_metadata"
	// ModeratedServersTable is the table that holds the moderatedServers relation/edge. The primary key declared below.
	ModeratedServersTable = "instance_metadata_moderators"
	// ModeratedServersInverseTable is the table name for the InstanceMetadata entity.
	// It exists in this package in order to avoid circular dependency with the "instancemetadata" package.
	ModeratedServersInverseTable = "instance_metadata"
	// AdministeredServersTable is the table that holds the administeredServers relation/edge. The primary key declared below.
	AdministeredServersTable = "instance_metadata_admins"
	// AdministeredServersInverseTable is the table name for the InstanceMetadata entity.
	// It exists in this package in order to avoid circular dependency with the "instancemetadata" package.
	AdministeredServersInverseTable = "instance_metadata"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldIsRemote,
	FieldURI,
	FieldExtensions,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldPasswordHash,
	FieldDisplayName,
	FieldBiography,
	FieldPublicKey,
	FieldPublicKeyActor,
	FieldPublicKeyAlgorithm,
	FieldPrivateKey,
	FieldIndexable,
	FieldPrivacyLevel,
	FieldFields,
	FieldInbox,
	FieldFeatured,
	FieldFollowers,
	FieldFollowing,
	FieldOutbox,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_avatar_image",
	"user_header_image",
}

var (
	// MentionedNotesPrimaryKey and MentionedNotesColumn2 are the table columns denoting the
	// primary key for the mentionedNotes relation (M2M).
	MentionedNotesPrimaryKey = []string{"note_id", "user_id"}
	// ServersPrimaryKey and ServersColumn2 are the table columns denoting the
	// primary key for the servers relation (M2M).
	ServersPrimaryKey = []string{"instance_metadata_id", "user_id"}
	// ModeratedServersPrimaryKey and ModeratedServersColumn2 are the table columns denoting the
	// primary key for the moderatedServers relation (M2M).
	ModeratedServersPrimaryKey = []string{"instance_metadata_id", "user_id"}
	// AdministeredServersPrimaryKey and AdministeredServersColumn2 are the table columns denoting the
	// primary key for the administeredServers relation (M2M).
	AdministeredServersPrimaryKey = []string{"instance_metadata_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// URIValidator is a validator for the "uri" field. It is called by the builders before save.
	URIValidator func(string) error
	// DefaultExtensions holds the default value on creation for the "extensions" field.
	DefaultExtensions lysand.Extensions
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DisplayNameValidator is a validator for the "displayName" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// DefaultIndexable holds the default value on creation for the "indexable" field.
	DefaultIndexable bool
	// DefaultFields holds the default value on creation for the "fields" field.
	DefaultFields []lysand.Field
	// InboxValidator is a validator for the "inbox" field. It is called by the builders before save.
	InboxValidator func(string) error
	// FeaturedValidator is a validator for the "featured" field. It is called by the builders before save.
	FeaturedValidator func(string) error
	// FollowersValidator is a validator for the "followers" field. It is called by the builders before save.
	FollowersValidator func(string) error
	// FollowingValidator is a validator for the "following" field. It is called by the builders before save.
	FollowingValidator func(string) error
	// OutboxValidator is a validator for the "outbox" field. It is called by the builders before save.
	OutboxValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// PrivacyLevel defines the type for the "privacyLevel" enum field.
type PrivacyLevel string

// PrivacyLevelPublic is the default value of the PrivacyLevel enum.
const DefaultPrivacyLevel = PrivacyLevelPublic

// PrivacyLevel values.
const (
	PrivacyLevelPublic     PrivacyLevel = "public"
	PrivacyLevelRestricted PrivacyLevel = "restricted"
	PrivacyLevelPrivate    PrivacyLevel = "private"
)

func (pl PrivacyLevel) String() string {
	return string(pl)
}

// PrivacyLevelValidator is a validator for the "privacyLevel" field enum values. It is called by the builders before save.
func PrivacyLevelValidator(pl PrivacyLevel) error {
	switch pl {
	case PrivacyLevelPublic, PrivacyLevelRestricted, PrivacyLevelPrivate:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for privacyLevel field: %q", pl)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsRemote orders the results by the isRemote field.
func ByIsRemote(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsRemote, opts...).ToFunc()
}

// ByURI orders the results by the uri field.
func ByURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURI, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByDisplayName orders the results by the displayName field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByBiography orders the results by the biography field.
func ByBiography(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBiography, opts...).ToFunc()
}

// ByPublicKeyActor orders the results by the publicKeyActor field.
func ByPublicKeyActor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicKeyActor, opts...).ToFunc()
}

// ByPublicKeyAlgorithm orders the results by the publicKeyAlgorithm field.
func ByPublicKeyAlgorithm(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicKeyAlgorithm, opts...).ToFunc()
}

// ByIndexable orders the results by the indexable field.
func ByIndexable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndexable, opts...).ToFunc()
}

// ByPrivacyLevel orders the results by the privacyLevel field.
func ByPrivacyLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivacyLevel, opts...).ToFunc()
}

// ByInbox orders the results by the inbox field.
func ByInbox(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInbox, opts...).ToFunc()
}

// ByFeatured orders the results by the featured field.
func ByFeatured(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeatured, opts...).ToFunc()
}

// ByFollowers orders the results by the followers field.
func ByFollowers(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowers, opts...).ToFunc()
}

// ByFollowing orders the results by the following field.
func ByFollowing(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFollowing, opts...).ToFunc()
}

// ByOutbox orders the results by the outbox field.
func ByOutbox(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutbox, opts...).ToFunc()
}

// ByAvatarImageField orders the results by avatarImage field.
func ByAvatarImageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAvatarImageStep(), sql.OrderByField(field, opts...))
	}
}

// ByHeaderImageField orders the results by headerImage field.
func ByHeaderImageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHeaderImageStep(), sql.OrderByField(field, opts...))
	}
}

// ByAuthoredNotesCount orders the results by authoredNotes count.
func ByAuthoredNotesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthoredNotesStep(), opts...)
	}
}

// ByAuthoredNotes orders the results by authoredNotes terms.
func ByAuthoredNotes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthoredNotesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMentionedNotesCount orders the results by mentionedNotes count.
func ByMentionedNotesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMentionedNotesStep(), opts...)
	}
}

// ByMentionedNotes orders the results by mentionedNotes terms.
func ByMentionedNotes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMentionedNotesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByServersCount orders the results by servers count.
func ByServersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newServersStep(), opts...)
	}
}

// ByServers orders the results by servers terms.
func ByServers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByModeratedServersCount orders the results by moderatedServers count.
func ByModeratedServersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newModeratedServersStep(), opts...)
	}
}

// ByModeratedServers orders the results by moderatedServers terms.
func ByModeratedServers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newModeratedServersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAdministeredServersCount orders the results by administeredServers count.
func ByAdministeredServersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAdministeredServersStep(), opts...)
	}
}

// ByAdministeredServers orders the results by administeredServers terms.
func ByAdministeredServers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAdministeredServersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAvatarImageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AvatarImageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AvatarImageTable, AvatarImageColumn),
	)
}
func newHeaderImageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HeaderImageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, HeaderImageTable, HeaderImageColumn),
	)
}
func newAuthoredNotesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthoredNotesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AuthoredNotesTable, AuthoredNotesColumn),
	)
}
func newMentionedNotesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MentionedNotesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, MentionedNotesTable, MentionedNotesPrimaryKey...),
	)
}
func newServersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ServersTable, ServersPrimaryKey...),
	)
}
func newModeratedServersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ModeratedServersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ModeratedServersTable, ModeratedServersPrimaryKey...),
	)
}
func newAdministeredServersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AdministeredServersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AdministeredServersTable, AdministeredServersPrimaryKey...),
	)
}
