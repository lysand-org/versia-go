// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/ent/follow"
	"github.com/lysand-org/versia-go/ent/predicate"
	"github.com/lysand-org/versia-go/ent/user"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

// FollowUpdate is the builder for updating Follow entities.
type FollowUpdate struct {
	config
	hooks    []Hook
	mutation *FollowMutation
}

// Where appends a list predicates to the FollowUpdate builder.
func (fu *FollowUpdate) Where(ps ...predicate.Follow) *FollowUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetIsRemote sets the "isRemote" field.
func (fu *FollowUpdate) SetIsRemote(b bool) *FollowUpdate {
	fu.mutation.SetIsRemote(b)
	return fu
}

// SetNillableIsRemote sets the "isRemote" field if the given value is not nil.
func (fu *FollowUpdate) SetNillableIsRemote(b *bool) *FollowUpdate {
	if b != nil {
		fu.SetIsRemote(*b)
	}
	return fu
}

// SetURI sets the "uri" field.
func (fu *FollowUpdate) SetURI(s string) *FollowUpdate {
	fu.mutation.SetURI(s)
	return fu
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (fu *FollowUpdate) SetNillableURI(s *string) *FollowUpdate {
	if s != nil {
		fu.SetURI(*s)
	}
	return fu
}

// SetExtensions sets the "extensions" field.
func (fu *FollowUpdate) SetExtensions(l lysand.Extensions) *FollowUpdate {
	fu.mutation.SetExtensions(l)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FollowUpdate) SetUpdatedAt(t time.Time) *FollowUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetStatus sets the "status" field.
func (fu *FollowUpdate) SetStatus(f follow.Status) *FollowUpdate {
	fu.mutation.SetStatus(f)
	return fu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fu *FollowUpdate) SetNillableStatus(f *follow.Status) *FollowUpdate {
	if f != nil {
		fu.SetStatus(*f)
	}
	return fu
}

// SetFollowerID sets the "follower" edge to the User entity by ID.
func (fu *FollowUpdate) SetFollowerID(id uuid.UUID) *FollowUpdate {
	fu.mutation.SetFollowerID(id)
	return fu
}

// SetFollower sets the "follower" edge to the User entity.
func (fu *FollowUpdate) SetFollower(u *User) *FollowUpdate {
	return fu.SetFollowerID(u.ID)
}

// SetFolloweeID sets the "followee" edge to the User entity by ID.
func (fu *FollowUpdate) SetFolloweeID(id uuid.UUID) *FollowUpdate {
	fu.mutation.SetFolloweeID(id)
	return fu
}

// SetFollowee sets the "followee" edge to the User entity.
func (fu *FollowUpdate) SetFollowee(u *User) *FollowUpdate {
	return fu.SetFolloweeID(u.ID)
}

// Mutation returns the FollowMutation object of the builder.
func (fu *FollowUpdate) Mutation() *FollowMutation {
	return fu.mutation
}

// ClearFollower clears the "follower" edge to the User entity.
func (fu *FollowUpdate) ClearFollower() *FollowUpdate {
	fu.mutation.ClearFollower()
	return fu
}

// ClearFollowee clears the "followee" edge to the User entity.
func (fu *FollowUpdate) ClearFollowee() *FollowUpdate {
	fu.mutation.ClearFollowee()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FollowUpdate) Save(ctx context.Context) (int, error) {
	fu.defaults()
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FollowUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FollowUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FollowUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FollowUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := follow.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FollowUpdate) check() error {
	if v, ok := fu.mutation.URI(); ok {
		if err := follow.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Follow.uri": %w`, err)}
		}
	}
	if v, ok := fu.mutation.Status(); ok {
		if err := follow.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Follow.status": %w`, err)}
		}
	}
	if _, ok := fu.mutation.FollowerID(); fu.mutation.FollowerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Follow.follower"`)
	}
	if _, ok := fu.mutation.FolloweeID(); fu.mutation.FolloweeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Follow.followee"`)
	}
	return nil
}

func (fu *FollowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(follow.Table, follow.Columns, sqlgraph.NewFieldSpec(follow.FieldID, field.TypeUUID))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.IsRemote(); ok {
		_spec.SetField(follow.FieldIsRemote, field.TypeBool, value)
	}
	if value, ok := fu.mutation.URI(); ok {
		_spec.SetField(follow.FieldURI, field.TypeString, value)
	}
	if value, ok := fu.mutation.Extensions(); ok {
		_spec.SetField(follow.FieldExtensions, field.TypeJSON, value)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(follow.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fu.mutation.Status(); ok {
		_spec.SetField(follow.FieldStatus, field.TypeEnum, value)
	}
	if fu.mutation.FollowerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FollowerTable,
			Columns: []string{follow.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FollowerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FollowerTable,
			Columns: []string{follow.FollowerColumn},
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
	if fu.mutation.FolloweeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FolloweeTable,
			Columns: []string{follow.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FolloweeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FolloweeTable,
			Columns: []string{follow.FolloweeColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{follow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FollowUpdateOne is the builder for updating a single Follow entity.
type FollowUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FollowMutation
}

// SetIsRemote sets the "isRemote" field.
func (fuo *FollowUpdateOne) SetIsRemote(b bool) *FollowUpdateOne {
	fuo.mutation.SetIsRemote(b)
	return fuo
}

// SetNillableIsRemote sets the "isRemote" field if the given value is not nil.
func (fuo *FollowUpdateOne) SetNillableIsRemote(b *bool) *FollowUpdateOne {
	if b != nil {
		fuo.SetIsRemote(*b)
	}
	return fuo
}

// SetURI sets the "uri" field.
func (fuo *FollowUpdateOne) SetURI(s string) *FollowUpdateOne {
	fuo.mutation.SetURI(s)
	return fuo
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (fuo *FollowUpdateOne) SetNillableURI(s *string) *FollowUpdateOne {
	if s != nil {
		fuo.SetURI(*s)
	}
	return fuo
}

// SetExtensions sets the "extensions" field.
func (fuo *FollowUpdateOne) SetExtensions(l lysand.Extensions) *FollowUpdateOne {
	fuo.mutation.SetExtensions(l)
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FollowUpdateOne) SetUpdatedAt(t time.Time) *FollowUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetStatus sets the "status" field.
func (fuo *FollowUpdateOne) SetStatus(f follow.Status) *FollowUpdateOne {
	fuo.mutation.SetStatus(f)
	return fuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fuo *FollowUpdateOne) SetNillableStatus(f *follow.Status) *FollowUpdateOne {
	if f != nil {
		fuo.SetStatus(*f)
	}
	return fuo
}

// SetFollowerID sets the "follower" edge to the User entity by ID.
func (fuo *FollowUpdateOne) SetFollowerID(id uuid.UUID) *FollowUpdateOne {
	fuo.mutation.SetFollowerID(id)
	return fuo
}

// SetFollower sets the "follower" edge to the User entity.
func (fuo *FollowUpdateOne) SetFollower(u *User) *FollowUpdateOne {
	return fuo.SetFollowerID(u.ID)
}

// SetFolloweeID sets the "followee" edge to the User entity by ID.
func (fuo *FollowUpdateOne) SetFolloweeID(id uuid.UUID) *FollowUpdateOne {
	fuo.mutation.SetFolloweeID(id)
	return fuo
}

// SetFollowee sets the "followee" edge to the User entity.
func (fuo *FollowUpdateOne) SetFollowee(u *User) *FollowUpdateOne {
	return fuo.SetFolloweeID(u.ID)
}

// Mutation returns the FollowMutation object of the builder.
func (fuo *FollowUpdateOne) Mutation() *FollowMutation {
	return fuo.mutation
}

// ClearFollower clears the "follower" edge to the User entity.
func (fuo *FollowUpdateOne) ClearFollower() *FollowUpdateOne {
	fuo.mutation.ClearFollower()
	return fuo
}

// ClearFollowee clears the "followee" edge to the User entity.
func (fuo *FollowUpdateOne) ClearFollowee() *FollowUpdateOne {
	fuo.mutation.ClearFollowee()
	return fuo
}

// Where appends a list predicates to the FollowUpdate builder.
func (fuo *FollowUpdateOne) Where(ps ...predicate.Follow) *FollowUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FollowUpdateOne) Select(field string, fields ...string) *FollowUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Follow entity.
func (fuo *FollowUpdateOne) Save(ctx context.Context) (*Follow, error) {
	fuo.defaults()
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FollowUpdateOne) SaveX(ctx context.Context) *Follow {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FollowUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FollowUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FollowUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := follow.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FollowUpdateOne) check() error {
	if v, ok := fuo.mutation.URI(); ok {
		if err := follow.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Follow.uri": %w`, err)}
		}
	}
	if v, ok := fuo.mutation.Status(); ok {
		if err := follow.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Follow.status": %w`, err)}
		}
	}
	if _, ok := fuo.mutation.FollowerID(); fuo.mutation.FollowerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Follow.follower"`)
	}
	if _, ok := fuo.mutation.FolloweeID(); fuo.mutation.FolloweeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Follow.followee"`)
	}
	return nil
}

func (fuo *FollowUpdateOne) sqlSave(ctx context.Context) (_node *Follow, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(follow.Table, follow.Columns, sqlgraph.NewFieldSpec(follow.FieldID, field.TypeUUID))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Follow.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, follow.FieldID)
		for _, f := range fields {
			if !follow.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != follow.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.IsRemote(); ok {
		_spec.SetField(follow.FieldIsRemote, field.TypeBool, value)
	}
	if value, ok := fuo.mutation.URI(); ok {
		_spec.SetField(follow.FieldURI, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Extensions(); ok {
		_spec.SetField(follow.FieldExtensions, field.TypeJSON, value)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(follow.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fuo.mutation.Status(); ok {
		_spec.SetField(follow.FieldStatus, field.TypeEnum, value)
	}
	if fuo.mutation.FollowerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FollowerTable,
			Columns: []string{follow.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FollowerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FollowerTable,
			Columns: []string{follow.FollowerColumn},
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
	if fuo.mutation.FolloweeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FolloweeTable,
			Columns: []string{follow.FolloweeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FolloweeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   follow.FolloweeTable,
			Columns: []string{follow.FolloweeColumn},
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
	_node = &Follow{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{follow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
