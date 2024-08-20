// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lysand-org/versia-go/ent/instancemetadata"
	"github.com/lysand-org/versia-go/ent/predicate"
)

// InstanceMetadataDelete is the builder for deleting a InstanceMetadata entity.
type InstanceMetadataDelete struct {
	config
	hooks    []Hook
	mutation *InstanceMetadataMutation
}

// Where appends a list predicates to the InstanceMetadataDelete builder.
func (imd *InstanceMetadataDelete) Where(ps ...predicate.InstanceMetadata) *InstanceMetadataDelete {
	imd.mutation.Where(ps...)
	return imd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (imd *InstanceMetadataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, imd.sqlExec, imd.mutation, imd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (imd *InstanceMetadataDelete) ExecX(ctx context.Context) int {
	n, err := imd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (imd *InstanceMetadataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(instancemetadata.Table, sqlgraph.NewFieldSpec(instancemetadata.FieldID, field.TypeUUID))
	if ps := imd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, imd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	imd.mutation.done = true
	return affected, err
}

// InstanceMetadataDeleteOne is the builder for deleting a single InstanceMetadata entity.
type InstanceMetadataDeleteOne struct {
	imd *InstanceMetadataDelete
}

// Where appends a list predicates to the InstanceMetadataDelete builder.
func (imdo *InstanceMetadataDeleteOne) Where(ps ...predicate.InstanceMetadata) *InstanceMetadataDeleteOne {
	imdo.imd.mutation.Where(ps...)
	return imdo
}

// Exec executes the deletion query.
func (imdo *InstanceMetadataDeleteOne) Exec(ctx context.Context) error {
	n, err := imdo.imd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{instancemetadata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (imdo *InstanceMetadataDeleteOne) ExecX(ctx context.Context) {
	if err := imdo.Exec(ctx); err != nil {
		panic(err)
	}
}
