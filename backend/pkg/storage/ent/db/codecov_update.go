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
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/codecov"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
)

// CodeCovUpdate is the builder for updating CodeCov entities.
type CodeCovUpdate struct {
	config
	hooks    []Hook
	mutation *CodeCovMutation
}

// Where appends a list predicates to the CodeCovUpdate builder.
func (ccu *CodeCovUpdate) Where(ps ...predicate.CodeCov) *CodeCovUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetRepositoryName sets the "repository_name" field.
func (ccu *CodeCovUpdate) SetRepositoryName(s string) *CodeCovUpdate {
	ccu.mutation.SetRepositoryName(s)
	return ccu
}

// SetGitOrganization sets the "git_organization" field.
func (ccu *CodeCovUpdate) SetGitOrganization(s string) *CodeCovUpdate {
	ccu.mutation.SetGitOrganization(s)
	return ccu
}

// SetCoveragePercentage sets the "coverage_percentage" field.
func (ccu *CodeCovUpdate) SetCoveragePercentage(f float64) *CodeCovUpdate {
	ccu.mutation.ResetCoveragePercentage()
	ccu.mutation.SetCoveragePercentage(f)
	return ccu
}

// AddCoveragePercentage adds f to the "coverage_percentage" field.
func (ccu *CodeCovUpdate) AddCoveragePercentage(f float64) *CodeCovUpdate {
	ccu.mutation.AddCoveragePercentage(f)
	return ccu
}

// SetCodecovID sets the "codecov" edge to the Repository entity by ID.
func (ccu *CodeCovUpdate) SetCodecovID(id uuid.UUID) *CodeCovUpdate {
	ccu.mutation.SetCodecovID(id)
	return ccu
}

// SetNillableCodecovID sets the "codecov" edge to the Repository entity by ID if the given value is not nil.
func (ccu *CodeCovUpdate) SetNillableCodecovID(id *uuid.UUID) *CodeCovUpdate {
	if id != nil {
		ccu = ccu.SetCodecovID(*id)
	}
	return ccu
}

// SetCodecov sets the "codecov" edge to the Repository entity.
func (ccu *CodeCovUpdate) SetCodecov(r *Repository) *CodeCovUpdate {
	return ccu.SetCodecovID(r.ID)
}

// Mutation returns the CodeCovMutation object of the builder.
func (ccu *CodeCovUpdate) Mutation() *CodeCovMutation {
	return ccu.mutation
}

// ClearCodecov clears the "codecov" edge to the Repository entity.
func (ccu *CodeCovUpdate) ClearCodecov() *CodeCovUpdate {
	ccu.mutation.ClearCodecov()
	return ccu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *CodeCovUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ccu.hooks) == 0 {
		if err = ccu.check(); err != nil {
			return 0, err
		}
		affected, err = ccu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodeCovMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ccu.check(); err != nil {
				return 0, err
			}
			ccu.mutation = mutation
			affected, err = ccu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccu.hooks) - 1; i >= 0; i-- {
			if ccu.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = ccu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *CodeCovUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *CodeCovUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *CodeCovUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccu *CodeCovUpdate) check() error {
	if v, ok := ccu.mutation.RepositoryName(); ok {
		if err := codecov.RepositoryNameValidator(v); err != nil {
			return &ValidationError{Name: "repository_name", err: fmt.Errorf(`db: validator failed for field "CodeCov.repository_name": %w`, err)}
		}
	}
	return nil
}

func (ccu *CodeCovUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codecov.Table,
			Columns: codecov.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: codecov.FieldID,
			},
		},
	}
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.RepositoryName(); ok {
		_spec.SetField(codecov.FieldRepositoryName, field.TypeString, value)
	}
	if value, ok := ccu.mutation.GitOrganization(); ok {
		_spec.SetField(codecov.FieldGitOrganization, field.TypeString, value)
	}
	if value, ok := ccu.mutation.CoveragePercentage(); ok {
		_spec.SetField(codecov.FieldCoveragePercentage, field.TypeFloat64, value)
	}
	if value, ok := ccu.mutation.AddedCoveragePercentage(); ok {
		_spec.AddField(codecov.FieldCoveragePercentage, field.TypeFloat64, value)
	}
	if ccu.mutation.CodecovCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codecov.CodecovTable,
			Columns: []string{codecov.CodecovColumn},
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
	if nodes := ccu.mutation.CodecovIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codecov.CodecovTable,
			Columns: []string{codecov.CodecovColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codecov.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CodeCovUpdateOne is the builder for updating a single CodeCov entity.
type CodeCovUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CodeCovMutation
}

// SetRepositoryName sets the "repository_name" field.
func (ccuo *CodeCovUpdateOne) SetRepositoryName(s string) *CodeCovUpdateOne {
	ccuo.mutation.SetRepositoryName(s)
	return ccuo
}

// SetGitOrganization sets the "git_organization" field.
func (ccuo *CodeCovUpdateOne) SetGitOrganization(s string) *CodeCovUpdateOne {
	ccuo.mutation.SetGitOrganization(s)
	return ccuo
}

// SetCoveragePercentage sets the "coverage_percentage" field.
func (ccuo *CodeCovUpdateOne) SetCoveragePercentage(f float64) *CodeCovUpdateOne {
	ccuo.mutation.ResetCoveragePercentage()
	ccuo.mutation.SetCoveragePercentage(f)
	return ccuo
}

// AddCoveragePercentage adds f to the "coverage_percentage" field.
func (ccuo *CodeCovUpdateOne) AddCoveragePercentage(f float64) *CodeCovUpdateOne {
	ccuo.mutation.AddCoveragePercentage(f)
	return ccuo
}

// SetCodecovID sets the "codecov" edge to the Repository entity by ID.
func (ccuo *CodeCovUpdateOne) SetCodecovID(id uuid.UUID) *CodeCovUpdateOne {
	ccuo.mutation.SetCodecovID(id)
	return ccuo
}

// SetNillableCodecovID sets the "codecov" edge to the Repository entity by ID if the given value is not nil.
func (ccuo *CodeCovUpdateOne) SetNillableCodecovID(id *uuid.UUID) *CodeCovUpdateOne {
	if id != nil {
		ccuo = ccuo.SetCodecovID(*id)
	}
	return ccuo
}

// SetCodecov sets the "codecov" edge to the Repository entity.
func (ccuo *CodeCovUpdateOne) SetCodecov(r *Repository) *CodeCovUpdateOne {
	return ccuo.SetCodecovID(r.ID)
}

// Mutation returns the CodeCovMutation object of the builder.
func (ccuo *CodeCovUpdateOne) Mutation() *CodeCovMutation {
	return ccuo.mutation
}

// ClearCodecov clears the "codecov" edge to the Repository entity.
func (ccuo *CodeCovUpdateOne) ClearCodecov() *CodeCovUpdateOne {
	ccuo.mutation.ClearCodecov()
	return ccuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *CodeCovUpdateOne) Select(field string, fields ...string) *CodeCovUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated CodeCov entity.
func (ccuo *CodeCovUpdateOne) Save(ctx context.Context) (*CodeCov, error) {
	var (
		err  error
		node *CodeCov
	)
	if len(ccuo.hooks) == 0 {
		if err = ccuo.check(); err != nil {
			return nil, err
		}
		node, err = ccuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CodeCovMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ccuo.check(); err != nil {
				return nil, err
			}
			ccuo.mutation = mutation
			node, err = ccuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ccuo.hooks) - 1; i >= 0; i-- {
			if ccuo.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = ccuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ccuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CodeCov)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CodeCovMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *CodeCovUpdateOne) SaveX(ctx context.Context) *CodeCov {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *CodeCovUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *CodeCovUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccuo *CodeCovUpdateOne) check() error {
	if v, ok := ccuo.mutation.RepositoryName(); ok {
		if err := codecov.RepositoryNameValidator(v); err != nil {
			return &ValidationError{Name: "repository_name", err: fmt.Errorf(`db: validator failed for field "CodeCov.repository_name": %w`, err)}
		}
	}
	return nil
}

func (ccuo *CodeCovUpdateOne) sqlSave(ctx context.Context) (_node *CodeCov, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   codecov.Table,
			Columns: codecov.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: codecov.FieldID,
			},
		},
	}
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "CodeCov.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, codecov.FieldID)
		for _, f := range fields {
			if !codecov.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != codecov.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.RepositoryName(); ok {
		_spec.SetField(codecov.FieldRepositoryName, field.TypeString, value)
	}
	if value, ok := ccuo.mutation.GitOrganization(); ok {
		_spec.SetField(codecov.FieldGitOrganization, field.TypeString, value)
	}
	if value, ok := ccuo.mutation.CoveragePercentage(); ok {
		_spec.SetField(codecov.FieldCoveragePercentage, field.TypeFloat64, value)
	}
	if value, ok := ccuo.mutation.AddedCoveragePercentage(); ok {
		_spec.AddField(codecov.FieldCoveragePercentage, field.TypeFloat64, value)
	}
	if ccuo.mutation.CodecovCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codecov.CodecovTable,
			Columns: []string{codecov.CodecovColumn},
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
	if nodes := ccuo.mutation.CodecovIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codecov.CodecovTable,
			Columns: []string{codecov.CodecovColumn},
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
	_node = &CodeCov{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{codecov.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
