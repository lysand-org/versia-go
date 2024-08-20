// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/ent/instancemetadata"
	"github.com/lysand-org/versia-go/ent/predicate"
	"github.com/lysand-org/versia-go/ent/user"
)

// InstanceMetadataQuery is the builder for querying InstanceMetadata entities.
type InstanceMetadataQuery struct {
	config
	ctx            *QueryContext
	order          []instancemetadata.OrderOption
	inters         []Interceptor
	predicates     []predicate.InstanceMetadata
	withUsers      *UserQuery
	withModerators *UserQuery
	withAdmins     *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InstanceMetadataQuery builder.
func (imq *InstanceMetadataQuery) Where(ps ...predicate.InstanceMetadata) *InstanceMetadataQuery {
	imq.predicates = append(imq.predicates, ps...)
	return imq
}

// Limit the number of records to be returned by this query.
func (imq *InstanceMetadataQuery) Limit(limit int) *InstanceMetadataQuery {
	imq.ctx.Limit = &limit
	return imq
}

// Offset to start from.
func (imq *InstanceMetadataQuery) Offset(offset int) *InstanceMetadataQuery {
	imq.ctx.Offset = &offset
	return imq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (imq *InstanceMetadataQuery) Unique(unique bool) *InstanceMetadataQuery {
	imq.ctx.Unique = &unique
	return imq
}

// Order specifies how the records should be ordered.
func (imq *InstanceMetadataQuery) Order(o ...instancemetadata.OrderOption) *InstanceMetadataQuery {
	imq.order = append(imq.order, o...)
	return imq
}

// QueryUsers chains the current query on the "users" edge.
func (imq *InstanceMetadataQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: imq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := imq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := imq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instancemetadata.Table, instancemetadata.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, instancemetadata.UsersTable, instancemetadata.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(imq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryModerators chains the current query on the "moderators" edge.
func (imq *InstanceMetadataQuery) QueryModerators() *UserQuery {
	query := (&UserClient{config: imq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := imq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := imq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instancemetadata.Table, instancemetadata.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, instancemetadata.ModeratorsTable, instancemetadata.ModeratorsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(imq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAdmins chains the current query on the "admins" edge.
func (imq *InstanceMetadataQuery) QueryAdmins() *UserQuery {
	query := (&UserClient{config: imq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := imq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := imq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instancemetadata.Table, instancemetadata.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, instancemetadata.AdminsTable, instancemetadata.AdminsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(imq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InstanceMetadata entity from the query.
// Returns a *NotFoundError when no InstanceMetadata was found.
func (imq *InstanceMetadataQuery) First(ctx context.Context) (*InstanceMetadata, error) {
	nodes, err := imq.Limit(1).All(setContextOp(ctx, imq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{instancemetadata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (imq *InstanceMetadataQuery) FirstX(ctx context.Context) *InstanceMetadata {
	node, err := imq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InstanceMetadata ID from the query.
// Returns a *NotFoundError when no InstanceMetadata ID was found.
func (imq *InstanceMetadataQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = imq.Limit(1).IDs(setContextOp(ctx, imq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{instancemetadata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (imq *InstanceMetadataQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := imq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InstanceMetadata entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InstanceMetadata entity is found.
// Returns a *NotFoundError when no InstanceMetadata entities are found.
func (imq *InstanceMetadataQuery) Only(ctx context.Context) (*InstanceMetadata, error) {
	nodes, err := imq.Limit(2).All(setContextOp(ctx, imq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{instancemetadata.Label}
	default:
		return nil, &NotSingularError{instancemetadata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (imq *InstanceMetadataQuery) OnlyX(ctx context.Context) *InstanceMetadata {
	node, err := imq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InstanceMetadata ID in the query.
// Returns a *NotSingularError when more than one InstanceMetadata ID is found.
// Returns a *NotFoundError when no entities are found.
func (imq *InstanceMetadataQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = imq.Limit(2).IDs(setContextOp(ctx, imq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{instancemetadata.Label}
	default:
		err = &NotSingularError{instancemetadata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (imq *InstanceMetadataQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := imq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InstanceMetadataSlice.
func (imq *InstanceMetadataQuery) All(ctx context.Context) ([]*InstanceMetadata, error) {
	ctx = setContextOp(ctx, imq.ctx, "All")
	if err := imq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*InstanceMetadata, *InstanceMetadataQuery]()
	return withInterceptors[[]*InstanceMetadata](ctx, imq, qr, imq.inters)
}

// AllX is like All, but panics if an error occurs.
func (imq *InstanceMetadataQuery) AllX(ctx context.Context) []*InstanceMetadata {
	nodes, err := imq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InstanceMetadata IDs.
func (imq *InstanceMetadataQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if imq.ctx.Unique == nil && imq.path != nil {
		imq.Unique(true)
	}
	ctx = setContextOp(ctx, imq.ctx, "IDs")
	if err = imq.Select(instancemetadata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (imq *InstanceMetadataQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := imq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (imq *InstanceMetadataQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, imq.ctx, "Count")
	if err := imq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, imq, querierCount[*InstanceMetadataQuery](), imq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (imq *InstanceMetadataQuery) CountX(ctx context.Context) int {
	count, err := imq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (imq *InstanceMetadataQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, imq.ctx, "Exist")
	switch _, err := imq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (imq *InstanceMetadataQuery) ExistX(ctx context.Context) bool {
	exist, err := imq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InstanceMetadataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (imq *InstanceMetadataQuery) Clone() *InstanceMetadataQuery {
	if imq == nil {
		return nil
	}
	return &InstanceMetadataQuery{
		config:         imq.config,
		ctx:            imq.ctx.Clone(),
		order:          append([]instancemetadata.OrderOption{}, imq.order...),
		inters:         append([]Interceptor{}, imq.inters...),
		predicates:     append([]predicate.InstanceMetadata{}, imq.predicates...),
		withUsers:      imq.withUsers.Clone(),
		withModerators: imq.withModerators.Clone(),
		withAdmins:     imq.withAdmins.Clone(),
		// clone intermediate query.
		sql:  imq.sql.Clone(),
		path: imq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (imq *InstanceMetadataQuery) WithUsers(opts ...func(*UserQuery)) *InstanceMetadataQuery {
	query := (&UserClient{config: imq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	imq.withUsers = query
	return imq
}

// WithModerators tells the query-builder to eager-load the nodes that are connected to
// the "moderators" edge. The optional arguments are used to configure the query builder of the edge.
func (imq *InstanceMetadataQuery) WithModerators(opts ...func(*UserQuery)) *InstanceMetadataQuery {
	query := (&UserClient{config: imq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	imq.withModerators = query
	return imq
}

// WithAdmins tells the query-builder to eager-load the nodes that are connected to
// the "admins" edge. The optional arguments are used to configure the query builder of the edge.
func (imq *InstanceMetadataQuery) WithAdmins(opts ...func(*UserQuery)) *InstanceMetadataQuery {
	query := (&UserClient{config: imq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	imq.withAdmins = query
	return imq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IsRemote bool `json:"isRemote,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.InstanceMetadata.Query().
//		GroupBy(instancemetadata.FieldIsRemote).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (imq *InstanceMetadataQuery) GroupBy(field string, fields ...string) *InstanceMetadataGroupBy {
	imq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InstanceMetadataGroupBy{build: imq}
	grbuild.flds = &imq.ctx.Fields
	grbuild.label = instancemetadata.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IsRemote bool `json:"isRemote,omitempty"`
//	}
//
//	client.InstanceMetadata.Query().
//		Select(instancemetadata.FieldIsRemote).
//		Scan(ctx, &v)
func (imq *InstanceMetadataQuery) Select(fields ...string) *InstanceMetadataSelect {
	imq.ctx.Fields = append(imq.ctx.Fields, fields...)
	sbuild := &InstanceMetadataSelect{InstanceMetadataQuery: imq}
	sbuild.label = instancemetadata.Label
	sbuild.flds, sbuild.scan = &imq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InstanceMetadataSelect configured with the given aggregations.
func (imq *InstanceMetadataQuery) Aggregate(fns ...AggregateFunc) *InstanceMetadataSelect {
	return imq.Select().Aggregate(fns...)
}

func (imq *InstanceMetadataQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range imq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, imq); err != nil {
				return err
			}
		}
	}
	for _, f := range imq.ctx.Fields {
		if !instancemetadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if imq.path != nil {
		prev, err := imq.path(ctx)
		if err != nil {
			return err
		}
		imq.sql = prev
	}
	return nil
}

func (imq *InstanceMetadataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*InstanceMetadata, error) {
	var (
		nodes       = []*InstanceMetadata{}
		_spec       = imq.querySpec()
		loadedTypes = [3]bool{
			imq.withUsers != nil,
			imq.withModerators != nil,
			imq.withAdmins != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*InstanceMetadata).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &InstanceMetadata{config: imq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, imq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := imq.withUsers; query != nil {
		if err := imq.loadUsers(ctx, query, nodes,
			func(n *InstanceMetadata) { n.Edges.Users = []*User{} },
			func(n *InstanceMetadata, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	if query := imq.withModerators; query != nil {
		if err := imq.loadModerators(ctx, query, nodes,
			func(n *InstanceMetadata) { n.Edges.Moderators = []*User{} },
			func(n *InstanceMetadata, e *User) { n.Edges.Moderators = append(n.Edges.Moderators, e) }); err != nil {
			return nil, err
		}
	}
	if query := imq.withAdmins; query != nil {
		if err := imq.loadAdmins(ctx, query, nodes,
			func(n *InstanceMetadata) { n.Edges.Admins = []*User{} },
			func(n *InstanceMetadata, e *User) { n.Edges.Admins = append(n.Edges.Admins, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (imq *InstanceMetadataQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*InstanceMetadata, init func(*InstanceMetadata), assign func(*InstanceMetadata, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*InstanceMetadata)
	nids := make(map[uuid.UUID]map[*InstanceMetadata]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(instancemetadata.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(instancemetadata.UsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(instancemetadata.UsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(instancemetadata.UsersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*InstanceMetadata]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (imq *InstanceMetadataQuery) loadModerators(ctx context.Context, query *UserQuery, nodes []*InstanceMetadata, init func(*InstanceMetadata), assign func(*InstanceMetadata, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*InstanceMetadata)
	nids := make(map[uuid.UUID]map[*InstanceMetadata]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(instancemetadata.ModeratorsTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(instancemetadata.ModeratorsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(instancemetadata.ModeratorsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(instancemetadata.ModeratorsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*InstanceMetadata]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "moderators" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (imq *InstanceMetadataQuery) loadAdmins(ctx context.Context, query *UserQuery, nodes []*InstanceMetadata, init func(*InstanceMetadata), assign func(*InstanceMetadata, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*InstanceMetadata)
	nids := make(map[uuid.UUID]map[*InstanceMetadata]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(instancemetadata.AdminsTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(instancemetadata.AdminsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(instancemetadata.AdminsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(instancemetadata.AdminsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*InstanceMetadata]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "admins" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (imq *InstanceMetadataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := imq.querySpec()
	_spec.Node.Columns = imq.ctx.Fields
	if len(imq.ctx.Fields) > 0 {
		_spec.Unique = imq.ctx.Unique != nil && *imq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, imq.driver, _spec)
}

func (imq *InstanceMetadataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(instancemetadata.Table, instancemetadata.Columns, sqlgraph.NewFieldSpec(instancemetadata.FieldID, field.TypeUUID))
	_spec.From = imq.sql
	if unique := imq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if imq.path != nil {
		_spec.Unique = true
	}
	if fields := imq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, instancemetadata.FieldID)
		for i := range fields {
			if fields[i] != instancemetadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := imq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := imq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := imq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := imq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (imq *InstanceMetadataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(imq.driver.Dialect())
	t1 := builder.Table(instancemetadata.Table)
	columns := imq.ctx.Fields
	if len(columns) == 0 {
		columns = instancemetadata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if imq.sql != nil {
		selector = imq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if imq.ctx.Unique != nil && *imq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range imq.predicates {
		p(selector)
	}
	for _, p := range imq.order {
		p(selector)
	}
	if offset := imq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := imq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InstanceMetadataGroupBy is the group-by builder for InstanceMetadata entities.
type InstanceMetadataGroupBy struct {
	selector
	build *InstanceMetadataQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (imgb *InstanceMetadataGroupBy) Aggregate(fns ...AggregateFunc) *InstanceMetadataGroupBy {
	imgb.fns = append(imgb.fns, fns...)
	return imgb
}

// Scan applies the selector query and scans the result into the given value.
func (imgb *InstanceMetadataGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, imgb.build.ctx, "GroupBy")
	if err := imgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InstanceMetadataQuery, *InstanceMetadataGroupBy](ctx, imgb.build, imgb, imgb.build.inters, v)
}

func (imgb *InstanceMetadataGroupBy) sqlScan(ctx context.Context, root *InstanceMetadataQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(imgb.fns))
	for _, fn := range imgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*imgb.flds)+len(imgb.fns))
		for _, f := range *imgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*imgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := imgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InstanceMetadataSelect is the builder for selecting fields of InstanceMetadata entities.
type InstanceMetadataSelect struct {
	*InstanceMetadataQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ims *InstanceMetadataSelect) Aggregate(fns ...AggregateFunc) *InstanceMetadataSelect {
	ims.fns = append(ims.fns, fns...)
	return ims
}

// Scan applies the selector query and scans the result into the given value.
func (ims *InstanceMetadataSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ims.ctx, "Select")
	if err := ims.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InstanceMetadataQuery, *InstanceMetadataSelect](ctx, ims.InstanceMetadataQuery, ims, ims.inters, v)
}

func (ims *InstanceMetadataSelect) sqlScan(ctx context.Context, root *InstanceMetadataQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ims.fns))
	for _, fn := range ims.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ims.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ims.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
