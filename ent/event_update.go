// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gmarcha/goswagger-ent-app/ent/event"
	"github.com/gmarcha/goswagger-ent-app/ent/predicate"
	"github.com/gmarcha/goswagger-ent-app/ent/user"
	"github.com/google/uuid"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EventUpdate) SetName(s string) *EventUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetDescription sets the "description" field.
func (eu *EventUpdate) SetDescription(s string) *EventUpdate {
	eu.mutation.SetDescription(s)
	return eu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (eu *EventUpdate) SetNillableDescription(s *string) *EventUpdate {
	if s != nil {
		eu.SetDescription(*s)
	}
	return eu
}

// SetDateUpdate sets the "date_update" field.
func (eu *EventUpdate) SetDateUpdate(t time.Time) *EventUpdate {
	eu.mutation.SetDateUpdate(t)
	return eu
}

// SetDateStart sets the "date_start" field.
func (eu *EventUpdate) SetDateStart(t time.Time) *EventUpdate {
	eu.mutation.SetDateStart(t)
	return eu
}

// SetDateEnd sets the "date_end" field.
func (eu *EventUpdate) SetDateEnd(t time.Time) *EventUpdate {
	eu.mutation.SetDateEnd(t)
	return eu
}

// SetNbTutors sets the "nb_tutors" field.
func (eu *EventUpdate) SetNbTutors(i int) *EventUpdate {
	eu.mutation.ResetNbTutors()
	eu.mutation.SetNbTutors(i)
	return eu
}

// AddNbTutors adds i to the "nb_tutors" field.
func (eu *EventUpdate) AddNbTutors(i int) *EventUpdate {
	eu.mutation.AddNbTutors(i)
	return eu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (eu *EventUpdate) AddUserIDs(ids ...uuid.UUID) *EventUpdate {
	eu.mutation.AddUserIDs(ids...)
	return eu
}

// AddUsers adds the "users" edges to the User entity.
func (eu *EventUpdate) AddUsers(u ...*User) *EventUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return eu.AddUserIDs(ids...)
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (eu *EventUpdate) ClearUsers() *EventUpdate {
	eu.mutation.ClearUsers()
	return eu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (eu *EventUpdate) RemoveUserIDs(ids ...uuid.UUID) *EventUpdate {
	eu.mutation.RemoveUserIDs(ids...)
	return eu
}

// RemoveUsers removes "users" edges to User entities.
func (eu *EventUpdate) RemoveUsers(u ...*User) *EventUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return eu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	eu.defaults()
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EventUpdate) defaults() {
	if _, ok := eu.mutation.DateUpdate(); !ok {
		v := event.UpdateDefaultDateUpdate()
		eu.mutation.SetDateUpdate(v)
	}
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldName,
		})
	}
	if value, ok := eu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldDescription,
		})
	}
	if value, ok := eu.mutation.DateUpdate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldDateUpdate,
		})
	}
	if value, ok := eu.mutation.DateStart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldDateStart,
		})
	}
	if value, ok := eu.mutation.DateEnd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldDateEnd,
		})
	}
	if value, ok := eu.mutation.NbTutors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldNbTutors,
		})
	}
	if value, ok := eu.mutation.AddedNbTutors(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldNbTutors,
		})
	}
	if eu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !eu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventMutation
}

// SetName sets the "name" field.
func (euo *EventUpdateOne) SetName(s string) *EventUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetDescription sets the "description" field.
func (euo *EventUpdateOne) SetDescription(s string) *EventUpdateOne {
	euo.mutation.SetDescription(s)
	return euo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableDescription(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetDescription(*s)
	}
	return euo
}

// SetDateUpdate sets the "date_update" field.
func (euo *EventUpdateOne) SetDateUpdate(t time.Time) *EventUpdateOne {
	euo.mutation.SetDateUpdate(t)
	return euo
}

// SetDateStart sets the "date_start" field.
func (euo *EventUpdateOne) SetDateStart(t time.Time) *EventUpdateOne {
	euo.mutation.SetDateStart(t)
	return euo
}

// SetDateEnd sets the "date_end" field.
func (euo *EventUpdateOne) SetDateEnd(t time.Time) *EventUpdateOne {
	euo.mutation.SetDateEnd(t)
	return euo
}

// SetNbTutors sets the "nb_tutors" field.
func (euo *EventUpdateOne) SetNbTutors(i int) *EventUpdateOne {
	euo.mutation.ResetNbTutors()
	euo.mutation.SetNbTutors(i)
	return euo
}

// AddNbTutors adds i to the "nb_tutors" field.
func (euo *EventUpdateOne) AddNbTutors(i int) *EventUpdateOne {
	euo.mutation.AddNbTutors(i)
	return euo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (euo *EventUpdateOne) AddUserIDs(ids ...uuid.UUID) *EventUpdateOne {
	euo.mutation.AddUserIDs(ids...)
	return euo
}

// AddUsers adds the "users" edges to the User entity.
func (euo *EventUpdateOne) AddUsers(u ...*User) *EventUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return euo.AddUserIDs(ids...)
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (euo *EventUpdateOne) ClearUsers() *EventUpdateOne {
	euo.mutation.ClearUsers()
	return euo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (euo *EventUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *EventUpdateOne {
	euo.mutation.RemoveUserIDs(ids...)
	return euo
}

// RemoveUsers removes "users" edges to User entities.
func (euo *EventUpdateOne) RemoveUsers(u ...*User) *EventUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return euo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	euo.defaults()
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EventUpdateOne) defaults() {
	if _, ok := euo.mutation.DateUpdate(); !ok {
		v := event.UpdateDefaultDateUpdate()
		euo.mutation.SetDateUpdate(v)
	}
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Event.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldName,
		})
	}
	if value, ok := euo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldDescription,
		})
	}
	if value, ok := euo.mutation.DateUpdate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldDateUpdate,
		})
	}
	if value, ok := euo.mutation.DateStart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldDateStart,
		})
	}
	if value, ok := euo.mutation.DateEnd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldDateEnd,
		})
	}
	if value, ok := euo.mutation.NbTutors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldNbTutors,
		})
	}
	if value, ok := euo.mutation.AddedNbTutors(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: event.FieldNbTutors,
		})
	}
	if euo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !euo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
