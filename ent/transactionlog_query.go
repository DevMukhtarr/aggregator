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
	"github.com/google/uuid"
	"github.com/paycrest/aggregator/ent/predicate"
	"github.com/paycrest/aggregator/ent/transactionlog"
)

// TransactionLogQuery is the builder for querying TransactionLog entities.
type TransactionLogQuery struct {
	config
	ctx        *QueryContext
	order      []transactionlog.OrderOption
	inters     []Interceptor
	predicates []predicate.TransactionLog
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TransactionLogQuery builder.
func (tlq *TransactionLogQuery) Where(ps ...predicate.TransactionLog) *TransactionLogQuery {
	tlq.predicates = append(tlq.predicates, ps...)
	return tlq
}

// Limit the number of records to be returned by this query.
func (tlq *TransactionLogQuery) Limit(limit int) *TransactionLogQuery {
	tlq.ctx.Limit = &limit
	return tlq
}

// Offset to start from.
func (tlq *TransactionLogQuery) Offset(offset int) *TransactionLogQuery {
	tlq.ctx.Offset = &offset
	return tlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tlq *TransactionLogQuery) Unique(unique bool) *TransactionLogQuery {
	tlq.ctx.Unique = &unique
	return tlq
}

// Order specifies how the records should be ordered.
func (tlq *TransactionLogQuery) Order(o ...transactionlog.OrderOption) *TransactionLogQuery {
	tlq.order = append(tlq.order, o...)
	return tlq
}

// First returns the first TransactionLog entity from the query.
// Returns a *NotFoundError when no TransactionLog was found.
func (tlq *TransactionLogQuery) First(ctx context.Context) (*TransactionLog, error) {
	nodes, err := tlq.Limit(1).All(setContextOp(ctx, tlq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transactionlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tlq *TransactionLogQuery) FirstX(ctx context.Context) *TransactionLog {
	node, err := tlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TransactionLog ID from the query.
// Returns a *NotFoundError when no TransactionLog ID was found.
func (tlq *TransactionLogQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tlq.Limit(1).IDs(setContextOp(ctx, tlq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transactionlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tlq *TransactionLogQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TransactionLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TransactionLog entity is found.
// Returns a *NotFoundError when no TransactionLog entities are found.
func (tlq *TransactionLogQuery) Only(ctx context.Context) (*TransactionLog, error) {
	nodes, err := tlq.Limit(2).All(setContextOp(ctx, tlq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transactionlog.Label}
	default:
		return nil, &NotSingularError{transactionlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tlq *TransactionLogQuery) OnlyX(ctx context.Context) *TransactionLog {
	node, err := tlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TransactionLog ID in the query.
// Returns a *NotSingularError when more than one TransactionLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (tlq *TransactionLogQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tlq.Limit(2).IDs(setContextOp(ctx, tlq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transactionlog.Label}
	default:
		err = &NotSingularError{transactionlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tlq *TransactionLogQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TransactionLogs.
func (tlq *TransactionLogQuery) All(ctx context.Context) ([]*TransactionLog, error) {
	ctx = setContextOp(ctx, tlq.ctx, ent.OpQueryAll)
	if err := tlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TransactionLog, *TransactionLogQuery]()
	return withInterceptors[[]*TransactionLog](ctx, tlq, qr, tlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tlq *TransactionLogQuery) AllX(ctx context.Context) []*TransactionLog {
	nodes, err := tlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TransactionLog IDs.
func (tlq *TransactionLogQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tlq.ctx.Unique == nil && tlq.path != nil {
		tlq.Unique(true)
	}
	ctx = setContextOp(ctx, tlq.ctx, ent.OpQueryIDs)
	if err = tlq.Select(transactionlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tlq *TransactionLogQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tlq *TransactionLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tlq.ctx, ent.OpQueryCount)
	if err := tlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tlq, querierCount[*TransactionLogQuery](), tlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tlq *TransactionLogQuery) CountX(ctx context.Context) int {
	count, err := tlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tlq *TransactionLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tlq.ctx, ent.OpQueryExist)
	switch _, err := tlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tlq *TransactionLogQuery) ExistX(ctx context.Context) bool {
	exist, err := tlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TransactionLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tlq *TransactionLogQuery) Clone() *TransactionLogQuery {
	if tlq == nil {
		return nil
	}
	return &TransactionLogQuery{
		config:     tlq.config,
		ctx:        tlq.ctx.Clone(),
		order:      append([]transactionlog.OrderOption{}, tlq.order...),
		inters:     append([]Interceptor{}, tlq.inters...),
		predicates: append([]predicate.TransactionLog{}, tlq.predicates...),
		// clone intermediate query.
		sql:  tlq.sql.Clone(),
		path: tlq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GatewayID string `json:"gateway_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TransactionLog.Query().
//		GroupBy(transactionlog.FieldGatewayID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tlq *TransactionLogQuery) GroupBy(field string, fields ...string) *TransactionLogGroupBy {
	tlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TransactionLogGroupBy{build: tlq}
	grbuild.flds = &tlq.ctx.Fields
	grbuild.label = transactionlog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GatewayID string `json:"gateway_id,omitempty"`
//	}
//
//	client.TransactionLog.Query().
//		Select(transactionlog.FieldGatewayID).
//		Scan(ctx, &v)
func (tlq *TransactionLogQuery) Select(fields ...string) *TransactionLogSelect {
	tlq.ctx.Fields = append(tlq.ctx.Fields, fields...)
	sbuild := &TransactionLogSelect{TransactionLogQuery: tlq}
	sbuild.label = transactionlog.Label
	sbuild.flds, sbuild.scan = &tlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TransactionLogSelect configured with the given aggregations.
func (tlq *TransactionLogQuery) Aggregate(fns ...AggregateFunc) *TransactionLogSelect {
	return tlq.Select().Aggregate(fns...)
}

func (tlq *TransactionLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tlq); err != nil {
				return err
			}
		}
	}
	for _, f := range tlq.ctx.Fields {
		if !transactionlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tlq.path != nil {
		prev, err := tlq.path(ctx)
		if err != nil {
			return err
		}
		tlq.sql = prev
	}
	return nil
}

func (tlq *TransactionLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TransactionLog, error) {
	var (
		nodes   = []*TransactionLog{}
		withFKs = tlq.withFKs
		_spec   = tlq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, transactionlog.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TransactionLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TransactionLog{config: tlq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tlq *TransactionLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tlq.querySpec()
	_spec.Node.Columns = tlq.ctx.Fields
	if len(tlq.ctx.Fields) > 0 {
		_spec.Unique = tlq.ctx.Unique != nil && *tlq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tlq.driver, _spec)
}

func (tlq *TransactionLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(transactionlog.Table, transactionlog.Columns, sqlgraph.NewFieldSpec(transactionlog.FieldID, field.TypeUUID))
	_spec.From = tlq.sql
	if unique := tlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tlq.path != nil {
		_spec.Unique = true
	}
	if fields := tlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transactionlog.FieldID)
		for i := range fields {
			if fields[i] != transactionlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tlq *TransactionLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tlq.driver.Dialect())
	t1 := builder.Table(transactionlog.Table)
	columns := tlq.ctx.Fields
	if len(columns) == 0 {
		columns = transactionlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tlq.sql != nil {
		selector = tlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tlq.ctx.Unique != nil && *tlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tlq.predicates {
		p(selector)
	}
	for _, p := range tlq.order {
		p(selector)
	}
	if offset := tlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TransactionLogGroupBy is the group-by builder for TransactionLog entities.
type TransactionLogGroupBy struct {
	selector
	build *TransactionLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tlgb *TransactionLogGroupBy) Aggregate(fns ...AggregateFunc) *TransactionLogGroupBy {
	tlgb.fns = append(tlgb.fns, fns...)
	return tlgb
}

// Scan applies the selector query and scans the result into the given value.
func (tlgb *TransactionLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tlgb.build.ctx, ent.OpQueryGroupBy)
	if err := tlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TransactionLogQuery, *TransactionLogGroupBy](ctx, tlgb.build, tlgb, tlgb.build.inters, v)
}

func (tlgb *TransactionLogGroupBy) sqlScan(ctx context.Context, root *TransactionLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tlgb.fns))
	for _, fn := range tlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tlgb.flds)+len(tlgb.fns))
		for _, f := range *tlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TransactionLogSelect is the builder for selecting fields of TransactionLog entities.
type TransactionLogSelect struct {
	*TransactionLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tls *TransactionLogSelect) Aggregate(fns ...AggregateFunc) *TransactionLogSelect {
	tls.fns = append(tls.fns, fns...)
	return tls
}

// Scan applies the selector query and scans the result into the given value.
func (tls *TransactionLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tls.ctx, ent.OpQuerySelect)
	if err := tls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TransactionLogQuery, *TransactionLogSelect](ctx, tls.TransactionLogQuery, tls, tls.inters, v)
}

func (tls *TransactionLogSelect) sqlScan(ctx context.Context, root *TransactionLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tls.fns))
	for _, fn := range tls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
