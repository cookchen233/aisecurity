// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/predicate"
	"aisecurity/ent/dao/riskcategory"
	"aisecurity/ent/dao/risklocation"
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

// SweepQuery is the builder for querying Sweep entities.
type SweepQuery struct {
	config
	ctx                    *QueryContext
	order                  []sweep.OrderOption
	inters                 []Interceptor
	predicates             []predicate.Sweep
	withCreator            *AdminQuery
	withUpdater            *AdminQuery
	withRiskCategory       *RiskCategoryQuery
	withRiskLocation       *RiskLocationQuery
	withSweepSchedule      *SweepScheduleQuery
	withSweepResult        *SweepResultQuery
	withSweepResultDetails *SweepResultDetailsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SweepQuery builder.
func (sq *SweepQuery) Where(ps ...predicate.Sweep) *SweepQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SweepQuery) Limit(limit int) *SweepQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SweepQuery) Offset(offset int) *SweepQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SweepQuery) Unique(unique bool) *SweepQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SweepQuery) Order(o ...sweep.OrderOption) *SweepQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryCreator chains the current query on the "creator" edge.
func (sq *SweepQuery) QueryCreator() *AdminQuery {
	query := (&AdminClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweep.Table, sweep.FieldID, selector),
			sqlgraph.To(admin.Table, admin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sweep.CreatorTable, sweep.CreatorColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUpdater chains the current query on the "updater" edge.
func (sq *SweepQuery) QueryUpdater() *AdminQuery {
	query := (&AdminClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweep.Table, sweep.FieldID, selector),
			sqlgraph.To(admin.Table, admin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sweep.UpdaterTable, sweep.UpdaterColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRiskCategory chains the current query on the "risk_category" edge.
func (sq *SweepQuery) QueryRiskCategory() *RiskCategoryQuery {
	query := (&RiskCategoryClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweep.Table, sweep.FieldID, selector),
			sqlgraph.To(riskcategory.Table, riskcategory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sweep.RiskCategoryTable, sweep.RiskCategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRiskLocation chains the current query on the "risk_location" edge.
func (sq *SweepQuery) QueryRiskLocation() *RiskLocationQuery {
	query := (&RiskLocationClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweep.Table, sweep.FieldID, selector),
			sqlgraph.To(risklocation.Table, risklocation.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sweep.RiskLocationTable, sweep.RiskLocationColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySweepSchedule chains the current query on the "sweep_schedule" edge.
func (sq *SweepQuery) QuerySweepSchedule() *SweepScheduleQuery {
	query := (&SweepScheduleClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweep.Table, sweep.FieldID, selector),
			sqlgraph.To(sweepschedule.Table, sweepschedule.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sweep.SweepScheduleTable, sweep.SweepScheduleColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySweepResult chains the current query on the "sweep_result" edge.
func (sq *SweepQuery) QuerySweepResult() *SweepResultQuery {
	query := (&SweepResultClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweep.Table, sweep.FieldID, selector),
			sqlgraph.To(sweepresult.Table, sweepresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sweep.SweepResultTable, sweep.SweepResultColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySweepResultDetails chains the current query on the "sweep_result_details" edge.
func (sq *SweepQuery) QuerySweepResultDetails() *SweepResultDetailsQuery {
	query := (&SweepResultDetailsClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sweep.Table, sweep.FieldID, selector),
			sqlgraph.To(sweepresultdetails.Table, sweepresultdetails.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sweep.SweepResultDetailsTable, sweep.SweepResultDetailsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Sweep entity from the query.
// Returns a *NotFoundError when no Sweep was found.
func (sq *SweepQuery) First(ctx context.Context) (*Sweep, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sweep.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SweepQuery) FirstX(ctx context.Context) *Sweep {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Sweep ID from the query.
// Returns a *NotFoundError when no Sweep ID was found.
func (sq *SweepQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sweep.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SweepQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Sweep entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Sweep entity is found.
// Returns a *NotFoundError when no Sweep entities are found.
func (sq *SweepQuery) Only(ctx context.Context) (*Sweep, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sweep.Label}
	default:
		return nil, &NotSingularError{sweep.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SweepQuery) OnlyX(ctx context.Context) *Sweep {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Sweep ID in the query.
// Returns a *NotSingularError when more than one Sweep ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SweepQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sweep.Label}
	default:
		err = &NotSingularError{sweep.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SweepQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Sweeps.
func (sq *SweepQuery) All(ctx context.Context) ([]*Sweep, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Sweep, *SweepQuery]()
	return withInterceptors[[]*Sweep](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SweepQuery) AllX(ctx context.Context) []*Sweep {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Sweep IDs.
func (sq *SweepQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(sweep.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SweepQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SweepQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SweepQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SweepQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SweepQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("dao: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SweepQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SweepQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SweepQuery) Clone() *SweepQuery {
	if sq == nil {
		return nil
	}
	return &SweepQuery{
		config:                 sq.config,
		ctx:                    sq.ctx.Clone(),
		order:                  append([]sweep.OrderOption{}, sq.order...),
		inters:                 append([]Interceptor{}, sq.inters...),
		predicates:             append([]predicate.Sweep{}, sq.predicates...),
		withCreator:            sq.withCreator.Clone(),
		withUpdater:            sq.withUpdater.Clone(),
		withRiskCategory:       sq.withRiskCategory.Clone(),
		withRiskLocation:       sq.withRiskLocation.Clone(),
		withSweepSchedule:      sq.withSweepSchedule.Clone(),
		withSweepResult:        sq.withSweepResult.Clone(),
		withSweepResultDetails: sq.withSweepResultDetails.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithCreator tells the query-builder to eager-load the nodes that are connected to
// the "creator" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SweepQuery) WithCreator(opts ...func(*AdminQuery)) *SweepQuery {
	query := (&AdminClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withCreator = query
	return sq
}

// WithUpdater tells the query-builder to eager-load the nodes that are connected to
// the "updater" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SweepQuery) WithUpdater(opts ...func(*AdminQuery)) *SweepQuery {
	query := (&AdminClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withUpdater = query
	return sq
}

// WithRiskCategory tells the query-builder to eager-load the nodes that are connected to
// the "risk_category" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SweepQuery) WithRiskCategory(opts ...func(*RiskCategoryQuery)) *SweepQuery {
	query := (&RiskCategoryClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withRiskCategory = query
	return sq
}

// WithRiskLocation tells the query-builder to eager-load the nodes that are connected to
// the "risk_location" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SweepQuery) WithRiskLocation(opts ...func(*RiskLocationQuery)) *SweepQuery {
	query := (&RiskLocationClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withRiskLocation = query
	return sq
}

// WithSweepSchedule tells the query-builder to eager-load the nodes that are connected to
// the "sweep_schedule" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SweepQuery) WithSweepSchedule(opts ...func(*SweepScheduleQuery)) *SweepQuery {
	query := (&SweepScheduleClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withSweepSchedule = query
	return sq
}

// WithSweepResult tells the query-builder to eager-load the nodes that are connected to
// the "sweep_result" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SweepQuery) WithSweepResult(opts ...func(*SweepResultQuery)) *SweepQuery {
	query := (&SweepResultClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withSweepResult = query
	return sq
}

// WithSweepResultDetails tells the query-builder to eager-load the nodes that are connected to
// the "sweep_result_details" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SweepQuery) WithSweepResultDetails(opts ...func(*SweepResultDetailsQuery)) *SweepQuery {
	query := (&SweepResultDetailsClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withSweepResultDetails = query
	return sq
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
//	client.Sweep.Query().
//		GroupBy(sweep.FieldCreateTime).
//		Aggregate(dao.Count()).
//		Scan(ctx, &v)
func (sq *SweepQuery) GroupBy(field string, fields ...string) *SweepGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SweepGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = sweep.Label
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
//	client.Sweep.Query().
//		Select(sweep.FieldCreateTime).
//		Scan(ctx, &v)
func (sq *SweepQuery) Select(fields ...string) *SweepSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SweepSelect{SweepQuery: sq}
	sbuild.label = sweep.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SweepSelect configured with the given aggregations.
func (sq *SweepQuery) Aggregate(fns ...AggregateFunc) *SweepSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SweepQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("dao: uninitialized interceptor (forgotten import dao/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !sweep.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SweepQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Sweep, error) {
	var (
		nodes       = []*Sweep{}
		_spec       = sq.querySpec()
		loadedTypes = [7]bool{
			sq.withCreator != nil,
			sq.withUpdater != nil,
			sq.withRiskCategory != nil,
			sq.withRiskLocation != nil,
			sq.withSweepSchedule != nil,
			sq.withSweepResult != nil,
			sq.withSweepResultDetails != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Sweep).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Sweep{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withCreator; query != nil {
		if err := sq.loadCreator(ctx, query, nodes, nil,
			func(n *Sweep, e *Admin) { n.Edges.Creator = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withUpdater; query != nil {
		if err := sq.loadUpdater(ctx, query, nodes, nil,
			func(n *Sweep, e *Admin) { n.Edges.Updater = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withRiskCategory; query != nil {
		if err := sq.loadRiskCategory(ctx, query, nodes, nil,
			func(n *Sweep, e *RiskCategory) { n.Edges.RiskCategory = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withRiskLocation; query != nil {
		if err := sq.loadRiskLocation(ctx, query, nodes, nil,
			func(n *Sweep, e *RiskLocation) { n.Edges.RiskLocation = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withSweepSchedule; query != nil {
		if err := sq.loadSweepSchedule(ctx, query, nodes,
			func(n *Sweep) { n.Edges.SweepSchedule = []*SweepSchedule{} },
			func(n *Sweep, e *SweepSchedule) { n.Edges.SweepSchedule = append(n.Edges.SweepSchedule, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withSweepResult; query != nil {
		if err := sq.loadSweepResult(ctx, query, nodes,
			func(n *Sweep) { n.Edges.SweepResult = []*SweepResult{} },
			func(n *Sweep, e *SweepResult) { n.Edges.SweepResult = append(n.Edges.SweepResult, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withSweepResultDetails; query != nil {
		if err := sq.loadSweepResultDetails(ctx, query, nodes,
			func(n *Sweep) { n.Edges.SweepResultDetails = []*SweepResultDetails{} },
			func(n *Sweep, e *SweepResultDetails) {
				n.Edges.SweepResultDetails = append(n.Edges.SweepResultDetails, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SweepQuery) loadCreator(ctx context.Context, query *AdminQuery, nodes []*Sweep, init func(*Sweep), assign func(*Sweep, *Admin)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Sweep)
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
func (sq *SweepQuery) loadUpdater(ctx context.Context, query *AdminQuery, nodes []*Sweep, init func(*Sweep), assign func(*Sweep, *Admin)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Sweep)
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
func (sq *SweepQuery) loadRiskCategory(ctx context.Context, query *RiskCategoryQuery, nodes []*Sweep, init func(*Sweep), assign func(*Sweep, *RiskCategory)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Sweep)
	for i := range nodes {
		fk := nodes[i].RiskCategoryID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(riskcategory.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "risk_category_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *SweepQuery) loadRiskLocation(ctx context.Context, query *RiskLocationQuery, nodes []*Sweep, init func(*Sweep), assign func(*Sweep, *RiskLocation)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Sweep)
	for i := range nodes {
		fk := nodes[i].RiskLocationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(risklocation.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "risk_location_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *SweepQuery) loadSweepSchedule(ctx context.Context, query *SweepScheduleQuery, nodes []*Sweep, init func(*Sweep), assign func(*Sweep, *SweepSchedule)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Sweep)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(sweepschedule.FieldSweepID)
	}
	query.Where(predicate.SweepSchedule(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sweep.SweepScheduleColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SweepID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "sweep_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *SweepQuery) loadSweepResult(ctx context.Context, query *SweepResultQuery, nodes []*Sweep, init func(*Sweep), assign func(*Sweep, *SweepResult)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Sweep)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(sweepresult.FieldSweepID)
	}
	query.Where(predicate.SweepResult(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sweep.SweepResultColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SweepID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "sweep_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *SweepQuery) loadSweepResultDetails(ctx context.Context, query *SweepResultDetailsQuery, nodes []*Sweep, init func(*Sweep), assign func(*Sweep, *SweepResultDetails)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Sweep)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(sweepresultdetails.FieldSweepID)
	}
	query.Where(predicate.SweepResultDetails(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sweep.SweepResultDetailsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SweepID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "sweep_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sq *SweepQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SweepQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sweep.Table, sweep.Columns, sqlgraph.NewFieldSpec(sweep.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sweep.FieldID)
		for i := range fields {
			if fields[i] != sweep.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if sq.withCreator != nil {
			_spec.Node.AddColumnOnce(sweep.FieldCreatorID)
		}
		if sq.withUpdater != nil {
			_spec.Node.AddColumnOnce(sweep.FieldUpdaterID)
		}
		if sq.withRiskCategory != nil {
			_spec.Node.AddColumnOnce(sweep.FieldRiskCategoryID)
		}
		if sq.withRiskLocation != nil {
			_spec.Node.AddColumnOnce(sweep.FieldRiskLocationID)
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SweepQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(sweep.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = sweep.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SweepGroupBy is the group-by builder for Sweep entities.
type SweepGroupBy struct {
	selector
	build *SweepQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SweepGroupBy) Aggregate(fns ...AggregateFunc) *SweepGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SweepGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SweepQuery, *SweepGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SweepGroupBy) sqlScan(ctx context.Context, root *SweepQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SweepSelect is the builder for selecting fields of Sweep entities.
type SweepSelect struct {
	*SweepQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SweepSelect) Aggregate(fns ...AggregateFunc) *SweepSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SweepSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SweepQuery, *SweepSelect](ctx, ss.SweepQuery, ss, ss.inters, v)
}

func (ss *SweepSelect) sqlScan(ctx context.Context, root *SweepQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
