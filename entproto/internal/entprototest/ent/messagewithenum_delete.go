// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithenum"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithEnumDelete is the builder for deleting a MessageWithEnum entity.
type MessageWithEnumDelete struct {
	config
	hooks    []Hook
	mutation *MessageWithEnumMutation
}

// Where appends a list predicates to the MessageWithEnumDelete builder.
func (mwed *MessageWithEnumDelete) Where(ps ...predicate.MessageWithEnum) *MessageWithEnumDelete {
	mwed.mutation.Where(ps...)
	return mwed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mwed *MessageWithEnumDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, MessageWithEnumMutation](ctx, mwed.sqlExec, mwed.mutation, mwed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mwed *MessageWithEnumDelete) ExecX(ctx context.Context) int {
	n, err := mwed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mwed *MessageWithEnumDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: messagewithenum.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithenum.FieldID,
			},
		},
	}
	if ps := mwed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mwed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mwed.mutation.done = true
	return affected, err
}

// MessageWithEnumDeleteOne is the builder for deleting a single MessageWithEnum entity.
type MessageWithEnumDeleteOne struct {
	mwed *MessageWithEnumDelete
}

// Exec executes the deletion query.
func (mwedo *MessageWithEnumDeleteOne) Exec(ctx context.Context) error {
	n, err := mwedo.mwed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{messagewithenum.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mwedo *MessageWithEnumDeleteOne) ExecX(ctx context.Context) {
	mwedo.mwed.ExecX(ctx)
}
