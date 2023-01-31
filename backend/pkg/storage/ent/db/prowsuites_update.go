// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/prowsuites"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
)

// ProwSuitesUpdate is the builder for updating ProwSuites entities.
type ProwSuitesUpdate struct {
	config
	hooks    []Hook
	mutation *ProwSuitesMutation
}

// Where appends a list predicates to the ProwSuitesUpdate builder.
func (psu *ProwSuitesUpdate) Where(ps ...predicate.ProwSuites) *ProwSuitesUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetJobID sets the "job_id" field.
func (psu *ProwSuitesUpdate) SetJobID(s string) *ProwSuitesUpdate {
	psu.mutation.SetJobID(s)
	return psu
}

// SetName sets the "name" field.
func (psu *ProwSuitesUpdate) SetName(s string) *ProwSuitesUpdate {
	psu.mutation.SetName(s)
	return psu
}

// SetStatus sets the "status" field.
func (psu *ProwSuitesUpdate) SetStatus(s string) *ProwSuitesUpdate {
	psu.mutation.SetStatus(s)
	return psu
}

// SetTime sets the "time" field.
func (psu *ProwSuitesUpdate) SetTime(f float64) *ProwSuitesUpdate {
	psu.mutation.ResetTime()
	psu.mutation.SetTime(f)
	return psu
}

// AddTime adds f to the "time" field.
func (psu *ProwSuitesUpdate) AddTime(f float64) *ProwSuitesUpdate {
	psu.mutation.AddTime(f)
	return psu
}

// SetProwSuitesID sets the "prow_suites" edge to the Repository entity by ID.
func (psu *ProwSuitesUpdate) SetProwSuitesID(id uuid.UUID) *ProwSuitesUpdate {
	psu.mutation.SetProwSuitesID(id)
	return psu
}

// SetNillableProwSuitesID sets the "prow_suites" edge to the Repository entity by ID if the given value is not nil.
func (psu *ProwSuitesUpdate) SetNillableProwSuitesID(id *uuid.UUID) *ProwSuitesUpdate {
	if id != nil {
		psu = psu.SetProwSuitesID(*id)
	}
	return psu
}

// SetProwSuites sets the "prow_suites" edge to the Repository entity.
func (psu *ProwSuitesUpdate) SetProwSuites(r *Repository) *ProwSuitesUpdate {
	return psu.SetProwSuitesID(r.ID)
}

// Mutation returns the ProwSuitesMutation object of the builder.
func (psu *ProwSuitesUpdate) Mutation() *ProwSuitesMutation {
	return psu.mutation
}

// ClearProwSuites clears the "prow_suites" edge to the Repository entity.
func (psu *ProwSuitesUpdate) ClearProwSuites() *ProwSuitesUpdate {
	psu.mutation.ClearProwSuites()
	return psu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *ProwSuitesUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(psu.hooks) == 0 {
		affected, err = psu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProwSuitesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psu.mutation = mutation
			affected, err = psu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(psu.hooks) - 1; i >= 0; i-- {
			if psu.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = psu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (psu *ProwSuitesUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *ProwSuitesUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *ProwSuitesUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (psu *ProwSuitesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prowsuites.Table,
			Columns: prowsuites.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prowsuites.FieldID,
			},
		},
	}
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psu.mutation.JobID(); ok {
		_spec.SetField(prowsuites.FieldJobID, field.TypeString, value)
	}
	if value, ok := psu.mutation.Name(); ok {
		_spec.SetField(prowsuites.FieldName, field.TypeString, value)
	}
	if value, ok := psu.mutation.Status(); ok {
		_spec.SetField(prowsuites.FieldStatus, field.TypeString, value)
	}
	if value, ok := psu.mutation.Time(); ok {
		_spec.SetField(prowsuites.FieldTime, field.TypeFloat64, value)
	}
	if value, ok := psu.mutation.AddedTime(); ok {
		_spec.AddField(prowsuites.FieldTime, field.TypeFloat64, value)
	}
	if psu.mutation.ProwSuitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prowsuites.ProwSuitesTable,
			Columns: []string{prowsuites.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.ProwSuitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prowsuites.ProwSuitesTable,
			Columns: []string{prowsuites.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prowsuites.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ProwSuitesUpdateOne is the builder for updating a single ProwSuites entity.
type ProwSuitesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProwSuitesMutation
}

// SetJobID sets the "job_id" field.
func (psuo *ProwSuitesUpdateOne) SetJobID(s string) *ProwSuitesUpdateOne {
	psuo.mutation.SetJobID(s)
	return psuo
}

// SetName sets the "name" field.
func (psuo *ProwSuitesUpdateOne) SetName(s string) *ProwSuitesUpdateOne {
	psuo.mutation.SetName(s)
	return psuo
}

// SetStatus sets the "status" field.
func (psuo *ProwSuitesUpdateOne) SetStatus(s string) *ProwSuitesUpdateOne {
	psuo.mutation.SetStatus(s)
	return psuo
}

// SetTime sets the "time" field.
func (psuo *ProwSuitesUpdateOne) SetTime(f float64) *ProwSuitesUpdateOne {
	psuo.mutation.ResetTime()
	psuo.mutation.SetTime(f)
	return psuo
}

// AddTime adds f to the "time" field.
func (psuo *ProwSuitesUpdateOne) AddTime(f float64) *ProwSuitesUpdateOne {
	psuo.mutation.AddTime(f)
	return psuo
}

// SetProwSuitesID sets the "prow_suites" edge to the Repository entity by ID.
func (psuo *ProwSuitesUpdateOne) SetProwSuitesID(id uuid.UUID) *ProwSuitesUpdateOne {
	psuo.mutation.SetProwSuitesID(id)
	return psuo
}

// SetNillableProwSuitesID sets the "prow_suites" edge to the Repository entity by ID if the given value is not nil.
func (psuo *ProwSuitesUpdateOne) SetNillableProwSuitesID(id *uuid.UUID) *ProwSuitesUpdateOne {
	if id != nil {
		psuo = psuo.SetProwSuitesID(*id)
	}
	return psuo
}

// SetProwSuites sets the "prow_suites" edge to the Repository entity.
func (psuo *ProwSuitesUpdateOne) SetProwSuites(r *Repository) *ProwSuitesUpdateOne {
	return psuo.SetProwSuitesID(r.ID)
}

// Mutation returns the ProwSuitesMutation object of the builder.
func (psuo *ProwSuitesUpdateOne) Mutation() *ProwSuitesMutation {
	return psuo.mutation
}

// ClearProwSuites clears the "prow_suites" edge to the Repository entity.
func (psuo *ProwSuitesUpdateOne) ClearProwSuites() *ProwSuitesUpdateOne {
	psuo.mutation.ClearProwSuites()
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *ProwSuitesUpdateOne) Select(field string, fields ...string) *ProwSuitesUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated ProwSuites entity.
func (psuo *ProwSuitesUpdateOne) Save(ctx context.Context) (*ProwSuites, error) {
	var (
		err  error
		node *ProwSuites
	)
	if len(psuo.hooks) == 0 {
		node, err = psuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProwSuitesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psuo.mutation = mutation
			node, err = psuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(psuo.hooks) - 1; i >= 0; i-- {
			if psuo.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = psuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, psuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ProwSuites)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProwSuitesMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *ProwSuitesUpdateOne) SaveX(ctx context.Context) *ProwSuites {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *ProwSuitesUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *ProwSuitesUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (psuo *ProwSuitesUpdateOne) sqlSave(ctx context.Context) (_node *ProwSuites, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prowsuites.Table,
			Columns: prowsuites.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prowsuites.FieldID,
			},
		},
	}
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "ProwSuites.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prowsuites.FieldID)
		for _, f := range fields {
			if !prowsuites.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != prowsuites.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psuo.mutation.JobID(); ok {
		_spec.SetField(prowsuites.FieldJobID, field.TypeString, value)
	}
	if value, ok := psuo.mutation.Name(); ok {
		_spec.SetField(prowsuites.FieldName, field.TypeString, value)
	}
	if value, ok := psuo.mutation.Status(); ok {
		_spec.SetField(prowsuites.FieldStatus, field.TypeString, value)
	}
	if value, ok := psuo.mutation.Time(); ok {
		_spec.SetField(prowsuites.FieldTime, field.TypeFloat64, value)
	}
	if value, ok := psuo.mutation.AddedTime(); ok {
		_spec.AddField(prowsuites.FieldTime, field.TypeFloat64, value)
	}
	if psuo.mutation.ProwSuitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prowsuites.ProwSuitesTable,
			Columns: []string{prowsuites.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.ProwSuitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prowsuites.ProwSuitesTable,
			Columns: []string{prowsuites.ProwSuitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProwSuites{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prowsuites.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
