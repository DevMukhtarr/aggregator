// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/paycrest/aggregator/ent/lockorderfulfillment"
	"github.com/paycrest/aggregator/ent/lockpaymentorder"
	"github.com/paycrest/aggregator/ent/predicate"
)

// LockOrderFulfillmentUpdate is the builder for updating LockOrderFulfillment entities.
type LockOrderFulfillmentUpdate struct {
	config
	hooks    []Hook
	mutation *LockOrderFulfillmentMutation
}

// Where appends a list predicates to the LockOrderFulfillmentUpdate builder.
func (lofu *LockOrderFulfillmentUpdate) Where(ps ...predicate.LockOrderFulfillment) *LockOrderFulfillmentUpdate {
	lofu.mutation.Where(ps...)
	return lofu
}

// SetUpdatedAt sets the "updated_at" field.
func (lofu *LockOrderFulfillmentUpdate) SetUpdatedAt(t time.Time) *LockOrderFulfillmentUpdate {
	lofu.mutation.SetUpdatedAt(t)
	return lofu
}

// SetTxID sets the "tx_id" field.
func (lofu *LockOrderFulfillmentUpdate) SetTxID(s string) *LockOrderFulfillmentUpdate {
	lofu.mutation.SetTxID(s)
	return lofu
}

// SetNillableTxID sets the "tx_id" field if the given value is not nil.
func (lofu *LockOrderFulfillmentUpdate) SetNillableTxID(s *string) *LockOrderFulfillmentUpdate {
	if s != nil {
		lofu.SetTxID(*s)
	}
	return lofu
}

// ClearTxID clears the value of the "tx_id" field.
func (lofu *LockOrderFulfillmentUpdate) ClearTxID() *LockOrderFulfillmentUpdate {
	lofu.mutation.ClearTxID()
	return lofu
}

// SetPsp sets the "psp" field.
func (lofu *LockOrderFulfillmentUpdate) SetPsp(s string) *LockOrderFulfillmentUpdate {
	lofu.mutation.SetPsp(s)
	return lofu
}

// SetNillablePsp sets the "psp" field if the given value is not nil.
func (lofu *LockOrderFulfillmentUpdate) SetNillablePsp(s *string) *LockOrderFulfillmentUpdate {
	if s != nil {
		lofu.SetPsp(*s)
	}
	return lofu
}

// ClearPsp clears the value of the "psp" field.
func (lofu *LockOrderFulfillmentUpdate) ClearPsp() *LockOrderFulfillmentUpdate {
	lofu.mutation.ClearPsp()
	return lofu
}

// SetValidationStatus sets the "validation_status" field.
func (lofu *LockOrderFulfillmentUpdate) SetValidationStatus(ls lockorderfulfillment.ValidationStatus) *LockOrderFulfillmentUpdate {
	lofu.mutation.SetValidationStatus(ls)
	return lofu
}

// SetNillableValidationStatus sets the "validation_status" field if the given value is not nil.
func (lofu *LockOrderFulfillmentUpdate) SetNillableValidationStatus(ls *lockorderfulfillment.ValidationStatus) *LockOrderFulfillmentUpdate {
	if ls != nil {
		lofu.SetValidationStatus(*ls)
	}
	return lofu
}

// SetValidationError sets the "validation_error" field.
func (lofu *LockOrderFulfillmentUpdate) SetValidationError(s string) *LockOrderFulfillmentUpdate {
	lofu.mutation.SetValidationError(s)
	return lofu
}

// SetNillableValidationError sets the "validation_error" field if the given value is not nil.
func (lofu *LockOrderFulfillmentUpdate) SetNillableValidationError(s *string) *LockOrderFulfillmentUpdate {
	if s != nil {
		lofu.SetValidationError(*s)
	}
	return lofu
}

// ClearValidationError clears the value of the "validation_error" field.
func (lofu *LockOrderFulfillmentUpdate) ClearValidationError() *LockOrderFulfillmentUpdate {
	lofu.mutation.ClearValidationError()
	return lofu
}

// SetOrderID sets the "order" edge to the LockPaymentOrder entity by ID.
func (lofu *LockOrderFulfillmentUpdate) SetOrderID(id uuid.UUID) *LockOrderFulfillmentUpdate {
	lofu.mutation.SetOrderID(id)
	return lofu
}

// SetOrder sets the "order" edge to the LockPaymentOrder entity.
func (lofu *LockOrderFulfillmentUpdate) SetOrder(l *LockPaymentOrder) *LockOrderFulfillmentUpdate {
	return lofu.SetOrderID(l.ID)
}

// Mutation returns the LockOrderFulfillmentMutation object of the builder.
func (lofu *LockOrderFulfillmentUpdate) Mutation() *LockOrderFulfillmentMutation {
	return lofu.mutation
}

// ClearOrder clears the "order" edge to the LockPaymentOrder entity.
func (lofu *LockOrderFulfillmentUpdate) ClearOrder() *LockOrderFulfillmentUpdate {
	lofu.mutation.ClearOrder()
	return lofu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lofu *LockOrderFulfillmentUpdate) Save(ctx context.Context) (int, error) {
	lofu.defaults()
	return withHooks(ctx, lofu.sqlSave, lofu.mutation, lofu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lofu *LockOrderFulfillmentUpdate) SaveX(ctx context.Context) int {
	affected, err := lofu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lofu *LockOrderFulfillmentUpdate) Exec(ctx context.Context) error {
	_, err := lofu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lofu *LockOrderFulfillmentUpdate) ExecX(ctx context.Context) {
	if err := lofu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lofu *LockOrderFulfillmentUpdate) defaults() {
	if _, ok := lofu.mutation.UpdatedAt(); !ok {
		v := lockorderfulfillment.UpdateDefaultUpdatedAt()
		lofu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lofu *LockOrderFulfillmentUpdate) check() error {
	if v, ok := lofu.mutation.ValidationStatus(); ok {
		if err := lockorderfulfillment.ValidationStatusValidator(v); err != nil {
			return &ValidationError{Name: "validation_status", err: fmt.Errorf(`ent: validator failed for field "LockOrderFulfillment.validation_status": %w`, err)}
		}
	}
	if lofu.mutation.OrderCleared() && len(lofu.mutation.OrderIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "LockOrderFulfillment.order"`)
	}
	return nil
}

func (lofu *LockOrderFulfillmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lofu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(lockorderfulfillment.Table, lockorderfulfillment.Columns, sqlgraph.NewFieldSpec(lockorderfulfillment.FieldID, field.TypeUUID))
	if ps := lofu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lofu.mutation.UpdatedAt(); ok {
		_spec.SetField(lockorderfulfillment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lofu.mutation.TxID(); ok {
		_spec.SetField(lockorderfulfillment.FieldTxID, field.TypeString, value)
	}
	if lofu.mutation.TxIDCleared() {
		_spec.ClearField(lockorderfulfillment.FieldTxID, field.TypeString)
	}
	if value, ok := lofu.mutation.Psp(); ok {
		_spec.SetField(lockorderfulfillment.FieldPsp, field.TypeString, value)
	}
	if lofu.mutation.PspCleared() {
		_spec.ClearField(lockorderfulfillment.FieldPsp, field.TypeString)
	}
	if value, ok := lofu.mutation.ValidationStatus(); ok {
		_spec.SetField(lockorderfulfillment.FieldValidationStatus, field.TypeEnum, value)
	}
	if value, ok := lofu.mutation.ValidationError(); ok {
		_spec.SetField(lockorderfulfillment.FieldValidationError, field.TypeString, value)
	}
	if lofu.mutation.ValidationErrorCleared() {
		_spec.ClearField(lockorderfulfillment.FieldValidationError, field.TypeString)
	}
	if lofu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lockorderfulfillment.OrderTable,
			Columns: []string{lockorderfulfillment.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lofu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lockorderfulfillment.OrderTable,
			Columns: []string{lockorderfulfillment.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lofu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lockorderfulfillment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lofu.mutation.done = true
	return n, nil
}

// LockOrderFulfillmentUpdateOne is the builder for updating a single LockOrderFulfillment entity.
type LockOrderFulfillmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LockOrderFulfillmentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (lofuo *LockOrderFulfillmentUpdateOne) SetUpdatedAt(t time.Time) *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.SetUpdatedAt(t)
	return lofuo
}

// SetTxID sets the "tx_id" field.
func (lofuo *LockOrderFulfillmentUpdateOne) SetTxID(s string) *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.SetTxID(s)
	return lofuo
}

// SetNillableTxID sets the "tx_id" field if the given value is not nil.
func (lofuo *LockOrderFulfillmentUpdateOne) SetNillableTxID(s *string) *LockOrderFulfillmentUpdateOne {
	if s != nil {
		lofuo.SetTxID(*s)
	}
	return lofuo
}

// ClearTxID clears the value of the "tx_id" field.
func (lofuo *LockOrderFulfillmentUpdateOne) ClearTxID() *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.ClearTxID()
	return lofuo
}

// SetPsp sets the "psp" field.
func (lofuo *LockOrderFulfillmentUpdateOne) SetPsp(s string) *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.SetPsp(s)
	return lofuo
}

// SetNillablePsp sets the "psp" field if the given value is not nil.
func (lofuo *LockOrderFulfillmentUpdateOne) SetNillablePsp(s *string) *LockOrderFulfillmentUpdateOne {
	if s != nil {
		lofuo.SetPsp(*s)
	}
	return lofuo
}

// ClearPsp clears the value of the "psp" field.
func (lofuo *LockOrderFulfillmentUpdateOne) ClearPsp() *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.ClearPsp()
	return lofuo
}

// SetValidationStatus sets the "validation_status" field.
func (lofuo *LockOrderFulfillmentUpdateOne) SetValidationStatus(ls lockorderfulfillment.ValidationStatus) *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.SetValidationStatus(ls)
	return lofuo
}

// SetNillableValidationStatus sets the "validation_status" field if the given value is not nil.
func (lofuo *LockOrderFulfillmentUpdateOne) SetNillableValidationStatus(ls *lockorderfulfillment.ValidationStatus) *LockOrderFulfillmentUpdateOne {
	if ls != nil {
		lofuo.SetValidationStatus(*ls)
	}
	return lofuo
}

// SetValidationError sets the "validation_error" field.
func (lofuo *LockOrderFulfillmentUpdateOne) SetValidationError(s string) *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.SetValidationError(s)
	return lofuo
}

// SetNillableValidationError sets the "validation_error" field if the given value is not nil.
func (lofuo *LockOrderFulfillmentUpdateOne) SetNillableValidationError(s *string) *LockOrderFulfillmentUpdateOne {
	if s != nil {
		lofuo.SetValidationError(*s)
	}
	return lofuo
}

// ClearValidationError clears the value of the "validation_error" field.
func (lofuo *LockOrderFulfillmentUpdateOne) ClearValidationError() *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.ClearValidationError()
	return lofuo
}

// SetOrderID sets the "order" edge to the LockPaymentOrder entity by ID.
func (lofuo *LockOrderFulfillmentUpdateOne) SetOrderID(id uuid.UUID) *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.SetOrderID(id)
	return lofuo
}

// SetOrder sets the "order" edge to the LockPaymentOrder entity.
func (lofuo *LockOrderFulfillmentUpdateOne) SetOrder(l *LockPaymentOrder) *LockOrderFulfillmentUpdateOne {
	return lofuo.SetOrderID(l.ID)
}

// Mutation returns the LockOrderFulfillmentMutation object of the builder.
func (lofuo *LockOrderFulfillmentUpdateOne) Mutation() *LockOrderFulfillmentMutation {
	return lofuo.mutation
}

// ClearOrder clears the "order" edge to the LockPaymentOrder entity.
func (lofuo *LockOrderFulfillmentUpdateOne) ClearOrder() *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.ClearOrder()
	return lofuo
}

// Where appends a list predicates to the LockOrderFulfillmentUpdate builder.
func (lofuo *LockOrderFulfillmentUpdateOne) Where(ps ...predicate.LockOrderFulfillment) *LockOrderFulfillmentUpdateOne {
	lofuo.mutation.Where(ps...)
	return lofuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lofuo *LockOrderFulfillmentUpdateOne) Select(field string, fields ...string) *LockOrderFulfillmentUpdateOne {
	lofuo.fields = append([]string{field}, fields...)
	return lofuo
}

// Save executes the query and returns the updated LockOrderFulfillment entity.
func (lofuo *LockOrderFulfillmentUpdateOne) Save(ctx context.Context) (*LockOrderFulfillment, error) {
	lofuo.defaults()
	return withHooks(ctx, lofuo.sqlSave, lofuo.mutation, lofuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lofuo *LockOrderFulfillmentUpdateOne) SaveX(ctx context.Context) *LockOrderFulfillment {
	node, err := lofuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lofuo *LockOrderFulfillmentUpdateOne) Exec(ctx context.Context) error {
	_, err := lofuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lofuo *LockOrderFulfillmentUpdateOne) ExecX(ctx context.Context) {
	if err := lofuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lofuo *LockOrderFulfillmentUpdateOne) defaults() {
	if _, ok := lofuo.mutation.UpdatedAt(); !ok {
		v := lockorderfulfillment.UpdateDefaultUpdatedAt()
		lofuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lofuo *LockOrderFulfillmentUpdateOne) check() error {
	if v, ok := lofuo.mutation.ValidationStatus(); ok {
		if err := lockorderfulfillment.ValidationStatusValidator(v); err != nil {
			return &ValidationError{Name: "validation_status", err: fmt.Errorf(`ent: validator failed for field "LockOrderFulfillment.validation_status": %w`, err)}
		}
	}
	if lofuo.mutation.OrderCleared() && len(lofuo.mutation.OrderIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "LockOrderFulfillment.order"`)
	}
	return nil
}

func (lofuo *LockOrderFulfillmentUpdateOne) sqlSave(ctx context.Context) (_node *LockOrderFulfillment, err error) {
	if err := lofuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(lockorderfulfillment.Table, lockorderfulfillment.Columns, sqlgraph.NewFieldSpec(lockorderfulfillment.FieldID, field.TypeUUID))
	id, ok := lofuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LockOrderFulfillment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lofuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lockorderfulfillment.FieldID)
		for _, f := range fields {
			if !lockorderfulfillment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lockorderfulfillment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lofuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lofuo.mutation.UpdatedAt(); ok {
		_spec.SetField(lockorderfulfillment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lofuo.mutation.TxID(); ok {
		_spec.SetField(lockorderfulfillment.FieldTxID, field.TypeString, value)
	}
	if lofuo.mutation.TxIDCleared() {
		_spec.ClearField(lockorderfulfillment.FieldTxID, field.TypeString)
	}
	if value, ok := lofuo.mutation.Psp(); ok {
		_spec.SetField(lockorderfulfillment.FieldPsp, field.TypeString, value)
	}
	if lofuo.mutation.PspCleared() {
		_spec.ClearField(lockorderfulfillment.FieldPsp, field.TypeString)
	}
	if value, ok := lofuo.mutation.ValidationStatus(); ok {
		_spec.SetField(lockorderfulfillment.FieldValidationStatus, field.TypeEnum, value)
	}
	if value, ok := lofuo.mutation.ValidationError(); ok {
		_spec.SetField(lockorderfulfillment.FieldValidationError, field.TypeString, value)
	}
	if lofuo.mutation.ValidationErrorCleared() {
		_spec.ClearField(lockorderfulfillment.FieldValidationError, field.TypeString)
	}
	if lofuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lockorderfulfillment.OrderTable,
			Columns: []string{lockorderfulfillment.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lofuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lockorderfulfillment.OrderTable,
			Columns: []string{lockorderfulfillment.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(lockpaymentorder.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LockOrderFulfillment{config: lofuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lofuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lockorderfulfillment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lofuo.mutation.done = true
	return _node, nil
}
