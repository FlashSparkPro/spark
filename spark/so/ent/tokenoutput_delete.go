// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lightsparkdev/spark/so/ent/predicate"
	"github.com/lightsparkdev/spark/so/ent/tokenoutput"
)

// TokenOutputDelete is the builder for deleting a TokenOutput entity.
type TokenOutputDelete struct {
	config
	hooks    []Hook
	mutation *TokenOutputMutation
}

// Where appends a list predicates to the TokenOutputDelete builder.
func (tod *TokenOutputDelete) Where(ps ...predicate.TokenOutput) *TokenOutputDelete {
	tod.mutation.Where(ps...)
	return tod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tod *TokenOutputDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tod.sqlExec, tod.mutation, tod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tod *TokenOutputDelete) ExecX(ctx context.Context) int {
	n, err := tod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tod *TokenOutputDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tokenoutput.Table, sqlgraph.NewFieldSpec(tokenoutput.FieldID, field.TypeUUID))
	if ps := tod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tod.mutation.done = true
	return affected, err
}

// TokenOutputDeleteOne is the builder for deleting a single TokenOutput entity.
type TokenOutputDeleteOne struct {
	tod *TokenOutputDelete
}

// Where appends a list predicates to the TokenOutputDelete builder.
func (todo *TokenOutputDeleteOne) Where(ps ...predicate.TokenOutput) *TokenOutputDeleteOne {
	todo.tod.mutation.Where(ps...)
	return todo
}

// Exec executes the deletion query.
func (todo *TokenOutputDeleteOne) Exec(ctx context.Context) error {
	n, err := todo.tod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tokenoutput.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (todo *TokenOutputDeleteOne) ExecX(ctx context.Context) {
	if err := todo.Exec(ctx); err != nil {
		panic(err)
	}
}
