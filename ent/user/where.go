// Code generated by ent, DO NOT EDIT.

package user

import (
	"crypto/ed25519"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IsRemote applies equality check predicate on the "isRemote" field. It's identical to IsRemoteEQ.
func IsRemote(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsRemote, v))
}

// URI applies equality check predicate on the "uri" field. It's identical to URIEQ.
func URI(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldURI, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// PasswordHash applies equality check predicate on the "passwordHash" field. It's identical to PasswordHashEQ.
func PasswordHash(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordHash, v))
}

// DisplayName applies equality check predicate on the "displayName" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// Biography applies equality check predicate on the "biography" field. It's identical to BiographyEQ.
func Biography(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBiography, v))
}

// PublicKey applies equality check predicate on the "publicKey" field. It's identical to PublicKeyEQ.
func PublicKey(v ed25519.PublicKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldEQ(FieldPublicKey, vc))
}

// PrivateKey applies equality check predicate on the "privateKey" field. It's identical to PrivateKeyEQ.
func PrivateKey(v ed25519.PrivateKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldEQ(FieldPrivateKey, vc))
}

// Indexable applies equality check predicate on the "indexable" field. It's identical to IndexableEQ.
func Indexable(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIndexable, v))
}

// Inbox applies equality check predicate on the "inbox" field. It's identical to InboxEQ.
func Inbox(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldInbox, v))
}

// Featured applies equality check predicate on the "featured" field. It's identical to FeaturedEQ.
func Featured(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFeatured, v))
}

// Followers applies equality check predicate on the "followers" field. It's identical to FollowersEQ.
func Followers(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFollowers, v))
}

// Following applies equality check predicate on the "following" field. It's identical to FollowingEQ.
func Following(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFollowing, v))
}

// Outbox applies equality check predicate on the "outbox" field. It's identical to OutboxEQ.
func Outbox(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOutbox, v))
}

// IsRemoteEQ applies the EQ predicate on the "isRemote" field.
func IsRemoteEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIsRemote, v))
}

// IsRemoteNEQ applies the NEQ predicate on the "isRemote" field.
func IsRemoteNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIsRemote, v))
}

// URIEQ applies the EQ predicate on the "uri" field.
func URIEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldURI, v))
}

// URINEQ applies the NEQ predicate on the "uri" field.
func URINEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldURI, v))
}

// URIIn applies the In predicate on the "uri" field.
func URIIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldURI, vs...))
}

// URINotIn applies the NotIn predicate on the "uri" field.
func URINotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldURI, vs...))
}

// URIGT applies the GT predicate on the "uri" field.
func URIGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldURI, v))
}

// URIGTE applies the GTE predicate on the "uri" field.
func URIGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldURI, v))
}

// URILT applies the LT predicate on the "uri" field.
func URILT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldURI, v))
}

// URILTE applies the LTE predicate on the "uri" field.
func URILTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldURI, v))
}

// URIContains applies the Contains predicate on the "uri" field.
func URIContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldURI, v))
}

// URIHasPrefix applies the HasPrefix predicate on the "uri" field.
func URIHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldURI, v))
}

// URIHasSuffix applies the HasSuffix predicate on the "uri" field.
func URIHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldURI, v))
}

// URIEqualFold applies the EqualFold predicate on the "uri" field.
func URIEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldURI, v))
}

// URIContainsFold applies the ContainsFold predicate on the "uri" field.
func URIContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldURI, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordHashEQ applies the EQ predicate on the "passwordHash" field.
func PasswordHashEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordHash, v))
}

// PasswordHashNEQ applies the NEQ predicate on the "passwordHash" field.
func PasswordHashNEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPasswordHash, v))
}

// PasswordHashIn applies the In predicate on the "passwordHash" field.
func PasswordHashIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldIn(FieldPasswordHash, vs...))
}

// PasswordHashNotIn applies the NotIn predicate on the "passwordHash" field.
func PasswordHashNotIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPasswordHash, vs...))
}

// PasswordHashGT applies the GT predicate on the "passwordHash" field.
func PasswordHashGT(v []byte) predicate.User {
	return predicate.User(sql.FieldGT(FieldPasswordHash, v))
}

// PasswordHashGTE applies the GTE predicate on the "passwordHash" field.
func PasswordHashGTE(v []byte) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPasswordHash, v))
}

// PasswordHashLT applies the LT predicate on the "passwordHash" field.
func PasswordHashLT(v []byte) predicate.User {
	return predicate.User(sql.FieldLT(FieldPasswordHash, v))
}

// PasswordHashLTE applies the LTE predicate on the "passwordHash" field.
func PasswordHashLTE(v []byte) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPasswordHash, v))
}

// PasswordHashIsNil applies the IsNil predicate on the "passwordHash" field.
func PasswordHashIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPasswordHash))
}

// PasswordHashNotNil applies the NotNil predicate on the "passwordHash" field.
func PasswordHashNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPasswordHash))
}

// DisplayNameEQ applies the EQ predicate on the "displayName" field.
func DisplayNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "displayName" field.
func DisplayNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "displayName" field.
func DisplayNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "displayName" field.
func DisplayNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "displayName" field.
func DisplayNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "displayName" field.
func DisplayNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "displayName" field.
func DisplayNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "displayName" field.
func DisplayNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "displayName" field.
func DisplayNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "displayName" field.
func DisplayNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "displayName" field.
func DisplayNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameIsNil applies the IsNil predicate on the "displayName" field.
func DisplayNameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldDisplayName))
}

// DisplayNameNotNil applies the NotNil predicate on the "displayName" field.
func DisplayNameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldDisplayName))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "displayName" field.
func DisplayNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "displayName" field.
func DisplayNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldDisplayName, v))
}

// BiographyEQ applies the EQ predicate on the "biography" field.
func BiographyEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBiography, v))
}

// BiographyNEQ applies the NEQ predicate on the "biography" field.
func BiographyNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBiography, v))
}

// BiographyIn applies the In predicate on the "biography" field.
func BiographyIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldBiography, vs...))
}

// BiographyNotIn applies the NotIn predicate on the "biography" field.
func BiographyNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBiography, vs...))
}

// BiographyGT applies the GT predicate on the "biography" field.
func BiographyGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldBiography, v))
}

// BiographyGTE applies the GTE predicate on the "biography" field.
func BiographyGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBiography, v))
}

// BiographyLT applies the LT predicate on the "biography" field.
func BiographyLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldBiography, v))
}

// BiographyLTE applies the LTE predicate on the "biography" field.
func BiographyLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBiography, v))
}

// BiographyContains applies the Contains predicate on the "biography" field.
func BiographyContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldBiography, v))
}

// BiographyHasPrefix applies the HasPrefix predicate on the "biography" field.
func BiographyHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldBiography, v))
}

// BiographyHasSuffix applies the HasSuffix predicate on the "biography" field.
func BiographyHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldBiography, v))
}

// BiographyIsNil applies the IsNil predicate on the "biography" field.
func BiographyIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldBiography))
}

// BiographyNotNil applies the NotNil predicate on the "biography" field.
func BiographyNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldBiography))
}

// BiographyEqualFold applies the EqualFold predicate on the "biography" field.
func BiographyEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldBiography, v))
}

// BiographyContainsFold applies the ContainsFold predicate on the "biography" field.
func BiographyContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldBiography, v))
}

// PublicKeyEQ applies the EQ predicate on the "publicKey" field.
func PublicKeyEQ(v ed25519.PublicKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldEQ(FieldPublicKey, vc))
}

// PublicKeyNEQ applies the NEQ predicate on the "publicKey" field.
func PublicKeyNEQ(v ed25519.PublicKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldNEQ(FieldPublicKey, vc))
}

// PublicKeyIn applies the In predicate on the "publicKey" field.
func PublicKeyIn(vs ...ed25519.PublicKey) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = []byte(vs[i])
	}
	return predicate.User(sql.FieldIn(FieldPublicKey, v...))
}

// PublicKeyNotIn applies the NotIn predicate on the "publicKey" field.
func PublicKeyNotIn(vs ...ed25519.PublicKey) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = []byte(vs[i])
	}
	return predicate.User(sql.FieldNotIn(FieldPublicKey, v...))
}

// PublicKeyGT applies the GT predicate on the "publicKey" field.
func PublicKeyGT(v ed25519.PublicKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldGT(FieldPublicKey, vc))
}

// PublicKeyGTE applies the GTE predicate on the "publicKey" field.
func PublicKeyGTE(v ed25519.PublicKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldGTE(FieldPublicKey, vc))
}

// PublicKeyLT applies the LT predicate on the "publicKey" field.
func PublicKeyLT(v ed25519.PublicKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldLT(FieldPublicKey, vc))
}

// PublicKeyLTE applies the LTE predicate on the "publicKey" field.
func PublicKeyLTE(v ed25519.PublicKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldLTE(FieldPublicKey, vc))
}

// PrivateKeyEQ applies the EQ predicate on the "privateKey" field.
func PrivateKeyEQ(v ed25519.PrivateKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldEQ(FieldPrivateKey, vc))
}

// PrivateKeyNEQ applies the NEQ predicate on the "privateKey" field.
func PrivateKeyNEQ(v ed25519.PrivateKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldNEQ(FieldPrivateKey, vc))
}

// PrivateKeyIn applies the In predicate on the "privateKey" field.
func PrivateKeyIn(vs ...ed25519.PrivateKey) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = []byte(vs[i])
	}
	return predicate.User(sql.FieldIn(FieldPrivateKey, v...))
}

// PrivateKeyNotIn applies the NotIn predicate on the "privateKey" field.
func PrivateKeyNotIn(vs ...ed25519.PrivateKey) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = []byte(vs[i])
	}
	return predicate.User(sql.FieldNotIn(FieldPrivateKey, v...))
}

// PrivateKeyGT applies the GT predicate on the "privateKey" field.
func PrivateKeyGT(v ed25519.PrivateKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldGT(FieldPrivateKey, vc))
}

// PrivateKeyGTE applies the GTE predicate on the "privateKey" field.
func PrivateKeyGTE(v ed25519.PrivateKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldGTE(FieldPrivateKey, vc))
}

// PrivateKeyLT applies the LT predicate on the "privateKey" field.
func PrivateKeyLT(v ed25519.PrivateKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldLT(FieldPrivateKey, vc))
}

// PrivateKeyLTE applies the LTE predicate on the "privateKey" field.
func PrivateKeyLTE(v ed25519.PrivateKey) predicate.User {
	vc := []byte(v)
	return predicate.User(sql.FieldLTE(FieldPrivateKey, vc))
}

// PrivateKeyIsNil applies the IsNil predicate on the "privateKey" field.
func PrivateKeyIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPrivateKey))
}

// PrivateKeyNotNil applies the NotNil predicate on the "privateKey" field.
func PrivateKeyNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPrivateKey))
}

// IndexableEQ applies the EQ predicate on the "indexable" field.
func IndexableEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIndexable, v))
}

// IndexableNEQ applies the NEQ predicate on the "indexable" field.
func IndexableNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIndexable, v))
}

// PrivacyLevelEQ applies the EQ predicate on the "privacyLevel" field.
func PrivacyLevelEQ(v PrivacyLevel) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPrivacyLevel, v))
}

// PrivacyLevelNEQ applies the NEQ predicate on the "privacyLevel" field.
func PrivacyLevelNEQ(v PrivacyLevel) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPrivacyLevel, v))
}

// PrivacyLevelIn applies the In predicate on the "privacyLevel" field.
func PrivacyLevelIn(vs ...PrivacyLevel) predicate.User {
	return predicate.User(sql.FieldIn(FieldPrivacyLevel, vs...))
}

// PrivacyLevelNotIn applies the NotIn predicate on the "privacyLevel" field.
func PrivacyLevelNotIn(vs ...PrivacyLevel) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPrivacyLevel, vs...))
}

// InboxEQ applies the EQ predicate on the "inbox" field.
func InboxEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldInbox, v))
}

// InboxNEQ applies the NEQ predicate on the "inbox" field.
func InboxNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldInbox, v))
}

// InboxIn applies the In predicate on the "inbox" field.
func InboxIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldInbox, vs...))
}

// InboxNotIn applies the NotIn predicate on the "inbox" field.
func InboxNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldInbox, vs...))
}

// InboxGT applies the GT predicate on the "inbox" field.
func InboxGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldInbox, v))
}

// InboxGTE applies the GTE predicate on the "inbox" field.
func InboxGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldInbox, v))
}

// InboxLT applies the LT predicate on the "inbox" field.
func InboxLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldInbox, v))
}

// InboxLTE applies the LTE predicate on the "inbox" field.
func InboxLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldInbox, v))
}

// InboxContains applies the Contains predicate on the "inbox" field.
func InboxContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldInbox, v))
}

// InboxHasPrefix applies the HasPrefix predicate on the "inbox" field.
func InboxHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldInbox, v))
}

// InboxHasSuffix applies the HasSuffix predicate on the "inbox" field.
func InboxHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldInbox, v))
}

// InboxEqualFold applies the EqualFold predicate on the "inbox" field.
func InboxEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldInbox, v))
}

// InboxContainsFold applies the ContainsFold predicate on the "inbox" field.
func InboxContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldInbox, v))
}

// FeaturedEQ applies the EQ predicate on the "featured" field.
func FeaturedEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFeatured, v))
}

// FeaturedNEQ applies the NEQ predicate on the "featured" field.
func FeaturedNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFeatured, v))
}

// FeaturedIn applies the In predicate on the "featured" field.
func FeaturedIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFeatured, vs...))
}

// FeaturedNotIn applies the NotIn predicate on the "featured" field.
func FeaturedNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFeatured, vs...))
}

// FeaturedGT applies the GT predicate on the "featured" field.
func FeaturedGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFeatured, v))
}

// FeaturedGTE applies the GTE predicate on the "featured" field.
func FeaturedGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFeatured, v))
}

// FeaturedLT applies the LT predicate on the "featured" field.
func FeaturedLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFeatured, v))
}

// FeaturedLTE applies the LTE predicate on the "featured" field.
func FeaturedLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFeatured, v))
}

// FeaturedContains applies the Contains predicate on the "featured" field.
func FeaturedContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFeatured, v))
}

// FeaturedHasPrefix applies the HasPrefix predicate on the "featured" field.
func FeaturedHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFeatured, v))
}

// FeaturedHasSuffix applies the HasSuffix predicate on the "featured" field.
func FeaturedHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFeatured, v))
}

// FeaturedEqualFold applies the EqualFold predicate on the "featured" field.
func FeaturedEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFeatured, v))
}

// FeaturedContainsFold applies the ContainsFold predicate on the "featured" field.
func FeaturedContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFeatured, v))
}

// FollowersEQ applies the EQ predicate on the "followers" field.
func FollowersEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFollowers, v))
}

// FollowersNEQ applies the NEQ predicate on the "followers" field.
func FollowersNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFollowers, v))
}

// FollowersIn applies the In predicate on the "followers" field.
func FollowersIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFollowers, vs...))
}

// FollowersNotIn applies the NotIn predicate on the "followers" field.
func FollowersNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFollowers, vs...))
}

// FollowersGT applies the GT predicate on the "followers" field.
func FollowersGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFollowers, v))
}

// FollowersGTE applies the GTE predicate on the "followers" field.
func FollowersGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFollowers, v))
}

// FollowersLT applies the LT predicate on the "followers" field.
func FollowersLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFollowers, v))
}

// FollowersLTE applies the LTE predicate on the "followers" field.
func FollowersLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFollowers, v))
}

// FollowersContains applies the Contains predicate on the "followers" field.
func FollowersContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFollowers, v))
}

// FollowersHasPrefix applies the HasPrefix predicate on the "followers" field.
func FollowersHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFollowers, v))
}

// FollowersHasSuffix applies the HasSuffix predicate on the "followers" field.
func FollowersHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFollowers, v))
}

// FollowersEqualFold applies the EqualFold predicate on the "followers" field.
func FollowersEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFollowers, v))
}

// FollowersContainsFold applies the ContainsFold predicate on the "followers" field.
func FollowersContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFollowers, v))
}

// FollowingEQ applies the EQ predicate on the "following" field.
func FollowingEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFollowing, v))
}

// FollowingNEQ applies the NEQ predicate on the "following" field.
func FollowingNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFollowing, v))
}

// FollowingIn applies the In predicate on the "following" field.
func FollowingIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFollowing, vs...))
}

// FollowingNotIn applies the NotIn predicate on the "following" field.
func FollowingNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFollowing, vs...))
}

// FollowingGT applies the GT predicate on the "following" field.
func FollowingGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFollowing, v))
}

// FollowingGTE applies the GTE predicate on the "following" field.
func FollowingGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFollowing, v))
}

// FollowingLT applies the LT predicate on the "following" field.
func FollowingLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFollowing, v))
}

// FollowingLTE applies the LTE predicate on the "following" field.
func FollowingLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFollowing, v))
}

// FollowingContains applies the Contains predicate on the "following" field.
func FollowingContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFollowing, v))
}

// FollowingHasPrefix applies the HasPrefix predicate on the "following" field.
func FollowingHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFollowing, v))
}

// FollowingHasSuffix applies the HasSuffix predicate on the "following" field.
func FollowingHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFollowing, v))
}

// FollowingEqualFold applies the EqualFold predicate on the "following" field.
func FollowingEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFollowing, v))
}

// FollowingContainsFold applies the ContainsFold predicate on the "following" field.
func FollowingContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFollowing, v))
}

// OutboxEQ applies the EQ predicate on the "outbox" field.
func OutboxEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOutbox, v))
}

// OutboxNEQ applies the NEQ predicate on the "outbox" field.
func OutboxNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOutbox, v))
}

// OutboxIn applies the In predicate on the "outbox" field.
func OutboxIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldOutbox, vs...))
}

// OutboxNotIn applies the NotIn predicate on the "outbox" field.
func OutboxNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldOutbox, vs...))
}

// OutboxGT applies the GT predicate on the "outbox" field.
func OutboxGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldOutbox, v))
}

// OutboxGTE applies the GTE predicate on the "outbox" field.
func OutboxGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldOutbox, v))
}

// OutboxLT applies the LT predicate on the "outbox" field.
func OutboxLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldOutbox, v))
}

// OutboxLTE applies the LTE predicate on the "outbox" field.
func OutboxLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldOutbox, v))
}

// OutboxContains applies the Contains predicate on the "outbox" field.
func OutboxContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldOutbox, v))
}

// OutboxHasPrefix applies the HasPrefix predicate on the "outbox" field.
func OutboxHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldOutbox, v))
}

// OutboxHasSuffix applies the HasSuffix predicate on the "outbox" field.
func OutboxHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldOutbox, v))
}

// OutboxEqualFold applies the EqualFold predicate on the "outbox" field.
func OutboxEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldOutbox, v))
}

// OutboxContainsFold applies the ContainsFold predicate on the "outbox" field.
func OutboxContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldOutbox, v))
}

// HasAvatarImage applies the HasEdge predicate on the "avatarImage" edge.
func HasAvatarImage() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AvatarImageTable, AvatarImageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAvatarImageWith applies the HasEdge predicate on the "avatarImage" edge with a given conditions (other predicates).
func HasAvatarImageWith(preds ...predicate.Image) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newAvatarImageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHeaderImage applies the HasEdge predicate on the "headerImage" edge.
func HasHeaderImage() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HeaderImageTable, HeaderImageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHeaderImageWith applies the HasEdge predicate on the "headerImage" edge with a given conditions (other predicates).
func HasHeaderImageWith(preds ...predicate.Image) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newHeaderImageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAuthoredNotes applies the HasEdge predicate on the "authoredNotes" edge.
func HasAuthoredNotes() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AuthoredNotesTable, AuthoredNotesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthoredNotesWith applies the HasEdge predicate on the "authoredNotes" edge with a given conditions (other predicates).
func HasAuthoredNotesWith(preds ...predicate.Note) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newAuthoredNotesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMentionedNotes applies the HasEdge predicate on the "mentionedNotes" edge.
func HasMentionedNotes() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MentionedNotesTable, MentionedNotesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMentionedNotesWith applies the HasEdge predicate on the "mentionedNotes" edge with a given conditions (other predicates).
func HasMentionedNotesWith(preds ...predicate.Note) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newMentionedNotesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
