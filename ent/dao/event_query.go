// Code generated by ent, DO NOT EDIT.

package dao

import (
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/device"
	"aisecurity/ent/dao/event"
	"aisecurity/ent/dao/eventlog"
	"aisecurity/ent/dao/fixing"
	"aisecurity/ent/dao/predicate"
	"aisecurity/ent/dao/video"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventQuery is the builder for querying Event entities.
type EventQuery struct {
	config
	ctx          *QueryContext
	order        []event.OrderOption
	inters       []Interceptor
	predicates   []predicate.Event
	withCreator  *AdminQuery
	withUpdater  *AdminQuery
	withVideo    *VideoQuery
	withDevice   *DeviceQuery
	withFixing   *FixingQuery
	withEventLog *EventLogQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EventQuery builder.
func (eq *EventQuery) Where(ps ...predicate.Event) *EventQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EventQuery) Limit(limit int) *EventQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EventQuery) Offset(offset int) *EventQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EventQuery) Unique(unique bool) *EventQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EventQuery) Order(o ...event.OrderOption) *EventQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryCreator chains the current query on the "creator" edge.
func (eq *EventQuery) QueryCreator() *AdminQuery {
	query := (&AdminClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, selector),
			sqlgraph.To(admin.Table, admin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, event.CreatorTable, event.CreatorColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUpdater chains the current query on the "updater" edge.
func (eq *EventQuery) QueryUpdater() *AdminQuery {
	query := (&AdminClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, selector),
			sqlgraph.To(admin.Table, admin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, event.UpdaterTable, event.UpdaterColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVideo chains the current query on the "video" edge.
func (eq *EventQuery) QueryVideo() *VideoQuery {
	query := (&VideoClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, selector),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, event.VideoTable, event.VideoColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDevice chains the current query on the "device" edge.
func (eq *EventQuery) QueryDevice() *DeviceQuery {
	query := (&DeviceClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, event.DeviceTable, event.DeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFixing chains the current query on the "fixing" edge.
func (eq *EventQuery) QueryFixing() *FixingQuery {
	query := (&FixingClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, selector),
			sqlgraph.To(fixing.Table, fixing.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, event.FixingTable, event.FixingColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEventLog chains the current query on the "event_log" edge.
func (eq *EventQuery) QueryEventLog() *EventLogQuery {
	query := (&EventLogClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, selector),
			sqlgraph.To(eventlog.Table, eventlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, event.EventLogTable, event.EventLogColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Event entity from the query.
// Returns a *NotFoundError when no Event was found.
func (eq *EventQuery) First(ctx context.Context) (*Event, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{event.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EventQuery) FirstX(ctx context.Context) *Event {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Event ID from the query.
// Returns a *NotFoundError when no Event ID was found.
func (eq *EventQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{event.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EventQuery) FirstIDX(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Event entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Event entity is found.
// Returns a *NotFoundError when no Event entities are found.
func (eq *EventQuery) Only(ctx context.Context) (*Event, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{event.Label}
	default:
		return nil, &NotSingularError{event.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EventQuery) OnlyX(ctx context.Context) *Event {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Event ID in the query.
// Returns a *NotSingularError when more than one Event ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EventQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{event.Label}
	default:
		err = &NotSingularError{event.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EventQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Events.
func (eq *EventQuery) All(ctx context.Context) ([]*Event, error) {
	ctx = setContextOp(ctx, eq.ctx, "All")
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Event, *EventQuery]()
	return withInterceptors[[]*Event](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EventQuery) AllX(ctx context.Context) []*Event {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Event IDs.
func (eq *EventQuery) IDs(ctx context.Context) (ids []int, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, "IDs")
	if err = eq.Select(event.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EventQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EventQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, "Count")
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EventQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EventQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EventQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, "Exist")
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("dao: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EventQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EventQuery) Clone() *EventQuery {
	if eq == nil {
		return nil
	}
	return &EventQuery{
		config:       eq.config,
		ctx:          eq.ctx.Clone(),
		order:        append([]event.OrderOption{}, eq.order...),
		inters:       append([]Interceptor{}, eq.inters...),
		predicates:   append([]predicate.Event{}, eq.predicates...),
		withCreator:  eq.withCreator.Clone(),
		withUpdater:  eq.withUpdater.Clone(),
		withVideo:    eq.withVideo.Clone(),
		withDevice:   eq.withDevice.Clone(),
		withFixing:   eq.withFixing.Clone(),
		withEventLog: eq.withEventLog.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithCreator tells the query-builder to eager-load the nodes that are connected to
// the "creator" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EventQuery) WithCreator(opts ...func(*AdminQuery)) *EventQuery {
	query := (&AdminClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withCreator = query
	return eq
}

// WithUpdater tells the query-builder to eager-load the nodes that are connected to
// the "updater" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EventQuery) WithUpdater(opts ...func(*AdminQuery)) *EventQuery {
	query := (&AdminClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withUpdater = query
	return eq
}

// WithVideo tells the query-builder to eager-load the nodes that are connected to
// the "video" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EventQuery) WithVideo(opts ...func(*VideoQuery)) *EventQuery {
	query := (&VideoClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withVideo = query
	return eq
}

// WithDevice tells the query-builder to eager-load the nodes that are connected to
// the "device" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EventQuery) WithDevice(opts ...func(*DeviceQuery)) *EventQuery {
	query := (&DeviceClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withDevice = query
	return eq
}

// WithFixing tells the query-builder to eager-load the nodes that are connected to
// the "fixing" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EventQuery) WithFixing(opts ...func(*FixingQuery)) *EventQuery {
	query := (&FixingClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withFixing = query
	return eq
}

// WithEventLog tells the query-builder to eager-load the nodes that are connected to
// the "event_log" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EventQuery) WithEventLog(opts ...func(*EventLogQuery)) *EventQuery {
	query := (&EventLogClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withEventLog = query
	return eq
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
//	client.Event.Query().
//		GroupBy(event.FieldCreateTime).
//		Aggregate(dao.Count()).
//		Scan(ctx, &v)
func (eq *EventQuery) GroupBy(field string, fields ...string) *EventGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EventGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = event.Label
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
//	client.Event.Query().
//		Select(event.FieldCreateTime).
//		Scan(ctx, &v)
func (eq *EventQuery) Select(fields ...string) *EventSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EventSelect{EventQuery: eq}
	sbuild.label = event.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EventSelect configured with the given aggregations.
func (eq *EventQuery) Aggregate(fns ...AggregateFunc) *EventSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EventQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("dao: uninitialized interceptor (forgotten import dao/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !event.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("dao: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Event, error) {
	var (
		nodes       = []*Event{}
		_spec       = eq.querySpec()
		loadedTypes = [6]bool{
			eq.withCreator != nil,
			eq.withUpdater != nil,
			eq.withVideo != nil,
			eq.withDevice != nil,
			eq.withFixing != nil,
			eq.withEventLog != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Event).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Event{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withCreator; query != nil {
		if err := eq.loadCreator(ctx, query, nodes, nil,
			func(n *Event, e *Admin) { n.Edges.Creator = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withUpdater; query != nil {
		if err := eq.loadUpdater(ctx, query, nodes, nil,
			func(n *Event, e *Admin) { n.Edges.Updater = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withVideo; query != nil {
		if err := eq.loadVideo(ctx, query, nodes, nil,
			func(n *Event, e *Video) { n.Edges.Video = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withDevice; query != nil {
		if err := eq.loadDevice(ctx, query, nodes, nil,
			func(n *Event, e *Device) { n.Edges.Device = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withFixing; query != nil {
		if err := eq.loadFixing(ctx, query, nodes, nil,
			func(n *Event, e *Fixing) { n.Edges.Fixing = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withEventLog; query != nil {
		if err := eq.loadEventLog(ctx, query, nodes,
			func(n *Event) { n.Edges.EventLog = []*EventLog{} },
			func(n *Event, e *EventLog) { n.Edges.EventLog = append(n.Edges.EventLog, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EventQuery) loadCreator(ctx context.Context, query *AdminQuery, nodes []*Event, init func(*Event), assign func(*Event, *Admin)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Event)
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
func (eq *EventQuery) loadUpdater(ctx context.Context, query *AdminQuery, nodes []*Event, init func(*Event), assign func(*Event, *Admin)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Event)
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
func (eq *EventQuery) loadVideo(ctx context.Context, query *VideoQuery, nodes []*Event, init func(*Event), assign func(*Event, *Video)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Event)
	for i := range nodes {
		fk := nodes[i].VideoID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(video.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "video_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EventQuery) loadDevice(ctx context.Context, query *DeviceQuery, nodes []*Event, init func(*Event), assign func(*Event, *Device)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Event)
	for i := range nodes {
		fk := nodes[i].DeviceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(device.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EventQuery) loadFixing(ctx context.Context, query *FixingQuery, nodes []*Event, init func(*Event), assign func(*Event, *Fixing)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Event)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(fixing.FieldEventID)
	}
	query.Where(predicate.Fixing(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(event.FixingColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.EventID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "event_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eq *EventQuery) loadEventLog(ctx context.Context, query *EventLogQuery, nodes []*Event, init func(*Event), assign func(*Event, *EventLog)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Event)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(eventlog.FieldEventID)
	}
	query.Where(predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(event.EventLogColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.EventID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "event_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (eq *EventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(event.Table, event.Columns, sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for i := range fields {
			if fields[i] != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if eq.withCreator != nil {
			_spec.Node.AddColumnOnce(event.FieldCreatorID)
		}
		if eq.withUpdater != nil {
			_spec.Node.AddColumnOnce(event.FieldUpdaterID)
		}
		if eq.withVideo != nil {
			_spec.Node.AddColumnOnce(event.FieldVideoID)
		}
		if eq.withDevice != nil {
			_spec.Node.AddColumnOnce(event.FieldDeviceID)
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(event.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = event.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EventGroupBy is the group-by builder for Event entities.
type EventGroupBy struct {
	selector
	build *EventQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EventGroupBy) Aggregate(fns ...AggregateFunc) *EventGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EventGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, "GroupBy")
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EventQuery, *EventGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EventGroupBy) sqlScan(ctx context.Context, root *EventQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EventSelect is the builder for selecting fields of Event entities.
type EventSelect struct {
	*EventQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EventSelect) Aggregate(fns ...AggregateFunc) *EventSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EventSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, "Select")
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EventQuery, *EventSelect](ctx, es.EventQuery, es, es.inters, v)
}

func (es *EventSelect) sqlScan(ctx context.Context, root *EventQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}