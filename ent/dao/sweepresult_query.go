// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/predicate"
	"aisecurity/ent/dao/sweep"
	"aisecurity/ent/dao/sweepresult"
	"aisecurity/ent/dao/sweepresultdetails"
	"aisecurity/ent/dao/sweepschedule"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SweepResultQuery is the builder for querying SweepResult entities.
type SweepResultQuery struct {
	config
	ctx                    *QueryContext
	order                  []sweepresult.OrderOption
	inters                 []Interceptor
	predicates             []predicate.SweepResult
	withCreator            *AdminQuery
	withUpdater            *AdminQuery
	withSweep              *SweepQuery
	withSweepSchedule      *SweepScheduleQuery
	withSweepResultDetails *SweepResultDetailsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SweepResultQuery builder.
func (srq *SweepResultQuery) Where(ps ...predicate.SweepResult) *SweepResultQuery {
	srq.predicates = append(srq.predicates, ps...)
	return srq
}

// Limit the number of records to be returned by this query.
func (srq *SweepResultQuery) Limit(limit int) *SweepResultQuery {
	srq.ctx.Limit = &limit
	return srq
}

// Offset to start from.
func (srq *SweepResultQuery) Offset(offset int) *SweepResultQuery {
	srq.ctx.Offset = &offset
	return srq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (srq *SweepResultQuery) Unique(unique bool) *SweepResultQuery {
	srq.ctx.Unique = &unique
	return srq
}

// Order specifies how the records should be ordered.
func (srq *SweepResultQuery) Order(o ...sweepresult.OrderOption) *SweepResultQuery {
	srq.order = append(srq.order, o...)
	return srq
}

// QueryCreator chains the current query on the "creator" edge.
func (srq *SweepResultQuery) QueryCreator() *AdminQuery {
	query := (&AdminClient{config: srq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := srq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweepresult.Table, sweepresult.FieldID, selector),
			sqlgraph.To(admin.Table, admin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sweepresult.CreatorTable, sweepresult.CreatorColumn),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUpdater chains the current query on the "updater" edge.
func (srq *SweepResultQuery) QueryUpdater() *AdminQuery {
	query := (&AdminClient{config: srq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := srq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweepresult.Table, sweepresult.FieldID, selector),
			sqlgraph.To(admin.Table, admin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sweepresult.UpdaterTable, sweepresult.UpdaterColumn),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySweep chains the current query on the "sweep" edge.
func (srq *SweepResultQuery) QuerySweep() *SweepQuery {
	query := (&SweepClient{config: srq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := srq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweepresult.Table, sweepresult.FieldID, selector),
			sqlgraph.To(sweep.Table, sweep.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sweepresult.SweepTable, sweepresult.SweepColumn),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySweepSchedule chains the current query on the "sweep_schedule" edge.
func (srq *SweepResultQuery) QuerySweepSchedule() *SweepScheduleQuery {
	query := (&SweepScheduleClient{config: srq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := srq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweepresult.Table, sweepresult.FieldID, selector),
			sqlgraph.To(sweepschedule.Table, sweepschedule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sweepresult.SweepScheduleTable, sweepresult.SweepScheduleColumn),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySweepResultDetails chains the current query on the "sweep_result_details" edge.
func (srq *SweepResultQuery) QuerySweepResultDetails() *SweepResultDetailsQuery {
	query := (&SweepResultDetailsClient{config: srq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := srq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweepresult.Table, sweepresult.FieldID, selector),
			sqlgraph.To(sweepresultdetails.Table, sweepresultdetails.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sweepresult.SweepResultDetailsTable, sweepresult.SweepResultDetailsColumn),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SweepResult entity from the query.
// Returns a *NotFoundError when no SweepResult was found.
func (srq *SweepResultQuery) First(ctx context.Context) (*SweepResult, error) {
	nodes, err := srq.Limit(1).All(setContextOp(ctx, srq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sweepresult.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (srq *SweepResultQuery) FirstX(ctx context.Context) *SweepResult {
	node, err := srq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SweepResult ID from the query.
// Returns a *NotFoundError when no SweepResult ID was found.
func (srq *SweepResultQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = srq.Limit(1).IDs(setContextOp(ctx, srq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sweepresult.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (srq *SweepResultQuery) FirstIDX(ctx context.Context) int {
	id, err := srq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SweepResult entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SweepResult entity is found.
// Returns a *NotFoundError when no SweepResult entities are found.
func (srq *SweepResultQuery) Only(ctx context.Context) (*SweepResult, error) {
	nodes, err := srq.Limit(2).All(setContextOp(ctx, srq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sweepresult.Label}
	default:
		return nil, &NotSingularError{sweepresult.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (srq *SweepResultQuery) OnlyX(ctx context.Context) *SweepResult {
	node, err := srq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SweepResult ID in the query.
// Returns a *NotSingularError when more than one SweepResult ID is found.
// Returns a *NotFoundError when no entities are found.
func (srq *SweepResultQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = srq.Limit(2).IDs(setContextOp(ctx, srq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sweepresult.Label}
	default:
		err = &NotSingularError{sweepresult.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (srq *SweepResultQuery) OnlyIDX(ctx context.Context) int {
	id, err := srq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SweepResults.
func (srq *SweepResultQuery) All(ctx context.Context) ([]*SweepResult, error) {
	ctx = setContextOp(ctx, srq.ctx, "All")
	if err := srq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SweepResult, *SweepResultQuery]()
	return withInterceptors[[]*SweepResult](ctx, srq, qr, srq.inters)
}

// AllX is like All, but panics if an error occurs.
func (srq *SweepResultQuery) AllX(ctx context.Context) []*SweepResult {
	nodes, err := srq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SweepResult IDs.
func (srq *SweepResultQuery) IDs(ctx context.Context) (ids []int, err error) {
	if srq.ctx.Unique == nil && srq.path != nil {
		srq.Unique(true)
	}
	ctx = setContextOp(ctx, srq.ctx, "IDs")
	if err = srq.Select(sweepresult.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (srq *SweepResultQuery) IDsX(ctx context.Context) []int {
	ids, err := srq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (srq *SweepResultQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, srq.ctx, "Count")
	if err := srq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, srq, querierCount[*SweepResultQuery](), srq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (srq *SweepResultQuery) CountX(ctx context.Context) int {
	count, err := srq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (srq *SweepResultQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, srq.ctx, "Exist")
	switch _, err := srq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("dao: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (srq *SweepResultQuery) ExistX(ctx context.Context) bool {
	exist, err := srq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SweepResultQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (srq *SweepResultQuery) Clone() *SweepResultQuery {
	if srq == nil {
		return nil
	}
	return &SweepResultQuery{
		config:                 srq.config,
		ctx:                    srq.ctx.Clone(),
		order:                  append([]sweepresult.OrderOption{}, srq.order...),
		inters:                 append([]Interceptor{}, srq.inters...),
		predicates:             append([]predicate.SweepResult{}, srq.predicates...),
		withCreator:            srq.withCreator.Clone(),
		withUpdater:            srq.withUpdater.Clone(),
		withSweep:              srq.withSweep.Clone(),
		withSweepSchedule:      srq.withSweepSchedule.Clone(),
		withSweepResultDetails: srq.withSweepResultDetails.Clone(),
		// clone intermediate query.
		sql:  srq.sql.Clone(),
		path: srq.path,
	}
}

// WithCreator tells the query-builder to eager-load the nodes that are connected to
// the "creator" edge. The optional arguments are used to configure the query builder of the edge.
func (srq *SweepResultQuery) WithCreator(opts ...func(*AdminQuery)) *SweepResultQuery {
	query := (&AdminClient{config: srq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	srq.withCreator = query
	return srq
}

// WithUpdater tells the query-builder to eager-load the nodes that are connected to
// the "updater" edge. The optional arguments are used to configure the query builder of the edge.
func (srq *SweepResultQuery) WithUpdater(opts ...func(*AdminQuery)) *SweepResultQuery {
	query := (&AdminClient{config: srq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	srq.withUpdater = query
	return srq
}

// WithSweep tells the query-builder to eager-load the nodes that are connected to
// the "sweep" edge. The optional arguments are used to configure the query builder of the edge.
func (srq *SweepResultQuery) WithSweep(opts ...func(*SweepQuery)) *SweepResultQuery {
	query := (&SweepClient{config: srq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	srq.withSweep = query
	return srq
}

// WithSweepSchedule tells the query-builder to eager-load the nodes that are connected to
// the "sweep_schedule" edge. The optional arguments are used to configure the query builder of the edge.
func (srq *SweepResultQuery) WithSweepSchedule(opts ...func(*SweepScheduleQuery)) *SweepResultQuery {
	query := (&SweepScheduleClient{config: srq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	srq.withSweepSchedule = query
	return srq
}

// WithSweepResultDetails tells the query-builder to eager-load the nodes that are connected to
// the "sweep_result_details" edge. The optional arguments are used to configure the query builder of the edge.
func (srq *SweepResultQuery) WithSweepResultDetails(opts ...func(*SweepResultDetailsQuery)) *SweepResultQuery {
	query := (&SweepResultDetailsClient{config: srq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	srq.withSweepResultDetails = query
	return srq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SweepResult.Query().
//		GroupBy(sweepresult.FieldCreateTime).
//		Aggregate(dao.Count()).
//		Scan(ctx, &v)
func (srq *SweepResultQuery) GroupBy(field string, fields ...string) *SweepResultGroupBy {
	srq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SweepResultGroupBy{build: srq}
	grbuild.flds = &srq.ctx.Fields
	grbuild.label = sweepresult.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time"`
//	}
//
//	client.SweepResult.Query().
//		Select(sweepresult.FieldCreateTime).
//		Scan(ctx, &v)
func (srq *SweepResultQuery) Select(fields ...string) *SweepResultSelect {
	srq.ctx.Fields = append(srq.ctx.Fields, fields...)
	sbuild := &SweepResultSelect{SweepResultQuery: srq}
	sbuild.label = sweepresult.Label
	sbuild.flds, sbuild.scan = &srq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SweepResultSelect configured with the given aggregations.
func (srq *SweepResultQuery) Aggregate(fns ...AggregateFunc) *SweepResultSelect {
	return srq.Select().Aggregate(fns...)
}

func (srq *SweepResultQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range srq.inters {
		if inter == nil {
			return fmt.Errorf("dao: uninitialized interceptor (forgotten import dao/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, srq); err != nil {
				return err
			}
		}
	}
	for _, f := range srq.ctx.Fields {
		if !sweepresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
		}
	}
	if srq.path != nil {
		prev, err := srq.path(ctx)
		if err != nil {
			return err
		}
		srq.sql = prev
	}
	return nil
}

func (srq *SweepResultQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SweepResult, error) {
	var (
		nodes       = []*SweepResult{}
		_spec       = srq.querySpec()
		loadedTypes = [5]bool{
			srq.withCreator != nil,
			srq.withUpdater != nil,
			srq.withSweep != nil,
			srq.withSweepSchedule != nil,
			srq.withSweepResultDetails != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SweepResult).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SweepResult{config: srq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, srq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := srq.withCreator; query != nil {
		if err := srq.loadCreator(ctx, query, nodes, nil,
			func(n *SweepResult, e *Admin) { n.Edges.Creator = e }); err != nil {
			return nil, err
		}
	}
	if query := srq.withUpdater; query != nil {
		if err := srq.loadUpdater(ctx, query, nodes, nil,
			func(n *SweepResult, e *Admin) { n.Edges.Updater = e }); err != nil {
			return nil, err
		}
	}
	if query := srq.withSweep; query != nil {
		if err := srq.loadSweep(ctx, query, nodes, nil,
			func(n *SweepResult, e *Sweep) { n.Edges.Sweep = e }); err != nil {
			return nil, err
		}
	}
	if query := srq.withSweepSchedule; query != nil {
		if err := srq.loadSweepSchedule(ctx, query, nodes, nil,
			func(n *SweepResult, e *SweepSchedule) { n.Edges.SweepSchedule = e }); err != nil {
			return nil, err
		}
	}
	if query := srq.withSweepResultDetails; query != nil {
		if err := srq.loadSweepResultDetails(ctx, query, nodes,
			func(n *SweepResult) { n.Edges.SweepResultDetails = []*SweepResultDetails{} },
			func(n *SweepResult, e *SweepResultDetails) {
				n.Edges.SweepResultDetails = append(n.Edges.SweepResultDetails, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (srq *SweepResultQuery) loadCreator(ctx context.Context, query *AdminQuery, nodes []*SweepResult, init func(*SweepResult), assign func(*SweepResult, *Admin)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SweepResult)
	for i := range nodes {
		fk := nodes[i].CreatorID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(admin.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "creator_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (srq *SweepResultQuery) loadUpdater(ctx context.Context, query *AdminQuery, nodes []*SweepResult, init func(*SweepResult), assign func(*SweepResult, *Admin)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SweepResult)
	for i := range nodes {
		fk := nodes[i].UpdaterID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(admin.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "updater_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (srq *SweepResultQuery) loadSweep(ctx context.Context, query *SweepQuery, nodes []*SweepResult, init func(*SweepResult), assign func(*SweepResult, *Sweep)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SweepResult)
	for i := range nodes {
		fk := nodes[i].SweepID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(sweep.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "sweep_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (srq *SweepResultQuery) loadSweepSchedule(ctx context.Context, query *SweepScheduleQuery, nodes []*SweepResult, init func(*SweepResult), assign func(*SweepResult, *SweepSchedule)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SweepResult)
	for i := range nodes {
		fk := nodes[i].SweepScheduleID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(sweepschedule.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "sweep_schedule_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (srq *SweepResultQuery) loadSweepResultDetails(ctx context.Context, query *SweepResultDetailsQuery, nodes []*SweepResult, init func(*SweepResult), assign func(*SweepResult, *SweepResultDetails)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*SweepResult)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(sweepresultdetails.FieldSweepResultID)
	}
	query.Where(predicate.SweepResultDetails(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sweepresult.SweepResultDetailsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SweepResultID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "sweep_result_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (srq *SweepResultQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := srq.querySpec()
	_spec.Node.Columns = srq.ctx.Fields
	if len(srq.ctx.Fields) > 0 {
		_spec.Unique = srq.ctx.Unique != nil && *srq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, srq.driver, _spec)
}

func (srq *SweepResultQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sweepresult.Table, sweepresult.Columns, sqlgraph.NewFieldSpec(sweepresult.FieldID, field.TypeInt))
	_spec.From = srq.sql
	if unique := srq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if srq.path != nil {
		_spec.Unique = true
	}
	if fields := srq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sweepresult.FieldID)
		for i := range fields {
			if fields[i] != sweepresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if srq.withCreator != nil {
			_spec.Node.AddColumnOnce(sweepresult.FieldCreatorID)
		}
		if srq.withUpdater != nil {
			_spec.Node.AddColumnOnce(sweepresult.FieldUpdaterID)
		}
		if srq.withSweep != nil {
			_spec.Node.AddColumnOnce(sweepresult.FieldSweepID)
		}
		if srq.withSweepSchedule != nil {
			_spec.Node.AddColumnOnce(sweepresult.FieldSweepScheduleID)
		}
	}
	if ps := srq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := srq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := srq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := srq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (srq *SweepResultQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(srq.driver.Dialect())
	t1 := builder.Table(sweepresult.Table)
	columns := srq.ctx.Fields
	if len(columns) == 0 {
		columns = sweepresult.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if srq.sql != nil {
		selector = srq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if srq.ctx.Unique != nil && *srq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range srq.predicates {
		p(selector)
	}
	for _, p := range srq.order {
		p(selector)
	}
	if offset := srq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := srq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SweepResultGroupBy is the group-by builder for SweepResult entities.
type SweepResultGroupBy struct {
	selector
	build *SweepResultQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (srgb *SweepResultGroupBy) Aggregate(fns ...AggregateFunc) *SweepResultGroupBy {
	srgb.fns = append(srgb.fns, fns...)
	return srgb
}

// Scan applies the selector query and scans the result into the given value.
func (srgb *SweepResultGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, srgb.build.ctx, "GroupBy")
	if err := srgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SweepResultQuery, *SweepResultGroupBy](ctx, srgb.build, srgb, srgb.build.inters, v)
}

func (srgb *SweepResultGroupBy) sqlScan(ctx context.Context, root *SweepResultQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(srgb.fns))
	for _, fn := range srgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*srgb.flds)+len(srgb.fns))
		for _, f := range *srgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*srgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SweepResultSelect is the builder for selecting fields of SweepResult entities.
type SweepResultSelect struct {
	*SweepResultQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (srs *SweepResultSelect) Aggregate(fns ...AggregateFunc) *SweepResultSelect {
	srs.fns = append(srs.fns, fns...)
	return srs
}

// Scan applies the selector query and scans the result into the given value.
func (srs *SweepResultSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, srs.ctx, "Select")
	if err := srs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SweepResultQuery, *SweepResultSelect](ctx, srs.SweepResultQuery, srs, srs.inters, v)
}

func (srs *SweepResultSelect) sqlScan(ctx context.Context, root *SweepResultQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(srs.fns))
	for _, fn := range srs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*srs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}