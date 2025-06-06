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
	"github.com/lightsparkdev/spark/so/ent/depositaddress"
	"github.com/lightsparkdev/spark/so/ent/predicate"
	"github.com/lightsparkdev/spark/so/ent/utxo"
)

// UtxoQuery is the builder for querying Utxo entities.
type UtxoQuery struct {
	config
	ctx                *QueryContext
	order              []utxo.OrderOption
	inters             []Interceptor
	predicates         []predicate.Utxo
	withDepositAddress *DepositAddressQuery
	withFKs            bool
	modifiers          []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UtxoQuery builder.
func (uq *UtxoQuery) Where(ps ...predicate.Utxo) *UtxoQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UtxoQuery) Limit(limit int) *UtxoQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UtxoQuery) Offset(offset int) *UtxoQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UtxoQuery) Unique(unique bool) *UtxoQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UtxoQuery) Order(o ...utxo.OrderOption) *UtxoQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryDepositAddress chains the current query on the "deposit_address" edge.
func (uq *UtxoQuery) QueryDepositAddress() *DepositAddressQuery {
	query := (&DepositAddressClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(utxo.Table, utxo.FieldID, selector),
			sqlgraph.To(depositaddress.Table, depositaddress.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, utxo.DepositAddressTable, utxo.DepositAddressColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Utxo entity from the query.
// Returns a *NotFoundError when no Utxo was found.
func (uq *UtxoQuery) First(ctx context.Context) (*Utxo, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{utxo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UtxoQuery) FirstX(ctx context.Context) *Utxo {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Utxo ID from the query.
// Returns a *NotFoundError when no Utxo ID was found.
func (uq *UtxoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{utxo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UtxoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Utxo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Utxo entity is found.
// Returns a *NotFoundError when no Utxo entities are found.
func (uq *UtxoQuery) Only(ctx context.Context) (*Utxo, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{utxo.Label}
	default:
		return nil, &NotSingularError{utxo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UtxoQuery) OnlyX(ctx context.Context) *Utxo {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Utxo ID in the query.
// Returns a *NotSingularError when more than one Utxo ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UtxoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{utxo.Label}
	default:
		err = &NotSingularError{utxo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UtxoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Utxos.
func (uq *UtxoQuery) All(ctx context.Context) ([]*Utxo, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryAll)
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Utxo, *UtxoQuery]()
	return withInterceptors[[]*Utxo](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UtxoQuery) AllX(ctx context.Context) []*Utxo {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Utxo IDs.
func (uq *UtxoQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryIDs)
	if err = uq.Select(utxo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UtxoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UtxoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryCount)
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UtxoQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UtxoQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UtxoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryExist)
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UtxoQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UtxoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UtxoQuery) Clone() *UtxoQuery {
	if uq == nil {
		return nil
	}
	return &UtxoQuery{
		config:             uq.config,
		ctx:                uq.ctx.Clone(),
		order:              append([]utxo.OrderOption{}, uq.order...),
		inters:             append([]Interceptor{}, uq.inters...),
		predicates:         append([]predicate.Utxo{}, uq.predicates...),
		withDepositAddress: uq.withDepositAddress.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithDepositAddress tells the query-builder to eager-load the nodes that are connected to
// the "deposit_address" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UtxoQuery) WithDepositAddress(opts ...func(*DepositAddressQuery)) *UtxoQuery {
	query := (&DepositAddressClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withDepositAddress = query
	return uq
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
//	client.Utxo.Query().
//		GroupBy(utxo.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UtxoQuery) GroupBy(field string, fields ...string) *UtxoGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UtxoGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = utxo.Label
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
//	client.Utxo.Query().
//		Select(utxo.FieldCreateTime).
//		Scan(ctx, &v)
func (uq *UtxoQuery) Select(fields ...string) *UtxoSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UtxoSelect{UtxoQuery: uq}
	sbuild.label = utxo.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UtxoSelect configured with the given aggregations.
func (uq *UtxoQuery) Aggregate(fns ...AggregateFunc) *UtxoSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UtxoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !utxo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UtxoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Utxo, error) {
	var (
		nodes       = []*Utxo{}
		withFKs     = uq.withFKs
		_spec       = uq.querySpec()
		loadedTypes = [1]bool{
			uq.withDepositAddress != nil,
		}
	)
	if uq.withDepositAddress != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, utxo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Utxo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Utxo{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withDepositAddress; query != nil {
		if err := uq.loadDepositAddress(ctx, query, nodes, nil,
			func(n *Utxo, e *DepositAddress) { n.Edges.DepositAddress = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UtxoQuery) loadDepositAddress(ctx context.Context, query *DepositAddressQuery, nodes []*Utxo, init func(*Utxo), assign func(*Utxo, *DepositAddress)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Utxo)
	for i := range nodes {
		if nodes[i].deposit_address_utxo == nil {
			continue
		}
		fk := *nodes[i].deposit_address_utxo
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(depositaddress.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "deposit_address_utxo" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uq *UtxoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UtxoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(utxo.Table, utxo.Columns, sqlgraph.NewFieldSpec(utxo.FieldID, field.TypeUUID))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, utxo.FieldID)
		for i := range fields {
			if fields[i] != utxo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UtxoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(utxo.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = utxo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range uq.modifiers {
		m(selector)
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (uq *UtxoQuery) ForUpdate(opts ...sql.LockOption) *UtxoQuery {
	if uq.driver.Dialect() == dialect.Postgres {
		uq.Unique(false)
	}
	uq.modifiers = append(uq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return uq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (uq *UtxoQuery) ForShare(opts ...sql.LockOption) *UtxoQuery {
	if uq.driver.Dialect() == dialect.Postgres {
		uq.Unique(false)
	}
	uq.modifiers = append(uq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return uq
}

// UtxoGroupBy is the group-by builder for Utxo entities.
type UtxoGroupBy struct {
	selector
	build *UtxoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UtxoGroupBy) Aggregate(fns ...AggregateFunc) *UtxoGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UtxoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, ent.OpQueryGroupBy)
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UtxoQuery, *UtxoGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UtxoGroupBy) sqlScan(ctx context.Context, root *UtxoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UtxoSelect is the builder for selecting fields of Utxo entities.
type UtxoSelect struct {
	*UtxoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UtxoSelect) Aggregate(fns ...AggregateFunc) *UtxoSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UtxoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, ent.OpQuerySelect)
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UtxoQuery, *UtxoSelect](ctx, us.UtxoQuery, us, us.inters, v)
}

func (us *UtxoSelect) sqlScan(ctx context.Context, root *UtxoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
