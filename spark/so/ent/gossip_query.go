// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lightsparkdev/spark/so/ent/gossip"
	"github.com/lightsparkdev/spark/so/ent/predicate"
)

// GossipQuery is the builder for querying Gossip entities.
type GossipQuery struct {
	config
	ctx        *QueryContext
	order      []gossip.OrderOption
	inters     []Interceptor
	predicates []predicate.Gossip
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GossipQuery builder.
func (gq *GossipQuery) Where(ps ...predicate.Gossip) *GossipQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit the number of records to be returned by this query.
func (gq *GossipQuery) Limit(limit int) *GossipQuery {
	gq.ctx.Limit = &limit
	return gq
}

// Offset to start from.
func (gq *GossipQuery) Offset(offset int) *GossipQuery {
	gq.ctx.Offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GossipQuery) Unique(unique bool) *GossipQuery {
	gq.ctx.Unique = &unique
	return gq
}

// Order specifies how the records should be ordered.
func (gq *GossipQuery) Order(o ...gossip.OrderOption) *GossipQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// First returns the first Gossip entity from the query.
// Returns a *NotFoundError when no Gossip was found.
func (gq *GossipQuery) First(ctx context.Context) (*Gossip, error) {
	nodes, err := gq.Limit(1).All(setContextOp(ctx, gq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gossip.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GossipQuery) FirstX(ctx context.Context) *Gossip {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Gossip ID from the query.
// Returns a *NotFoundError when no Gossip ID was found.
func (gq *GossipQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gq.Limit(1).IDs(setContextOp(ctx, gq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gossip.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GossipQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Gossip entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Gossip entity is found.
// Returns a *NotFoundError when no Gossip entities are found.
func (gq *GossipQuery) Only(ctx context.Context) (*Gossip, error) {
	nodes, err := gq.Limit(2).All(setContextOp(ctx, gq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gossip.Label}
	default:
		return nil, &NotSingularError{gossip.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GossipQuery) OnlyX(ctx context.Context) *Gossip {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Gossip ID in the query.
// Returns a *NotSingularError when more than one Gossip ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GossipQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gq.Limit(2).IDs(setContextOp(ctx, gq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gossip.Label}
	default:
		err = &NotSingularError{gossip.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GossipQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Gossips.
func (gq *GossipQuery) All(ctx context.Context) ([]*Gossip, error) {
	ctx = setContextOp(ctx, gq.ctx, ent.OpQueryAll)
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Gossip, *GossipQuery]()
	return withInterceptors[[]*Gossip](ctx, gq, qr, gq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gq *GossipQuery) AllX(ctx context.Context) []*Gossip {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Gossip IDs.
func (gq *GossipQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if gq.ctx.Unique == nil && gq.path != nil {
		gq.Unique(true)
	}
	ctx = setContextOp(ctx, gq.ctx, ent.OpQueryIDs)
	if err = gq.Select(gossip.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GossipQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GossipQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gq.ctx, ent.OpQueryCount)
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gq, querierCount[*GossipQuery](), gq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GossipQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GossipQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gq.ctx, ent.OpQueryExist)
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GossipQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GossipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GossipQuery) Clone() *GossipQuery {
	if gq == nil {
		return nil
	}
	return &GossipQuery{
		config:     gq.config,
		ctx:        gq.ctx.Clone(),
		order:      append([]gossip.OrderOption{}, gq.order...),
		inters:     append([]Interceptor{}, gq.inters...),
		predicates: append([]predicate.Gossip{}, gq.predicates...),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Gossip.Query().
//		GroupBy(gossip.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gq *GossipQuery) GroupBy(field string, fields ...string) *GossipGroupBy {
	gq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GossipGroupBy{build: gq}
	grbuild.flds = &gq.ctx.Fields
	grbuild.label = gossip.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Gossip.Query().
//		Select(gossip.FieldCreateTime).
//		Scan(ctx, &v)
func (gq *GossipQuery) Select(fields ...string) *GossipSelect {
	gq.ctx.Fields = append(gq.ctx.Fields, fields...)
	sbuild := &GossipSelect{GossipQuery: gq}
	sbuild.label = gossip.Label
	sbuild.flds, sbuild.scan = &gq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GossipSelect configured with the given aggregations.
func (gq *GossipQuery) Aggregate(fns ...AggregateFunc) *GossipSelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GossipQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gq); err != nil {
				return err
			}
		}
	}
	for _, f := range gq.ctx.Fields {
		if !gossip.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GossipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Gossip, error) {
	var (
		nodes = []*Gossip{}
		_spec = gq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Gossip).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Gossip{config: gq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gq *GossipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	_spec.Node.Columns = gq.ctx.Fields
	if len(gq.ctx.Fields) > 0 {
		_spec.Unique = gq.ctx.Unique != nil && *gq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GossipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(gossip.Table, gossip.Columns, sqlgraph.NewFieldSpec(gossip.FieldID, field.TypeUUID))
	_spec.From = gq.sql
	if unique := gq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gq.path != nil {
		_spec.Unique = true
	}
	if fields := gq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gossip.FieldID)
		for i := range fields {
			if fields[i] != gossip.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GossipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(gossip.Table)
	columns := gq.ctx.Fields
	if len(columns) == 0 {
		columns = gossip.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.ctx.Unique != nil && *gq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range gq.modifiers {
		m(selector)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gq *GossipQuery) ForUpdate(opts ...sql.LockOption) *GossipQuery {
	if gq.driver.Dialect() == dialect.Postgres {
		gq.Unique(false)
	}
	gq.modifiers = append(gq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gq *GossipQuery) ForShare(opts ...sql.LockOption) *GossipQuery {
	if gq.driver.Dialect() == dialect.Postgres {
		gq.Unique(false)
	}
	gq.modifiers = append(gq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gq
}

// GossipGroupBy is the group-by builder for Gossip entities.
type GossipGroupBy struct {
	selector
	build *GossipQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GossipGroupBy) Aggregate(fns ...AggregateFunc) *GossipGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GossipGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, ent.OpQueryGroupBy)
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GossipQuery, *GossipGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GossipGroupBy) sqlScan(ctx context.Context, root *GossipQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GossipSelect is the builder for selecting fields of Gossip entities.
type GossipSelect struct {
	*GossipQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GossipSelect) Aggregate(fns ...AggregateFunc) *GossipSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GossipSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, ent.OpQuerySelect)
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GossipQuery, *GossipSelect](ctx, gs.GossipQuery, gs, gs.inters, v)
}

func (gs *GossipSelect) sqlScan(ctx context.Context, root *GossipQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
