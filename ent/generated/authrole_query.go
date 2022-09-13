// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/ulid-test/ent/generated/account"
	"github.com/chenningg/ulid-test/ent/generated/authrole"
	"github.com/chenningg/ulid-test/ent/generated/predicate"
	"github.com/oklog/ulid/v2"
	ulid "github.com/oklog/ulid/v2"
)

// AuthRoleQuery is the builder for querying AuthRole entities.
type AuthRoleQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.AuthRole
	withAccounts *AccountQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthRoleQuery builder.
func (arq *AuthRoleQuery) Where(ps ...predicate.AuthRole) *AuthRoleQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit adds a limit step to the query.
func (arq *AuthRoleQuery) Limit(limit int) *AuthRoleQuery {
	arq.limit = &limit
	return arq
}

// Offset adds an offset step to the query.
func (arq *AuthRoleQuery) Offset(offset int) *AuthRoleQuery {
	arq.offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AuthRoleQuery) Unique(unique bool) *AuthRoleQuery {
	arq.unique = &unique
	return arq
}

// Order adds an order step to the query.
func (arq *AuthRoleQuery) Order(o ...OrderFunc) *AuthRoleQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// QueryAccounts chains the current query on the "accounts" edge.
func (arq *AuthRoleQuery) QueryAccounts() *AccountQuery {
	query := &AccountQuery{config: arq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := arq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authrole.Table, authrole.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, authrole.AccountsTable, authrole.AccountsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(arq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AuthRole entity from the query.
// Returns a *NotFoundError when no AuthRole was found.
func (arq *AuthRoleQuery) First(ctx context.Context) (*AuthRole, error) {
	nodes, err := arq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AuthRoleQuery) FirstX(ctx context.Context) *AuthRole {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthRole ID from the query.
// Returns a *NotFoundError when no AuthRole ID was found.
func (arq *AuthRoleQuery) FirstID(ctx context.Context) (id ulid.ULID, err error) {
	var ids []ulid.ULID
	if ids, err = arq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AuthRoleQuery) FirstIDX(ctx context.Context) ulid.ULID {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AuthRole entity is found.
// Returns a *NotFoundError when no AuthRole entities are found.
func (arq *AuthRoleQuery) Only(ctx context.Context) (*AuthRole, error) {
	nodes, err := arq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authrole.Label}
	default:
		return nil, &NotSingularError{authrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AuthRoleQuery) OnlyX(ctx context.Context) *AuthRole {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthRole ID in the query.
// Returns a *NotSingularError when more than one AuthRole ID is found.
// Returns a *NotFoundError when no entities are found.
func (arq *AuthRoleQuery) OnlyID(ctx context.Context) (id ulid.ULID, err error) {
	var ids []ulid.ULID
	if ids, err = arq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authrole.Label}
	default:
		err = &NotSingularError{authrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AuthRoleQuery) OnlyIDX(ctx context.Context) ulid.ULID {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthRoles.
func (arq *AuthRoleQuery) All(ctx context.Context) ([]*AuthRole, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return arq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (arq *AuthRoleQuery) AllX(ctx context.Context) []*AuthRole {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthRole IDs.
func (arq *AuthRoleQuery) IDs(ctx context.Context) ([]ulid.ULID, error) {
	var ids []ulid.ULID
	if err := arq.Select(authrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AuthRoleQuery) IDsX(ctx context.Context) []ulid.ULID {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AuthRoleQuery) Count(ctx context.Context) (int, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return arq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AuthRoleQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AuthRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return arq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AuthRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AuthRoleQuery) Clone() *AuthRoleQuery {
	if arq == nil {
		return nil
	}
	return &AuthRoleQuery{
		config:       arq.config,
		limit:        arq.limit,
		offset:       arq.offset,
		order:        append([]OrderFunc{}, arq.order...),
		predicates:   append([]predicate.AuthRole{}, arq.predicates...),
		withAccounts: arq.withAccounts.Clone(),
		// clone intermediate query.
		sql:    arq.sql.Clone(),
		path:   arq.path,
		unique: arq.unique,
	}
}

// WithAccounts tells the query-builder to eager-load the nodes that are connected to
// the "accounts" edge. The optional arguments are used to configure the query builder of the edge.
func (arq *AuthRoleQuery) WithAccounts(opts ...func(*AccountQuery)) *AuthRoleQuery {
	query := &AccountQuery{config: arq.config}
	for _, opt := range opts {
		opt(query)
	}
	arq.withAccounts = query
	return arq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthRole.Query().
//		GroupBy(authrole.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (arq *AuthRoleQuery) GroupBy(field string, fields ...string) *AuthRoleGroupBy {
	grbuild := &AuthRoleGroupBy{config: arq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return arq.sqlQuery(ctx), nil
	}
	grbuild.label = authrole.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.AuthRole.Query().
//		Select(authrole.FieldCreatedAt).
//		Scan(ctx, &v)
func (arq *AuthRoleQuery) Select(fields ...string) *AuthRoleSelect {
	arq.fields = append(arq.fields, fields...)
	selbuild := &AuthRoleSelect{AuthRoleQuery: arq}
	selbuild.label = authrole.Label
	selbuild.flds, selbuild.scan = &arq.fields, selbuild.Scan
	return selbuild
}

func (arq *AuthRoleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range arq.fields {
		if !authrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AuthRoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AuthRole, error) {
	var (
		nodes       = []*AuthRole{}
		_spec       = arq.querySpec()
		loadedTypes = [1]bool{
			arq.withAccounts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AuthRole).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AuthRole{config: arq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := arq.withAccounts; query != nil {
		if err := arq.loadAccounts(ctx, query, nodes,
			func(n *AuthRole) { n.Edges.Accounts = []*Account{} },
			func(n *AuthRole, e *Account) { n.Edges.Accounts = append(n.Edges.Accounts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (arq *AuthRoleQuery) loadAccounts(ctx context.Context, query *AccountQuery, nodes []*AuthRole, init func(*AuthRole), assign func(*AuthRole, *Account)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[ulid.ULID]*AuthRole)
	nids := make(map[ulid.ULID]map[*AuthRole]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(authrole.AccountsTable)
		s.Join(joinT).On(s.C(account.FieldID), joinT.C(authrole.AccountsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(authrole.AccountsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(authrole.AccountsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(ulid.ULID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := *values[0].(*ulid.ULID)
			inValue := *values[1].(*ulid.ULID)
			if nids[inValue] == nil {
				nids[inValue] = map[*AuthRole]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "accounts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (arq *AuthRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	_spec.Node.Columns = arq.fields
	if len(arq.fields) > 0 {
		_spec.Unique = arq.unique != nil && *arq.unique
	}
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AuthRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := arq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("generated: check existence: %w", err)
	}
	return n > 0, nil
}

func (arq *AuthRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authrole.Table,
			Columns: authrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: authrole.FieldID,
			},
		},
		From:   arq.sql,
		Unique: true,
	}
	if unique := arq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := arq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authrole.FieldID)
		for i := range fields {
			if fields[i] != authrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AuthRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(authrole.Table)
	columns := arq.fields
	if len(columns) == 0 {
		columns = authrole.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if arq.unique != nil && *arq.unique {
		selector.Distinct()
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthRoleGroupBy is the group-by builder for AuthRole entities.
type AuthRoleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AuthRoleGroupBy) Aggregate(fns ...AggregateFunc) *AuthRoleGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the group-by query and scans the result into the given value.
func (argb *AuthRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := argb.path(ctx)
	if err != nil {
		return err
	}
	argb.sql = query
	return argb.sqlScan(ctx, v)
}

func (argb *AuthRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range argb.fields {
		if !authrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := argb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (argb *AuthRoleGroupBy) sqlQuery() *sql.Selector {
	selector := argb.sql.Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(argb.fields)+len(argb.fns))
		for _, f := range argb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(argb.fields...)...)
}

// AuthRoleSelect is the builder for selecting fields of AuthRole entities.
type AuthRoleSelect struct {
	*AuthRoleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AuthRoleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	ars.sql = ars.AuthRoleQuery.sqlQuery(ctx)
	return ars.sqlScan(ctx, v)
}

func (ars *AuthRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ars.sql.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}