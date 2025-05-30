// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lightsparkdev/spark/so/ent/predicate"
	"github.com/lightsparkdev/spark/so/ent/tree"
	"github.com/lightsparkdev/spark/so/ent/treenode"
)

// TreeQuery is the builder for querying Tree entities.
type TreeQuery struct {
	config
	ctx        *QueryContext
	order      []tree.OrderOption
	inters     []Interceptor
	predicates []predicate.Tree
	withRoot   *TreeNodeQuery
	withNodes  *TreeNodeQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TreeQuery builder.
func (tq *TreeQuery) Where(ps ...predicate.Tree) *TreeQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TreeQuery) Limit(limit int) *TreeQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TreeQuery) Offset(offset int) *TreeQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TreeQuery) Unique(unique bool) *TreeQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TreeQuery) Order(o ...tree.OrderOption) *TreeQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryRoot chains the current query on the "root" edge.
func (tq *TreeQuery) QueryRoot() *TreeNodeQuery {
	query := (&TreeNodeClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tree.Table, tree.FieldID, selector),
			sqlgraph.To(treenode.Table, treenode.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tree.RootTable, tree.RootColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNodes chains the current query on the "nodes" edge.
func (tq *TreeQuery) QueryNodes() *TreeNodeQuery {
	query := (&TreeNodeClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tree.Table, tree.FieldID, selector),
			sqlgraph.To(treenode.Table, treenode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, tree.NodesTable, tree.NodesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Tree entity from the query.
// Returns a *NotFoundError when no Tree was found.
func (tq *TreeQuery) First(ctx context.Context) (*Tree, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tree.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TreeQuery) FirstX(ctx context.Context) *Tree {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Tree ID from the query.
// Returns a *NotFoundError when no Tree ID was found.
func (tq *TreeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tree.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TreeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Tree entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Tree entity is found.
// Returns a *NotFoundError when no Tree entities are found.
func (tq *TreeQuery) Only(ctx context.Context) (*Tree, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tree.Label}
	default:
		return nil, &NotSingularError{tree.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TreeQuery) OnlyX(ctx context.Context) *Tree {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Tree ID in the query.
// Returns a *NotSingularError when more than one Tree ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TreeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tree.Label}
	default:
		err = &NotSingularError{tree.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TreeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Trees.
func (tq *TreeQuery) All(ctx context.Context) ([]*Tree, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryAll)
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Tree, *TreeQuery]()
	return withInterceptors[[]*Tree](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TreeQuery) AllX(ctx context.Context) []*Tree {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Tree IDs.
func (tq *TreeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryIDs)
	if err = tq.Select(tree.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TreeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TreeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryCount)
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TreeQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TreeQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TreeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryExist)
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TreeQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TreeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TreeQuery) Clone() *TreeQuery {
	if tq == nil {
		return nil
	}
	return &TreeQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]tree.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Tree{}, tq.predicates...),
		withRoot:   tq.withRoot.Clone(),
		withNodes:  tq.withNodes.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithRoot tells the query-builder to eager-load the nodes that are connected to
// the "root" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TreeQuery) WithRoot(opts ...func(*TreeNodeQuery)) *TreeQuery {
	query := (&TreeNodeClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withRoot = query
	return tq
}

// WithNodes tells the query-builder to eager-load the nodes that are connected to
// the "nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TreeQuery) WithNodes(opts ...func(*TreeNodeQuery)) *TreeQuery {
	query := (&TreeNodeClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withNodes = query
	return tq
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
//	client.Tree.Query().
//		GroupBy(tree.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TreeQuery) GroupBy(field string, fields ...string) *TreeGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TreeGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = tree.Label
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
//	client.Tree.Query().
//		Select(tree.FieldCreateTime).
//		Scan(ctx, &v)
func (tq *TreeQuery) Select(fields ...string) *TreeSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TreeSelect{TreeQuery: tq}
	sbuild.label = tree.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TreeSelect configured with the given aggregations.
func (tq *TreeQuery) Aggregate(fns ...AggregateFunc) *TreeSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TreeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !tree.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TreeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Tree, error) {
	var (
		nodes       = []*Tree{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withRoot != nil,
			tq.withNodes != nil,
		}
	)
	if tq.withRoot != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, tree.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Tree).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Tree{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withRoot; query != nil {
		if err := tq.loadRoot(ctx, query, nodes, nil,
			func(n *Tree, e *TreeNode) { n.Edges.Root = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withNodes; query != nil {
		if err := tq.loadNodes(ctx, query, nodes,
			func(n *Tree) { n.Edges.Nodes = []*TreeNode{} },
			func(n *Tree, e *TreeNode) { n.Edges.Nodes = append(n.Edges.Nodes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TreeQuery) loadRoot(ctx context.Context, query *TreeNodeQuery, nodes []*Tree, init func(*Tree), assign func(*Tree, *TreeNode)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Tree)
	for i := range nodes {
		if nodes[i].tree_root == nil {
			continue
		}
		fk := *nodes[i].tree_root
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(treenode.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tree_root" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TreeQuery) loadNodes(ctx context.Context, query *TreeNodeQuery, nodes []*Tree, init func(*Tree), assign func(*Tree, *TreeNode)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Tree)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TreeNode(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(tree.NodesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.tree_node_tree
		if fk == nil {
			return fmt.Errorf(`foreign-key "tree_node_tree" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "tree_node_tree" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TreeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TreeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tree.Table, tree.Columns, sqlgraph.NewFieldSpec(tree.FieldID, field.TypeUUID))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tree.FieldID)
		for i := range fields {
			if fields[i] != tree.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TreeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(tree.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = tree.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tq.modifiers {
		m(selector)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tq *TreeQuery) ForUpdate(opts ...sql.LockOption) *TreeQuery {
	if tq.driver.Dialect() == dialect.Postgres {
		tq.Unique(false)
	}
	tq.modifiers = append(tq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tq *TreeQuery) ForShare(opts ...sql.LockOption) *TreeQuery {
	if tq.driver.Dialect() == dialect.Postgres {
		tq.Unique(false)
	}
	tq.modifiers = append(tq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tq
}

// TreeGroupBy is the group-by builder for Tree entities.
type TreeGroupBy struct {
	selector
	build *TreeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TreeGroupBy) Aggregate(fns ...AggregateFunc) *TreeGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TreeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, ent.OpQueryGroupBy)
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TreeQuery, *TreeGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TreeGroupBy) sqlScan(ctx context.Context, root *TreeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TreeSelect is the builder for selecting fields of Tree entities.
type TreeSelect struct {
	*TreeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TreeSelect) Aggregate(fns ...AggregateFunc) *TreeSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TreeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, ent.OpQuerySelect)
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TreeQuery, *TreeSelect](ctx, ts.TreeQuery, ts, ts.inters, v)
}

func (ts *TreeSelect) sqlScan(ctx context.Context, root *TreeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
