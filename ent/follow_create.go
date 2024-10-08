// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/ent/follow"
	"github.com/versia-pub/versia-go/ent/user"
	"github.com/versia-pub/versia-go/pkg/versia"
)

// FollowCreate is the builder for creating a Follow entity.
type FollowCreate struct {
	config
	mutation *FollowMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsRemote sets the "isRemote" field.
func (fc *FollowCreate) SetIsRemote(b bool) *FollowCreate {
	fc.mutation.SetIsRemote(b)
	return fc
}

// SetURI sets the "uri" field.
func (fc *FollowCreate) SetURI(s string) *FollowCreate {
	fc.mutation.SetURI(s)
	return fc
}

// SetExtensions sets the "extensions" field.
func (fc *FollowCreate) SetExtensions(v versia.Extensions) *FollowCreate {
	fc.mutation.SetExtensions(v)
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FollowCreate) SetCreatedAt(t time.Time) *FollowCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FollowCreate) SetNillableCreatedAt(t *time.Time) *FollowCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FollowCreate) SetUpdatedAt(t time.Time) *FollowCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FollowCreate) SetNillableUpdatedAt(t *time.Time) *FollowCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetStatus sets the "status" field.
func (fc *FollowCreate) SetStatus(f follow.Status) *FollowCreate {
	fc.mutation.SetStatus(f)
	return fc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fc *FollowCreate) SetNillableStatus(f *follow.Status) *FollowCreate {
	if f != nil {
		fc.SetStatus(*f)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FollowCreate) SetID(u uuid.UUID) *FollowCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FollowCreate) SetNillableID(u *uuid.UUID) *FollowCreate {
	if u != nil {
		fc.SetID(*u)
	}
	return fc
}

// SetFollowerID sets the "follower" edge to the User entity by ID.
func (fc *FollowCreate) SetFollowerID(id uuid.UUID) *FollowCreate {
	fc.mutation.SetFollowerID(id)
	return fc
}

// SetFollower sets the "follower" edge to the User entity.
func (fc *FollowCreate) SetFollower(u *User) *FollowCreate {
	return fc.SetFollowerID(u.ID)
}

// SetFolloweeID sets the "followee" edge to the User entity by ID.
func (fc *FollowCreate) SetFolloweeID(id uuid.UUID) *FollowCreate {
	fc.mutation.SetFolloweeID(id)
	return fc
}

// SetFollowee sets the "followee" edge to the User entity.
func (fc *FollowCreate) SetFollowee(u *User) *FollowCreate {
	return fc.SetFolloweeID(u.ID)
}

// Mutation returns the FollowMutation object of the builder.
func (fc *FollowCreate) Mutation() *FollowMutation {
	return fc.mutation
}

// Save creates the Follow in the database.
func (fc *FollowCreate) Save(ctx context.Context) (*Follow, error) {
	fc.defaults()
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FollowCreate) SaveX(ctx context.Context) *Follow {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FollowCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FollowCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FollowCreate) defaults() {
	if _, ok := fc.mutation.Extensions(); !ok {
		v := follow.DefaultExtensions
		fc.mutation.SetExtensions(v)
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := follow.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := follow.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.Status(); !ok {
		v := follow.DefaultStatus
		fc.mutation.SetStatus(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := follow.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FollowCreate) check() error {
	if _, ok := fc.mutation.IsRemote(); !ok {
		return &ValidationError{Name: "isRemote", err: errors.New(`ent: missing required field "Follow.isRemote"`)}
	}
	if _, ok := fc.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New(`ent: missing required field "Follow.uri"`)}
	}
	if v, ok := fc.mutation.URI(); ok {
		if err := follow.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Follow.uri": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Extensions(); !ok {
		return &ValidationError{Name: "extensions", err: errors.New(`ent: missing required field "Follow.extensions"`)}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Follow.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Follow.updated_at"`)}
	}
	if _, ok := fc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Follow.status"`)}
	}
	if v, ok := fc.mutation.Status(); ok {
		if err := follow.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Follow.status": %w`, err)}
		}
	}
	if _, ok := fc.mutation.FollowerID(); !ok {
		return &ValidationError{Name: "follower", err: errors.New(`ent: missing required edge "Follow.follower"`)}
	}
	if _, ok := fc.mutation.FolloweeID(); !ok {
		return &ValidationError{Name: "followee", err: errors.New(`ent: missing required edge "Follow.followee"`)}
	}
	return nil
}

func (fc *FollowCreate) sqlSave(ctx context.Context) (*Follow, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FollowCreate) createSpec() (*Follow, *sqlgraph.CreateSpec) {
	var (
		_node = &Follow{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(follow.Table, sqlgraph.NewFieldSpec(follow.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = fc.conflict
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.IsRemote(); ok {
		_spec.SetField(follow.FieldIsRemote, field.TypeBool, value)
		_node.IsRemote = value
	}
	if value, ok := fc.mutation.URI(); ok {
		_spec.SetField(follow.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	if value, ok := fc.mutation.Extensions(); ok {
		_spec.SetField(follow.FieldExtensions, field.TypeJSON, value)
		_node.Extensions = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(follow.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(follow.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.Status(); ok {
		_spec.SetField(follow.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := fc.mutation.FollowerIDs(); len(nodes) > 0 {
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
		_node.follow_follower = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FolloweeIDs(); len(nodes) > 0 {
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
		_node.follow_followee = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Follow.Create().
//		SetIsRemote(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FollowUpsert) {
//			SetIsRemote(v+v).
//		}).
//		Exec(ctx)
func (fc *FollowCreate) OnConflict(opts ...sql.ConflictOption) *FollowUpsertOne {
	fc.conflict = opts
	return &FollowUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Follow.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FollowCreate) OnConflictColumns(columns ...string) *FollowUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FollowUpsertOne{
		create: fc,
	}
}

type (
	// FollowUpsertOne is the builder for "upsert"-ing
	//  one Follow node.
	FollowUpsertOne struct {
		create *FollowCreate
	}

	// FollowUpsert is the "OnConflict" setter.
	FollowUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsRemote sets the "isRemote" field.
func (u *FollowUpsert) SetIsRemote(v bool) *FollowUpsert {
	u.Set(follow.FieldIsRemote, v)
	return u
}

// UpdateIsRemote sets the "isRemote" field to the value that was provided on create.
func (u *FollowUpsert) UpdateIsRemote() *FollowUpsert {
	u.SetExcluded(follow.FieldIsRemote)
	return u
}

// SetURI sets the "uri" field.
func (u *FollowUpsert) SetURI(v string) *FollowUpsert {
	u.Set(follow.FieldURI, v)
	return u
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *FollowUpsert) UpdateURI() *FollowUpsert {
	u.SetExcluded(follow.FieldURI)
	return u
}

// SetExtensions sets the "extensions" field.
func (u *FollowUpsert) SetExtensions(v versia.Extensions) *FollowUpsert {
	u.Set(follow.FieldExtensions, v)
	return u
}

// UpdateExtensions sets the "extensions" field to the value that was provided on create.
func (u *FollowUpsert) UpdateExtensions() *FollowUpsert {
	u.SetExcluded(follow.FieldExtensions)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FollowUpsert) SetUpdatedAt(v time.Time) *FollowUpsert {
	u.Set(follow.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FollowUpsert) UpdateUpdatedAt() *FollowUpsert {
	u.SetExcluded(follow.FieldUpdatedAt)
	return u
}

// SetStatus sets the "status" field.
func (u *FollowUpsert) SetStatus(v follow.Status) *FollowUpsert {
	u.Set(follow.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FollowUpsert) UpdateStatus() *FollowUpsert {
	u.SetExcluded(follow.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Follow.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(follow.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FollowUpsertOne) UpdateNewValues() *FollowUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(follow.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(follow.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Follow.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FollowUpsertOne) Ignore() *FollowUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FollowUpsertOne) DoNothing() *FollowUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FollowCreate.OnConflict
// documentation for more info.
func (u *FollowUpsertOne) Update(set func(*FollowUpsert)) *FollowUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FollowUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsRemote sets the "isRemote" field.
func (u *FollowUpsertOne) SetIsRemote(v bool) *FollowUpsertOne {
	return u.Update(func(s *FollowUpsert) {
		s.SetIsRemote(v)
	})
}

// UpdateIsRemote sets the "isRemote" field to the value that was provided on create.
func (u *FollowUpsertOne) UpdateIsRemote() *FollowUpsertOne {
	return u.Update(func(s *FollowUpsert) {
		s.UpdateIsRemote()
	})
}

// SetURI sets the "uri" field.
func (u *FollowUpsertOne) SetURI(v string) *FollowUpsertOne {
	return u.Update(func(s *FollowUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *FollowUpsertOne) UpdateURI() *FollowUpsertOne {
	return u.Update(func(s *FollowUpsert) {
		s.UpdateURI()
	})
}

// SetExtensions sets the "extensions" field.
func (u *FollowUpsertOne) SetExtensions(v versia.Extensions) *FollowUpsertOne {
	return u.Update(func(s *FollowUpsert) {
		s.SetExtensions(v)
	})
}

// UpdateExtensions sets the "extensions" field to the value that was provided on create.
func (u *FollowUpsertOne) UpdateExtensions() *FollowUpsertOne {
	return u.Update(func(s *FollowUpsert) {
		s.UpdateExtensions()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FollowUpsertOne) SetUpdatedAt(v time.Time) *FollowUpsertOne {
	return u.Update(func(s *FollowUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FollowUpsertOne) UpdateUpdatedAt() *FollowUpsertOne {
	return u.Update(func(s *FollowUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetStatus sets the "status" field.
func (u *FollowUpsertOne) SetStatus(v follow.Status) *FollowUpsertOne {
	return u.Update(func(s *FollowUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FollowUpsertOne) UpdateStatus() *FollowUpsertOne {
	return u.Update(func(s *FollowUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *FollowUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FollowCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FollowUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FollowUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: FollowUpsertOne.ID is not supported by MySQL driver. Use FollowUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FollowUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FollowCreateBulk is the builder for creating many Follow entities in bulk.
type FollowCreateBulk struct {
	config
	err      error
	builders []*FollowCreate
	conflict []sql.ConflictOption
}

// Save creates the Follow entities in the database.
func (fcb *FollowCreateBulk) Save(ctx context.Context) ([]*Follow, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Follow, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FollowMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FollowCreateBulk) SaveX(ctx context.Context) []*Follow {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FollowCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FollowCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Follow.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FollowUpsert) {
//			SetIsRemote(v+v).
//		}).
//		Exec(ctx)
func (fcb *FollowCreateBulk) OnConflict(opts ...sql.ConflictOption) *FollowUpsertBulk {
	fcb.conflict = opts
	return &FollowUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Follow.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FollowCreateBulk) OnConflictColumns(columns ...string) *FollowUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FollowUpsertBulk{
		create: fcb,
	}
}

// FollowUpsertBulk is the builder for "upsert"-ing
// a bulk of Follow nodes.
type FollowUpsertBulk struct {
	create *FollowCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Follow.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(follow.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FollowUpsertBulk) UpdateNewValues() *FollowUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(follow.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(follow.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Follow.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FollowUpsertBulk) Ignore() *FollowUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FollowUpsertBulk) DoNothing() *FollowUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FollowCreateBulk.OnConflict
// documentation for more info.
func (u *FollowUpsertBulk) Update(set func(*FollowUpsert)) *FollowUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FollowUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsRemote sets the "isRemote" field.
func (u *FollowUpsertBulk) SetIsRemote(v bool) *FollowUpsertBulk {
	return u.Update(func(s *FollowUpsert) {
		s.SetIsRemote(v)
	})
}

// UpdateIsRemote sets the "isRemote" field to the value that was provided on create.
func (u *FollowUpsertBulk) UpdateIsRemote() *FollowUpsertBulk {
	return u.Update(func(s *FollowUpsert) {
		s.UpdateIsRemote()
	})
}

// SetURI sets the "uri" field.
func (u *FollowUpsertBulk) SetURI(v string) *FollowUpsertBulk {
	return u.Update(func(s *FollowUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *FollowUpsertBulk) UpdateURI() *FollowUpsertBulk {
	return u.Update(func(s *FollowUpsert) {
		s.UpdateURI()
	})
}

// SetExtensions sets the "extensions" field.
func (u *FollowUpsertBulk) SetExtensions(v versia.Extensions) *FollowUpsertBulk {
	return u.Update(func(s *FollowUpsert) {
		s.SetExtensions(v)
	})
}

// UpdateExtensions sets the "extensions" field to the value that was provided on create.
func (u *FollowUpsertBulk) UpdateExtensions() *FollowUpsertBulk {
	return u.Update(func(s *FollowUpsert) {
		s.UpdateExtensions()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FollowUpsertBulk) SetUpdatedAt(v time.Time) *FollowUpsertBulk {
	return u.Update(func(s *FollowUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FollowUpsertBulk) UpdateUpdatedAt() *FollowUpsertBulk {
	return u.Update(func(s *FollowUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetStatus sets the "status" field.
func (u *FollowUpsertBulk) SetStatus(v follow.Status) *FollowUpsertBulk {
	return u.Update(func(s *FollowUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FollowUpsertBulk) UpdateStatus() *FollowUpsertBulk {
	return u.Update(func(s *FollowUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *FollowUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FollowCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FollowCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FollowUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
