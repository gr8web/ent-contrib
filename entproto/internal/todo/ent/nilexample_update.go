// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entproto/internal/todo/ent/nilexample"
	"entgo.io/contrib/entproto/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NilExampleUpdate is the builder for updating NilExample entities.
type NilExampleUpdate struct {
	config
	hooks    []Hook
	mutation *NilExampleMutation
}

// Where appends a list predicates to the NilExampleUpdate builder.
func (neu *NilExampleUpdate) Where(ps ...predicate.NilExample) *NilExampleUpdate {
	neu.mutation.Where(ps...)
	return neu
}

// SetStrNil sets the "str_nil" field.
func (neu *NilExampleUpdate) SetStrNil(s string) *NilExampleUpdate {
	neu.mutation.SetStrNil(s)
	return neu
}

// SetNillableStrNil sets the "str_nil" field if the given value is not nil.
func (neu *NilExampleUpdate) SetNillableStrNil(s *string) *NilExampleUpdate {
	if s != nil {
		neu.SetStrNil(*s)
	}
	return neu
}

// ClearStrNil clears the value of the "str_nil" field.
func (neu *NilExampleUpdate) ClearStrNil() *NilExampleUpdate {
	neu.mutation.ClearStrNil()
	return neu
}

// SetTimeNil sets the "time_nil" field.
func (neu *NilExampleUpdate) SetTimeNil(t time.Time) *NilExampleUpdate {
	neu.mutation.SetTimeNil(t)
	return neu
}

// SetNillableTimeNil sets the "time_nil" field if the given value is not nil.
func (neu *NilExampleUpdate) SetNillableTimeNil(t *time.Time) *NilExampleUpdate {
	if t != nil {
		neu.SetTimeNil(*t)
	}
	return neu
}

// ClearTimeNil clears the value of the "time_nil" field.
func (neu *NilExampleUpdate) ClearTimeNil() *NilExampleUpdate {
	neu.mutation.ClearTimeNil()
	return neu
}

// Mutation returns the NilExampleMutation object of the builder.
func (neu *NilExampleUpdate) Mutation() *NilExampleMutation {
	return neu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (neu *NilExampleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, NilExampleMutation](ctx, neu.sqlSave, neu.mutation, neu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (neu *NilExampleUpdate) SaveX(ctx context.Context) int {
	affected, err := neu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (neu *NilExampleUpdate) Exec(ctx context.Context) error {
	_, err := neu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (neu *NilExampleUpdate) ExecX(ctx context.Context) {
	if err := neu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (neu *NilExampleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nilexample.Table,
			Columns: nilexample.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nilexample.FieldID,
			},
		},
	}
	if ps := neu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := neu.mutation.StrNil(); ok {
		_spec.SetField(nilexample.FieldStrNil, field.TypeString, value)
	}
	if neu.mutation.StrNilCleared() {
		_spec.ClearField(nilexample.FieldStrNil, field.TypeString)
	}
	if value, ok := neu.mutation.TimeNil(); ok {
		_spec.SetField(nilexample.FieldTimeNil, field.TypeTime, value)
	}
	if neu.mutation.TimeNilCleared() {
		_spec.ClearField(nilexample.FieldTimeNil, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, neu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nilexample.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	neu.mutation.done = true
	return n, nil
}

// NilExampleUpdateOne is the builder for updating a single NilExample entity.
type NilExampleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NilExampleMutation
}

// SetStrNil sets the "str_nil" field.
func (neuo *NilExampleUpdateOne) SetStrNil(s string) *NilExampleUpdateOne {
	neuo.mutation.SetStrNil(s)
	return neuo
}

// SetNillableStrNil sets the "str_nil" field if the given value is not nil.
func (neuo *NilExampleUpdateOne) SetNillableStrNil(s *string) *NilExampleUpdateOne {
	if s != nil {
		neuo.SetStrNil(*s)
	}
	return neuo
}

// ClearStrNil clears the value of the "str_nil" field.
func (neuo *NilExampleUpdateOne) ClearStrNil() *NilExampleUpdateOne {
	neuo.mutation.ClearStrNil()
	return neuo
}

// SetTimeNil sets the "time_nil" field.
func (neuo *NilExampleUpdateOne) SetTimeNil(t time.Time) *NilExampleUpdateOne {
	neuo.mutation.SetTimeNil(t)
	return neuo
}

// SetNillableTimeNil sets the "time_nil" field if the given value is not nil.
func (neuo *NilExampleUpdateOne) SetNillableTimeNil(t *time.Time) *NilExampleUpdateOne {
	if t != nil {
		neuo.SetTimeNil(*t)
	}
	return neuo
}

// ClearTimeNil clears the value of the "time_nil" field.
func (neuo *NilExampleUpdateOne) ClearTimeNil() *NilExampleUpdateOne {
	neuo.mutation.ClearTimeNil()
	return neuo
}

// Mutation returns the NilExampleMutation object of the builder.
func (neuo *NilExampleUpdateOne) Mutation() *NilExampleMutation {
	return neuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (neuo *NilExampleUpdateOne) Select(field string, fields ...string) *NilExampleUpdateOne {
	neuo.fields = append([]string{field}, fields...)
	return neuo
}

// Save executes the query and returns the updated NilExample entity.
func (neuo *NilExampleUpdateOne) Save(ctx context.Context) (*NilExample, error) {
	return withHooks[*NilExample, NilExampleMutation](ctx, neuo.sqlSave, neuo.mutation, neuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (neuo *NilExampleUpdateOne) SaveX(ctx context.Context) *NilExample {
	node, err := neuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (neuo *NilExampleUpdateOne) Exec(ctx context.Context) error {
	_, err := neuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (neuo *NilExampleUpdateOne) ExecX(ctx context.Context) {
	if err := neuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (neuo *NilExampleUpdateOne) sqlSave(ctx context.Context) (_node *NilExample, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nilexample.Table,
			Columns: nilexample.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nilexample.FieldID,
			},
		},
	}
	id, ok := neuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NilExample.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := neuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nilexample.FieldID)
		for _, f := range fields {
			if !nilexample.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nilexample.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := neuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := neuo.mutation.StrNil(); ok {
		_spec.SetField(nilexample.FieldStrNil, field.TypeString, value)
	}
	if neuo.mutation.StrNilCleared() {
		_spec.ClearField(nilexample.FieldStrNil, field.TypeString)
	}
	if value, ok := neuo.mutation.TimeNil(); ok {
		_spec.SetField(nilexample.FieldTimeNil, field.TypeTime, value)
	}
	if neuo.mutation.TimeNilCleared() {
		_spec.ClearField(nilexample.FieldTimeNil, field.TypeTime)
	}
	_node = &NilExample{config: neuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, neuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nilexample.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	neuo.mutation.done = true
	return _node, nil
}
