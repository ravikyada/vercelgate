// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khanakia/vercelgate/gen/ent/predicate"
	"github.com/khanakia/vercelgate/gen/ent/team"
	"github.com/khanakia/vercelgate/gen/ent/user"
)

// TeamQuery is the builder for querying Team entities.
type TeamQuery struct {
	config
	ctx        *QueryContext
	order      []team.OrderOption
	inters     []Interceptor
	predicates []predicate.Team
	withUser   *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TeamQuery builder.
func (tq *TeamQuery) Where(ps ...predicate.Team) *TeamQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TeamQuery) Limit(limit int) *TeamQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TeamQuery) Offset(offset int) *TeamQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TeamQuery) Unique(unique bool) *TeamQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TeamQuery) Order(o ...team.OrderOption) *TeamQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryUser chains the current query on the "user" edge.
func (tq *TeamQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(team.Table, team.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, team.UserTable, team.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Team entity from the query.
// Returns a *NotFoundError when no Team was found.
func (tq *TeamQuery) First(ctx context.Context) (*Team, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{team.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TeamQuery) FirstX(ctx context.Context) *Team {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Team ID from the query.
// Returns a *NotFoundError when no Team ID was found.
func (tq *TeamQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{team.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TeamQuery) FirstIDX(ctx context.Context) string {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Team entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Team entity is found.
// Returns a *NotFoundError when no Team entities are found.
func (tq *TeamQuery) Only(ctx context.Context) (*Team, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{team.Label}
	default:
		return nil, &NotSingularError{team.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TeamQuery) OnlyX(ctx context.Context) *Team {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Team ID in the query.
// Returns a *NotSingularError when more than one Team ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TeamQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{team.Label}
	default:
		err = &NotSingularError{team.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TeamQuery) OnlyIDX(ctx context.Context) string {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Teams.
func (tq *TeamQuery) All(ctx context.Context) ([]*Team, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryAll)
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Team, *TeamQuery]()
	return withInterceptors[[]*Team](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TeamQuery) AllX(ctx context.Context) []*Team {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Team IDs.
func (tq *TeamQuery) IDs(ctx context.Context) (ids []string, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryIDs)
	if err = tq.Select(team.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TeamQuery) IDsX(ctx context.Context) []string {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TeamQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryCount)
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TeamQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TeamQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TeamQuery) Exist(ctx context.Context) (bool, error) {
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
func (tq *TeamQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TeamQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TeamQuery) Clone() *TeamQuery {
	if tq == nil {
		return nil
	}
	return &TeamQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]team.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Team{}, tq.predicates...),
		withUser:   tq.withUser.Clone(),
		// clone intermediate query.
		sql:       tq.sql.Clone(),
		path:      tq.path,
		modifiers: append([]func(*sql.Selector){}, tq.modifiers...),
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TeamQuery) WithUser(opts ...func(*UserQuery)) *TeamQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUser = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Team.Query().
//		GroupBy(team.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TeamQuery) GroupBy(field string, fields ...string) *TeamGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TeamGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = team.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Team.Query().
//		Select(team.FieldName).
//		Scan(ctx, &v)
func (tq *TeamQuery) Select(fields ...string) *TeamSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TeamSelect{TeamQuery: tq}
	sbuild.label = team.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TeamSelect configured with the given aggregations.
func (tq *TeamQuery) Aggregate(fns ...AggregateFunc) *TeamSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TeamQuery) prepareQuery(ctx context.Context) error {
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
		if !team.ValidColumn(f) {
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

func (tq *TeamQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Team, error) {
	var (
		nodes       = []*Team{}
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Team).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Team{config: tq.config}
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
	if query := tq.withUser; query != nil {
		if err := tq.loadUser(ctx, query, nodes, nil,
			func(n *Team, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TeamQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Team, init func(*Team), assign func(*Team, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Team)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TeamQuery) sqlCount(ctx context.Context) (int, error) {
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

func (tq *TeamQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeString))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, team.FieldID)
		for i := range fields {
			if fields[i] != team.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tq.withUser != nil {
			_spec.Node.AddColumnOnce(team.FieldUserID)
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

func (tq *TeamQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(team.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = team.Columns
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

// Modify adds a query modifier for attaching custom logic to queries.
func (tq *TeamQuery) Modify(modifiers ...func(s *sql.Selector)) *TeamSelect {
	tq.modifiers = append(tq.modifiers, modifiers...)
	return tq.Select()
}

// TeamGroupBy is the group-by builder for Team entities.
type TeamGroupBy struct {
	selector
	build *TeamQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TeamGroupBy) Aggregate(fns ...AggregateFunc) *TeamGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TeamGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, ent.OpQueryGroupBy)
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeamQuery, *TeamGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TeamGroupBy) sqlScan(ctx context.Context, root *TeamQuery, v any) error {
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

// TeamSelect is the builder for selecting fields of Team entities.
type TeamSelect struct {
	*TeamQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TeamSelect) Aggregate(fns ...AggregateFunc) *TeamSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TeamSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, ent.OpQuerySelect)
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeamQuery, *TeamSelect](ctx, ts.TeamQuery, ts, ts.inters, v)
}

func (ts *TeamSelect) sqlScan(ctx context.Context, root *TeamQuery, v any) error {
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

// Modify adds a query modifier for attaching custom logic to queries.
func (ts *TeamSelect) Modify(modifiers ...func(s *sql.Selector)) *TeamSelect {
	ts.modifiers = append(ts.modifiers, modifiers...)
	return ts
}
