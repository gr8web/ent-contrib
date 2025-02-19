// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/explicitskippedmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExplicitSkippedMessageUpdate is the builder for updating ExplicitSkippedMessage entities.
type ExplicitSkippedMessageUpdate struct {
	config
	hooks    []Hook
	mutation *ExplicitSkippedMessageMutation
}

// Where appends a list predicates to the ExplicitSkippedMessageUpdate builder.
func (esmu *ExplicitSkippedMessageUpdate) Where(ps ...predicate.ExplicitSkippedMessage) *ExplicitSkippedMessageUpdate {
	esmu.mutation.Where(ps...)
	return esmu
}

// Mutation returns the ExplicitSkippedMessageMutation object of the builder.
func (esmu *ExplicitSkippedMessageUpdate) Mutation() *ExplicitSkippedMessageMutation {
	return esmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (esmu *ExplicitSkippedMessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ExplicitSkippedMessageMutation](ctx, esmu.sqlSave, esmu.mutation, esmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (esmu *ExplicitSkippedMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := esmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (esmu *ExplicitSkippedMessageUpdate) Exec(ctx context.Context) error {
	_, err := esmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esmu *ExplicitSkippedMessageUpdate) ExecX(ctx context.Context) {
	if err := esmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (esmu *ExplicitSkippedMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   explicitskippedmessage.Table,
			Columns: explicitskippedmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: explicitskippedmessage.FieldID,
			},
		},
	}
	if ps := esmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, esmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{explicitskippedmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	esmu.mutation.done = true
	return n, nil
}

// ExplicitSkippedMessageUpdateOne is the builder for updating a single ExplicitSkippedMessage entity.
type ExplicitSkippedMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExplicitSkippedMessageMutation
}

// Mutation returns the ExplicitSkippedMessageMutation object of the builder.
func (esmuo *ExplicitSkippedMessageUpdateOne) Mutation() *ExplicitSkippedMessageMutation {
	return esmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (esmuo *ExplicitSkippedMessageUpdateOne) Select(field string, fields ...string) *ExplicitSkippedMessageUpdateOne {
	esmuo.fields = append([]string{field}, fields...)
	return esmuo
}

// Save executes the query and returns the updated ExplicitSkippedMessage entity.
func (esmuo *ExplicitSkippedMessageUpdateOne) Save(ctx context.Context) (*ExplicitSkippedMessage, error) {
	return withHooks[*ExplicitSkippedMessage, ExplicitSkippedMessageMutation](ctx, esmuo.sqlSave, esmuo.mutation, esmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (esmuo *ExplicitSkippedMessageUpdateOne) SaveX(ctx context.Context) *ExplicitSkippedMessage {
	node, err := esmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (esmuo *ExplicitSkippedMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := esmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esmuo *ExplicitSkippedMessageUpdateOne) ExecX(ctx context.Context) {
	if err := esmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (esmuo *ExplicitSkippedMessageUpdateOne) sqlSave(ctx context.Context) (_node *ExplicitSkippedMessage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   explicitskippedmessage.Table,
			Columns: explicitskippedmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: explicitskippedmessage.FieldID,
			},
		},
	}
	id, ok := esmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExplicitSkippedMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := esmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, explicitskippedmessage.FieldID)
		for _, f := range fields {
			if !explicitskippedmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != explicitskippedmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := esmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &ExplicitSkippedMessage{config: esmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, esmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{explicitskippedmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	esmuo.mutation.done = true
	return _node, nil
}
