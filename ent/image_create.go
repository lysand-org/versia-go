// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/versia-pub/versia-go/ent/image"
)

// ImageCreate is the builder for creating a Image entity.
type ImageCreate struct {
	config
	mutation *ImageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetURL sets the "url" field.
func (ic *ImageCreate) SetURL(s string) *ImageCreate {
	ic.mutation.SetURL(s)
	return ic
}

// SetMimeType sets the "mimeType" field.
func (ic *ImageCreate) SetMimeType(s string) *ImageCreate {
	ic.mutation.SetMimeType(s)
	return ic
}

// Mutation returns the ImageMutation object of the builder.
func (ic *ImageCreate) Mutation() *ImageMutation {
	return ic.mutation
}

// Save creates the Image in the database.
func (ic *ImageCreate) Save(ctx context.Context) (*Image, error) {
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ImageCreate) SaveX(ctx context.Context) *Image {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ImageCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ImageCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ImageCreate) check() error {
	if _, ok := ic.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Image.url"`)}
	}
	if v, ok := ic.mutation.URL(); ok {
		if err := image.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Image.url": %w`, err)}
		}
	}
	if _, ok := ic.mutation.MimeType(); !ok {
		return &ValidationError{Name: "mimeType", err: errors.New(`ent: missing required field "Image.mimeType"`)}
	}
	return nil
}

func (ic *ImageCreate) sqlSave(ctx context.Context) (*Image, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ImageCreate) createSpec() (*Image, *sqlgraph.CreateSpec) {
	var (
		_node = &Image{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(image.Table, sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ic.conflict
	if value, ok := ic.mutation.URL(); ok {
		_spec.SetField(image.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ic.mutation.MimeType(); ok {
		_spec.SetField(image.FieldMimeType, field.TypeString, value)
		_node.MimeType = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Image.Create().
//		SetURL(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ImageUpsert) {
//			SetURL(v+v).
//		}).
//		Exec(ctx)
func (ic *ImageCreate) OnConflict(opts ...sql.ConflictOption) *ImageUpsertOne {
	ic.conflict = opts
	return &ImageUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Image.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ic *ImageCreate) OnConflictColumns(columns ...string) *ImageUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &ImageUpsertOne{
		create: ic,
	}
}

type (
	// ImageUpsertOne is the builder for "upsert"-ing
	//  one Image node.
	ImageUpsertOne struct {
		create *ImageCreate
	}

	// ImageUpsert is the "OnConflict" setter.
	ImageUpsert struct {
		*sql.UpdateSet
	}
)

// SetURL sets the "url" field.
func (u *ImageUpsert) SetURL(v string) *ImageUpsert {
	u.Set(image.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ImageUpsert) UpdateURL() *ImageUpsert {
	u.SetExcluded(image.FieldURL)
	return u
}

// SetMimeType sets the "mimeType" field.
func (u *ImageUpsert) SetMimeType(v string) *ImageUpsert {
	u.Set(image.FieldMimeType, v)
	return u
}

// UpdateMimeType sets the "mimeType" field to the value that was provided on create.
func (u *ImageUpsert) UpdateMimeType() *ImageUpsert {
	u.SetExcluded(image.FieldMimeType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Image.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ImageUpsertOne) UpdateNewValues() *ImageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Image.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ImageUpsertOne) Ignore() *ImageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ImageUpsertOne) DoNothing() *ImageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ImageCreate.OnConflict
// documentation for more info.
func (u *ImageUpsertOne) Update(set func(*ImageUpsert)) *ImageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ImageUpsert{UpdateSet: update})
	}))
	return u
}

// SetURL sets the "url" field.
func (u *ImageUpsertOne) SetURL(v string) *ImageUpsertOne {
	return u.Update(func(s *ImageUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ImageUpsertOne) UpdateURL() *ImageUpsertOne {
	return u.Update(func(s *ImageUpsert) {
		s.UpdateURL()
	})
}

// SetMimeType sets the "mimeType" field.
func (u *ImageUpsertOne) SetMimeType(v string) *ImageUpsertOne {
	return u.Update(func(s *ImageUpsert) {
		s.SetMimeType(v)
	})
}

// UpdateMimeType sets the "mimeType" field to the value that was provided on create.
func (u *ImageUpsertOne) UpdateMimeType() *ImageUpsertOne {
	return u.Update(func(s *ImageUpsert) {
		s.UpdateMimeType()
	})
}

// Exec executes the query.
func (u *ImageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ImageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ImageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ImageUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ImageUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ImageCreateBulk is the builder for creating many Image entities in bulk.
type ImageCreateBulk struct {
	config
	err      error
	builders []*ImageCreate
	conflict []sql.ConflictOption
}

// Save creates the Image entities in the database.
func (icb *ImageCreateBulk) Save(ctx context.Context) ([]*Image, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Image, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImageMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = icb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ImageCreateBulk) SaveX(ctx context.Context) []*Image {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ImageCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ImageCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Image.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ImageUpsert) {
//			SetURL(v+v).
//		}).
//		Exec(ctx)
func (icb *ImageCreateBulk) OnConflict(opts ...sql.ConflictOption) *ImageUpsertBulk {
	icb.conflict = opts
	return &ImageUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Image.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (icb *ImageCreateBulk) OnConflictColumns(columns ...string) *ImageUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &ImageUpsertBulk{
		create: icb,
	}
}

// ImageUpsertBulk is the builder for "upsert"-ing
// a bulk of Image nodes.
type ImageUpsertBulk struct {
	create *ImageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Image.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ImageUpsertBulk) UpdateNewValues() *ImageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Image.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ImageUpsertBulk) Ignore() *ImageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ImageUpsertBulk) DoNothing() *ImageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ImageCreateBulk.OnConflict
// documentation for more info.
func (u *ImageUpsertBulk) Update(set func(*ImageUpsert)) *ImageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ImageUpsert{UpdateSet: update})
	}))
	return u
}

// SetURL sets the "url" field.
func (u *ImageUpsertBulk) SetURL(v string) *ImageUpsertBulk {
	return u.Update(func(s *ImageUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ImageUpsertBulk) UpdateURL() *ImageUpsertBulk {
	return u.Update(func(s *ImageUpsert) {
		s.UpdateURL()
	})
}

// SetMimeType sets the "mimeType" field.
func (u *ImageUpsertBulk) SetMimeType(v string) *ImageUpsertBulk {
	return u.Update(func(s *ImageUpsert) {
		s.SetMimeType(v)
	})
}

// UpdateMimeType sets the "mimeType" field to the value that was provided on create.
func (u *ImageUpsertBulk) UpdateMimeType() *ImageUpsertBulk {
	return u.Update(func(s *ImageUpsert) {
		s.UpdateMimeType()
	})
}

// Exec executes the query.
func (u *ImageUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ImageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ImageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ImageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
