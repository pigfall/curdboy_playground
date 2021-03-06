// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pigfall/curdboy_playground/ent/dept"
	"github.com/pigfall/curdboy_playground/ent/predicate"
	"github.com/pigfall/curdboy_playground/ent/user"
)

// DeptQuery is the builder for querying Dept entities.
type DeptQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Dept
	// eager-loading edges.
	withUser *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeptQuery builder.
func (dq *DeptQuery) Where(ps ...predicate.Dept) *DeptQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DeptQuery) Limit(limit int) *DeptQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DeptQuery) Offset(offset int) *DeptQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DeptQuery) Unique(unique bool) *DeptQuery {
	dq.unique = &unique
	return dq
}

// Order adds an order step to the query.
func (dq *DeptQuery) Order(o ...OrderFunc) *DeptQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryUser chains the current query on the "user" edge.
func (dq *DeptQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dept.Table, dept.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, dept.UserTable, dept.UserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Dept entity from the query.
// Returns a *NotFoundError when no Dept was found.
func (dq *DeptQuery) First(ctx context.Context) (*Dept, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dept.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DeptQuery) FirstX(ctx context.Context) *Dept {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Dept ID from the query.
// Returns a *NotFoundError when no Dept ID was found.
func (dq *DeptQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dept.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DeptQuery) FirstIDX(ctx context.Context) string {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Dept entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Dept entity is not found.
// Returns a *NotFoundError when no Dept entities are found.
func (dq *DeptQuery) Only(ctx context.Context) (*Dept, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dept.Label}
	default:
		return nil, &NotSingularError{dept.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DeptQuery) OnlyX(ctx context.Context) *Dept {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Dept ID in the query.
// Returns a *NotSingularError when exactly one Dept ID is not found.
// Returns a *NotFoundError when no entities are found.
func (dq *DeptQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dept.Label}
	default:
		err = &NotSingularError{dept.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DeptQuery) OnlyIDX(ctx context.Context) string {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Depts.
func (dq *DeptQuery) All(ctx context.Context) ([]*Dept, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DeptQuery) AllX(ctx context.Context) []*Dept {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Dept IDs.
func (dq *DeptQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := dq.Select(dept.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DeptQuery) IDsX(ctx context.Context) []string {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DeptQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DeptQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DeptQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DeptQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeptQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DeptQuery) Clone() *DeptQuery {
	if dq == nil {
		return nil
	}
	return &DeptQuery{
		config:     dq.config,
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]OrderFunc{}, dq.order...),
		predicates: append([]predicate.Dept{}, dq.predicates...),
		withUser:   dq.withUser.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeptQuery) WithUser(opts ...func(*UserQuery)) *DeptQuery {
	query := &UserQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withUser = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (dq *DeptQuery) GroupBy(field string, fields ...string) *DeptGroupBy {
	group := &DeptGroupBy{config: dq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (dq *DeptQuery) Select(fields ...string) *DeptSelect {
	dq.fields = append(dq.fields, fields...)
	return &DeptSelect{DeptQuery: dq}
}

func (dq *DeptQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !dept.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DeptQuery) sqlAll(ctx context.Context) ([]*Dept, error) {
	var (
		nodes       = []*Dept{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Dept{config: dq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dq.withUser; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[string]*Dept, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.User = []*User{}
		}
		var (
			edgeids []string
			edges   = make(map[string][]*Dept)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   dept.UserTable,
				Columns: dept.UserPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(dept.UserPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullString), new(sql.NullString)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullString)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullString)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := eout.String
				inValue := ein.String
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, dq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "user": %w`, err)
		}
		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "user" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = append(nodes[i].Edges.User, n)
			}
		}
	}

	return nodes, nil
}

func (dq *DeptQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DeptQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dq *DeptQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dept.Table,
			Columns: dept.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: dept.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dept.FieldID)
		for i := range fields {
			if fields[i] != dept.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DeptQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(dept.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = dept.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeptGroupBy is the group-by builder for Dept entities.
type DeptGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DeptGroupBy) Aggregate(fns ...AggregateFunc) *DeptGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DeptGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgb *DeptGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DeptGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DeptGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgb *DeptGroupBy) StringsX(ctx context.Context) []string {
	v, err := dgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DeptGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dept.Label}
	default:
		err = fmt.Errorf("ent: DeptGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dgb *DeptGroupBy) StringX(ctx context.Context) string {
	v, err := dgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DeptGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DeptGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgb *DeptGroupBy) IntsX(ctx context.Context) []int {
	v, err := dgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DeptGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dept.Label}
	default:
		err = fmt.Errorf("ent: DeptGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dgb *DeptGroupBy) IntX(ctx context.Context) int {
	v, err := dgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DeptGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DeptGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgb *DeptGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DeptGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dept.Label}
	default:
		err = fmt.Errorf("ent: DeptGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dgb *DeptGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DeptGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DeptGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgb *DeptGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DeptGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dept.Label}
	default:
		err = fmt.Errorf("ent: DeptGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dgb *DeptGroupBy) BoolX(ctx context.Context) bool {
	v, err := dgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgb *DeptGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dgb.fields {
		if !dept.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DeptGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql.Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
		for _, f := range dgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dgb.fields...)...)
}

// DeptSelect is the builder for selecting fields of Dept entities.
type DeptSelect struct {
	*DeptQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DeptSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DeptQuery.sqlQuery(ctx)
	return ds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ds *DeptSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ds *DeptSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DeptSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ds *DeptSelect) StringsX(ctx context.Context) []string {
	v, err := ds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ds *DeptSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dept.Label}
	default:
		err = fmt.Errorf("ent: DeptSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ds *DeptSelect) StringX(ctx context.Context) string {
	v, err := ds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ds *DeptSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DeptSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ds *DeptSelect) IntsX(ctx context.Context) []int {
	v, err := ds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ds *DeptSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dept.Label}
	default:
		err = fmt.Errorf("ent: DeptSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ds *DeptSelect) IntX(ctx context.Context) int {
	v, err := ds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ds *DeptSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DeptSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ds *DeptSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ds *DeptSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dept.Label}
	default:
		err = fmt.Errorf("ent: DeptSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ds *DeptSelect) Float64X(ctx context.Context) float64 {
	v, err := ds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ds *DeptSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DeptSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ds *DeptSelect) BoolsX(ctx context.Context) []bool {
	v, err := ds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ds *DeptSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dept.Label}
	default:
		err = fmt.Errorf("ent: DeptSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ds *DeptSelect) BoolX(ctx context.Context) bool {
	v, err := ds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ds *DeptSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sql.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
