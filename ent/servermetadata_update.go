// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/ent/predicate"
	"github.com/lysand-org/versia-go/ent/servermetadata"
	"github.com/lysand-org/versia-go/ent/user"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

// ServerMetadataUpdate is the builder for updating ServerMetadata entities.
type ServerMetadataUpdate struct {
	config
	hooks    []Hook
	mutation *ServerMetadataMutation
}

// Where appends a list predicates to the ServerMetadataUpdate builder.
func (smu *ServerMetadataUpdate) Where(ps ...predicate.ServerMetadata) *ServerMetadataUpdate {
	smu.mutation.Where(ps...)
	return smu
}

// SetIsRemote sets the "isRemote" field.
func (smu *ServerMetadataUpdate) SetIsRemote(b bool) *ServerMetadataUpdate {
	smu.mutation.SetIsRemote(b)
	return smu
}

// SetNillableIsRemote sets the "isRemote" field if the given value is not nil.
func (smu *ServerMetadataUpdate) SetNillableIsRemote(b *bool) *ServerMetadataUpdate {
	if b != nil {
		smu.SetIsRemote(*b)
	}
	return smu
}

// SetURI sets the "uri" field.
func (smu *ServerMetadataUpdate) SetURI(s string) *ServerMetadataUpdate {
	smu.mutation.SetURI(s)
	return smu
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (smu *ServerMetadataUpdate) SetNillableURI(s *string) *ServerMetadataUpdate {
	if s != nil {
		smu.SetURI(*s)
	}
	return smu
}

// SetExtensions sets the "extensions" field.
func (smu *ServerMetadataUpdate) SetExtensions(l lysand.Extensions) *ServerMetadataUpdate {
	smu.mutation.SetExtensions(l)
	return smu
}

// SetUpdatedAt sets the "updated_at" field.
func (smu *ServerMetadataUpdate) SetUpdatedAt(t time.Time) *ServerMetadataUpdate {
	smu.mutation.SetUpdatedAt(t)
	return smu
}

// SetName sets the "name" field.
func (smu *ServerMetadataUpdate) SetName(s string) *ServerMetadataUpdate {
	smu.mutation.SetName(s)
	return smu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (smu *ServerMetadataUpdate) SetNillableName(s *string) *ServerMetadataUpdate {
	if s != nil {
		smu.SetName(*s)
	}
	return smu
}

// SetDescription sets the "description" field.
func (smu *ServerMetadataUpdate) SetDescription(s string) *ServerMetadataUpdate {
	smu.mutation.SetDescription(s)
	return smu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (smu *ServerMetadataUpdate) SetNillableDescription(s *string) *ServerMetadataUpdate {
	if s != nil {
		smu.SetDescription(*s)
	}
	return smu
}

// ClearDescription clears the value of the "description" field.
func (smu *ServerMetadataUpdate) ClearDescription() *ServerMetadataUpdate {
	smu.mutation.ClearDescription()
	return smu
}

// SetVersion sets the "version" field.
func (smu *ServerMetadataUpdate) SetVersion(s string) *ServerMetadataUpdate {
	smu.mutation.SetVersion(s)
	return smu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (smu *ServerMetadataUpdate) SetNillableVersion(s *string) *ServerMetadataUpdate {
	if s != nil {
		smu.SetVersion(*s)
	}
	return smu
}

// SetSupportedExtensions sets the "supportedExtensions" field.
func (smu *ServerMetadataUpdate) SetSupportedExtensions(s []string) *ServerMetadataUpdate {
	smu.mutation.SetSupportedExtensions(s)
	return smu
}

// AppendSupportedExtensions appends s to the "supportedExtensions" field.
func (smu *ServerMetadataUpdate) AppendSupportedExtensions(s []string) *ServerMetadataUpdate {
	smu.mutation.AppendSupportedExtensions(s)
	return smu
}

// SetFollowerID sets the "follower" edge to the User entity by ID.
func (smu *ServerMetadataUpdate) SetFollowerID(id uuid.UUID) *ServerMetadataUpdate {
	smu.mutation.SetFollowerID(id)
	return smu
}

// SetFollower sets the "follower" edge to the User entity.
func (smu *ServerMetadataUpdate) SetFollower(u *User) *ServerMetadataUpdate {
	return smu.SetFollowerID(u.ID)
}

// SetFolloweeID sets the "followee" edge to the User entity by ID.
func (smu *ServerMetadataUpdate) SetFolloweeID(id uuid.UUID) *ServerMetadataUpdate {
	smu.mutation.SetFolloweeID(id)
	return smu
}

// SetFollowee sets the "followee" edge to the User entity.
func (smu *ServerMetadataUpdate) SetFollowee(u *User) *ServerMetadataUpdate {
	return smu.SetFolloweeID(u.ID)
}

// Mutation returns the ServerMetadataMutation object of the builder.
func (smu *ServerMetadataUpdate) Mutation() *ServerMetadataMutation {
	return smu.mutation
}

// ClearFollower clears the "follower" edge to the User entity.
func (smu *ServerMetadataUpdate) ClearFollower() *ServerMetadataUpdate {
	smu.mutation.ClearFollower()
	return smu
}

// ClearFollowee clears the "followee" edge to the User entity.
func (smu *ServerMetadataUpdate) ClearFollowee() *ServerMetadataUpdate {
	smu.mutation.ClearFollowee()
	return smu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (smu *ServerMetadataUpdate) Save(ctx context.Context) (int, error) {
	smu.defaults()
	return withHooks(ctx, smu.sqlSave, smu.mutation, smu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (smu *ServerMetadataUpdate) SaveX(ctx context.Context) int {
	affected, err := smu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (smu *ServerMetadataUpdate) Exec(ctx context.Context) error {
	_, err := smu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smu *ServerMetadataUpdate) ExecX(ctx context.Context) {
	if err := smu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smu *ServerMetadataUpdate) defaults() {
	if _, ok := smu.mutation.UpdatedAt(); !ok {
		v := servermetadata.UpdateDefaultUpdatedAt()
		smu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smu *ServerMetadataUpdate) check() error {
	if v, ok := smu.mutation.URI(); ok {
		if err := servermetadata.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "ServerMetadata.uri": %w`, err)}
		}
	}
	if v, ok := smu.mutation.Name(); ok {
		if err := servermetadata.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ServerMetadata.name": %w`, err)}
		}
	}
	if v, ok := smu.mutation.Version(); ok {
		if err := servermetadata.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`ent: validator failed for field "ServerMetadata.version": %w`, err)}
		}
	}
	if _, ok := smu.mutation.FollowerID(); smu.mutation.FollowerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ServerMetadata.follower"`)
	}
	if _, ok := smu.mutation.FolloweeID(); smu.mutation.FolloweeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ServerMetadata.followee"`)
	}
	return nil
}

func (smu *ServerMetadataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := smu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(servermetadata.Table, servermetadata.Columns, sqlgraph.NewFieldSpec(servermetadata.FieldID, field.TypeUUID))
	if ps := smu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smu.mutation.IsRemote(); ok {
		_spec.SetField(servermetadata.FieldIsRemote, field.TypeBool, value)
	}
	if value, ok := smu.mutation.URI(); ok {
		_spec.SetField(servermetadata.FieldURI, field.TypeString, value)
	}
	if value, ok := smu.mutation.Extensions(); ok {
		_spec.SetField(servermetadata.FieldExtensions, field.TypeJSON, value)
	}
	if value, ok := smu.mutation.UpdatedAt(); ok {
		_spec.SetField(servermetadata.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := smu.mutation.Name(); ok {
		_spec.SetField(servermetadata.FieldName, field.TypeString, value)
	}
	if value, ok := smu.mutation.Description(); ok {
		_spec.SetField(servermetadata.FieldDescription, field.TypeString, value)
	}
	if smu.mutation.DescriptionCleared() {
		_spec.ClearField(servermetadata.FieldDescription, field.TypeString)
	}
	if value, ok := smu.mutation.Version(); ok {
		_spec.SetField(servermetadata.FieldVersion, field.TypeString, value)
	}
	if value, ok := smu.mutation.SupportedExtensions(); ok {
		_spec.SetField(servermetadata.FieldSupportedExtensions, field.TypeJSON, value)
	}
	if value, ok := smu.mutation.AppendedSupportedExtensions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, servermetadata.FieldSupportedExtensions, value)
		})
	}
	if smu.mutation.FollowerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servermetadata.FollowerTable,
			Columns: []string{servermetadata.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smu.mutation.FollowerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servermetadata.FollowerTable,
			Columns: []string{servermetadata.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if smu.mutation.FolloweeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servermetadata.FolloweeTable,
			Columns: []string{servermetadata.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smu.mutation.FolloweeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servermetadata.FolloweeTable,
			Columns: []string{servermetadata.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, smu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servermetadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	smu.mutation.done = true
	return n, nil
}

// ServerMetadataUpdateOne is the builder for updating a single ServerMetadata entity.
type ServerMetadataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServerMetadataMutation
}

// SetIsRemote sets the "isRemote" field.
func (smuo *ServerMetadataUpdateOne) SetIsRemote(b bool) *ServerMetadataUpdateOne {
	smuo.mutation.SetIsRemote(b)
	return smuo
}

// SetNillableIsRemote sets the "isRemote" field if the given value is not nil.
func (smuo *ServerMetadataUpdateOne) SetNillableIsRemote(b *bool) *ServerMetadataUpdateOne {
	if b != nil {
		smuo.SetIsRemote(*b)
	}
	return smuo
}

// SetURI sets the "uri" field.
func (smuo *ServerMetadataUpdateOne) SetURI(s string) *ServerMetadataUpdateOne {
	smuo.mutation.SetURI(s)
	return smuo
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (smuo *ServerMetadataUpdateOne) SetNillableURI(s *string) *ServerMetadataUpdateOne {
	if s != nil {
		smuo.SetURI(*s)
	}
	return smuo
}

// SetExtensions sets the "extensions" field.
func (smuo *ServerMetadataUpdateOne) SetExtensions(l lysand.Extensions) *ServerMetadataUpdateOne {
	smuo.mutation.SetExtensions(l)
	return smuo
}

// SetUpdatedAt sets the "updated_at" field.
func (smuo *ServerMetadataUpdateOne) SetUpdatedAt(t time.Time) *ServerMetadataUpdateOne {
	smuo.mutation.SetUpdatedAt(t)
	return smuo
}

// SetName sets the "name" field.
func (smuo *ServerMetadataUpdateOne) SetName(s string) *ServerMetadataUpdateOne {
	smuo.mutation.SetName(s)
	return smuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (smuo *ServerMetadataUpdateOne) SetNillableName(s *string) *ServerMetadataUpdateOne {
	if s != nil {
		smuo.SetName(*s)
	}
	return smuo
}

// SetDescription sets the "description" field.
func (smuo *ServerMetadataUpdateOne) SetDescription(s string) *ServerMetadataUpdateOne {
	smuo.mutation.SetDescription(s)
	return smuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (smuo *ServerMetadataUpdateOne) SetNillableDescription(s *string) *ServerMetadataUpdateOne {
	if s != nil {
		smuo.SetDescription(*s)
	}
	return smuo
}

// ClearDescription clears the value of the "description" field.
func (smuo *ServerMetadataUpdateOne) ClearDescription() *ServerMetadataUpdateOne {
	smuo.mutation.ClearDescription()
	return smuo
}

// SetVersion sets the "version" field.
func (smuo *ServerMetadataUpdateOne) SetVersion(s string) *ServerMetadataUpdateOne {
	smuo.mutation.SetVersion(s)
	return smuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (smuo *ServerMetadataUpdateOne) SetNillableVersion(s *string) *ServerMetadataUpdateOne {
	if s != nil {
		smuo.SetVersion(*s)
	}
	return smuo
}

// SetSupportedExtensions sets the "supportedExtensions" field.
func (smuo *ServerMetadataUpdateOne) SetSupportedExtensions(s []string) *ServerMetadataUpdateOne {
	smuo.mutation.SetSupportedExtensions(s)
	return smuo
}

// AppendSupportedExtensions appends s to the "supportedExtensions" field.
func (smuo *ServerMetadataUpdateOne) AppendSupportedExtensions(s []string) *ServerMetadataUpdateOne {
	smuo.mutation.AppendSupportedExtensions(s)
	return smuo
}

// SetFollowerID sets the "follower" edge to the User entity by ID.
func (smuo *ServerMetadataUpdateOne) SetFollowerID(id uuid.UUID) *ServerMetadataUpdateOne {
	smuo.mutation.SetFollowerID(id)
	return smuo
}

// SetFollower sets the "follower" edge to the User entity.
func (smuo *ServerMetadataUpdateOne) SetFollower(u *User) *ServerMetadataUpdateOne {
	return smuo.SetFollowerID(u.ID)
}

// SetFolloweeID sets the "followee" edge to the User entity by ID.
func (smuo *ServerMetadataUpdateOne) SetFolloweeID(id uuid.UUID) *ServerMetadataUpdateOne {
	smuo.mutation.SetFolloweeID(id)
	return smuo
}

// SetFollowee sets the "followee" edge to the User entity.
func (smuo *ServerMetadataUpdateOne) SetFollowee(u *User) *ServerMetadataUpdateOne {
	return smuo.SetFolloweeID(u.ID)
}

// Mutation returns the ServerMetadataMutation object of the builder.
func (smuo *ServerMetadataUpdateOne) Mutation() *ServerMetadataMutation {
	return smuo.mutation
}

// ClearFollower clears the "follower" edge to the User entity.
func (smuo *ServerMetadataUpdateOne) ClearFollower() *ServerMetadataUpdateOne {
	smuo.mutation.ClearFollower()
	return smuo
}

// ClearFollowee clears the "followee" edge to the User entity.
func (smuo *ServerMetadataUpdateOne) ClearFollowee() *ServerMetadataUpdateOne {
	smuo.mutation.ClearFollowee()
	return smuo
}

// Where appends a list predicates to the ServerMetadataUpdate builder.
func (smuo *ServerMetadataUpdateOne) Where(ps ...predicate.ServerMetadata) *ServerMetadataUpdateOne {
	smuo.mutation.Where(ps...)
	return smuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (smuo *ServerMetadataUpdateOne) Select(field string, fields ...string) *ServerMetadataUpdateOne {
	smuo.fields = append([]string{field}, fields...)
	return smuo
}

// Save executes the query and returns the updated ServerMetadata entity.
func (smuo *ServerMetadataUpdateOne) Save(ctx context.Context) (*ServerMetadata, error) {
	smuo.defaults()
	return withHooks(ctx, smuo.sqlSave, smuo.mutation, smuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (smuo *ServerMetadataUpdateOne) SaveX(ctx context.Context) *ServerMetadata {
	node, err := smuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (smuo *ServerMetadataUpdateOne) Exec(ctx context.Context) error {
	_, err := smuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smuo *ServerMetadataUpdateOne) ExecX(ctx context.Context) {
	if err := smuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smuo *ServerMetadataUpdateOne) defaults() {
	if _, ok := smuo.mutation.UpdatedAt(); !ok {
		v := servermetadata.UpdateDefaultUpdatedAt()
		smuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smuo *ServerMetadataUpdateOne) check() error {
	if v, ok := smuo.mutation.URI(); ok {
		if err := servermetadata.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "ServerMetadata.uri": %w`, err)}
		}
	}
	if v, ok := smuo.mutation.Name(); ok {
		if err := servermetadata.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ServerMetadata.name": %w`, err)}
		}
	}
	if v, ok := smuo.mutation.Version(); ok {
		if err := servermetadata.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`ent: validator failed for field "ServerMetadata.version": %w`, err)}
		}
	}
	if _, ok := smuo.mutation.FollowerID(); smuo.mutation.FollowerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ServerMetadata.follower"`)
	}
	if _, ok := smuo.mutation.FolloweeID(); smuo.mutation.FolloweeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ServerMetadata.followee"`)
	}
	return nil
}

func (smuo *ServerMetadataUpdateOne) sqlSave(ctx context.Context) (_node *ServerMetadata, err error) {
	if err := smuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(servermetadata.Table, servermetadata.Columns, sqlgraph.NewFieldSpec(servermetadata.FieldID, field.TypeUUID))
	id, ok := smuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ServerMetadata.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := smuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servermetadata.FieldID)
		for _, f := range fields {
			if !servermetadata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != servermetadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := smuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smuo.mutation.IsRemote(); ok {
		_spec.SetField(servermetadata.FieldIsRemote, field.TypeBool, value)
	}
	if value, ok := smuo.mutation.URI(); ok {
		_spec.SetField(servermetadata.FieldURI, field.TypeString, value)
	}
	if value, ok := smuo.mutation.Extensions(); ok {
		_spec.SetField(servermetadata.FieldExtensions, field.TypeJSON, value)
	}
	if value, ok := smuo.mutation.UpdatedAt(); ok {
		_spec.SetField(servermetadata.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := smuo.mutation.Name(); ok {
		_spec.SetField(servermetadata.FieldName, field.TypeString, value)
	}
	if value, ok := smuo.mutation.Description(); ok {
		_spec.SetField(servermetadata.FieldDescription, field.TypeString, value)
	}
	if smuo.mutation.DescriptionCleared() {
		_spec.ClearField(servermetadata.FieldDescription, field.TypeString)
	}
	if value, ok := smuo.mutation.Version(); ok {
		_spec.SetField(servermetadata.FieldVersion, field.TypeString, value)
	}
	if value, ok := smuo.mutation.SupportedExtensions(); ok {
		_spec.SetField(servermetadata.FieldSupportedExtensions, field.TypeJSON, value)
	}
	if value, ok := smuo.mutation.AppendedSupportedExtensions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, servermetadata.FieldSupportedExtensions, value)
		})
	}
	if smuo.mutation.FollowerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servermetadata.FollowerTable,
			Columns: []string{servermetadata.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smuo.mutation.FollowerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servermetadata.FollowerTable,
			Columns: []string{servermetadata.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if smuo.mutation.FolloweeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servermetadata.FolloweeTable,
			Columns: []string{servermetadata.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := smuo.mutation.FolloweeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   servermetadata.FolloweeTable,
			Columns: []string{servermetadata.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ServerMetadata{config: smuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, smuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servermetadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	smuo.mutation.done = true
	return _node, nil
}
