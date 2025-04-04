// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/paycrest/aggregator/ent/fiatcurrency"
	"github.com/paycrest/aggregator/ent/provisionbucket"
	"github.com/shopspring/decimal"
)

// ProvisionBucket is the model entity for the ProvisionBucket schema.
type ProvisionBucket struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MinAmount holds the value of the "min_amount" field.
	MinAmount decimal.Decimal `json:"min_amount,omitempty"`
	// MaxAmount holds the value of the "max_amount" field.
	MaxAmount decimal.Decimal `json:"max_amount,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProvisionBucketQuery when eager-loading is set.
	Edges                           ProvisionBucketEdges `json:"edges"`
	fiat_currency_provision_buckets *uuid.UUID
	selectValues                    sql.SelectValues
}

// ProvisionBucketEdges holds the relations/edges for other nodes in the graph.
type ProvisionBucketEdges struct {
	// Currency holds the value of the currency edge.
	Currency *FiatCurrency `json:"currency,omitempty"`
	// LockPaymentOrders holds the value of the lock_payment_orders edge.
	LockPaymentOrders []*LockPaymentOrder `json:"lock_payment_orders,omitempty"`
	// ProviderProfiles holds the value of the provider_profiles edge.
	ProviderProfiles []*ProviderProfile `json:"provider_profiles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CurrencyOrErr returns the Currency value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProvisionBucketEdges) CurrencyOrErr() (*FiatCurrency, error) {
	if e.Currency != nil {
		return e.Currency, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: fiatcurrency.Label}
	}
	return nil, &NotLoadedError{edge: "currency"}
}

// LockPaymentOrdersOrErr returns the LockPaymentOrders value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionBucketEdges) LockPaymentOrdersOrErr() ([]*LockPaymentOrder, error) {
	if e.loadedTypes[1] {
		return e.LockPaymentOrders, nil
	}
	return nil, &NotLoadedError{edge: "lock_payment_orders"}
}

// ProviderProfilesOrErr returns the ProviderProfiles value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionBucketEdges) ProviderProfilesOrErr() ([]*ProviderProfile, error) {
	if e.loadedTypes[2] {
		return e.ProviderProfiles, nil
	}
	return nil, &NotLoadedError{edge: "provider_profiles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProvisionBucket) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case provisionbucket.FieldMinAmount, provisionbucket.FieldMaxAmount:
			values[i] = new(decimal.Decimal)
		case provisionbucket.FieldID:
			values[i] = new(sql.NullInt64)
		case provisionbucket.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case provisionbucket.ForeignKeys[0]: // fiat_currency_provision_buckets
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProvisionBucket fields.
func (pb *ProvisionBucket) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provisionbucket.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pb.ID = int(value.Int64)
		case provisionbucket.FieldMinAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field min_amount", values[i])
			} else if value != nil {
				pb.MinAmount = *value
			}
		case provisionbucket.FieldMaxAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field max_amount", values[i])
			} else if value != nil {
				pb.MaxAmount = *value
			}
		case provisionbucket.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pb.CreatedAt = value.Time
			}
		case provisionbucket.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field fiat_currency_provision_buckets", values[i])
			} else if value.Valid {
				pb.fiat_currency_provision_buckets = new(uuid.UUID)
				*pb.fiat_currency_provision_buckets = *value.S.(*uuid.UUID)
			}
		default:
			pb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProvisionBucket.
// This includes values selected through modifiers, order, etc.
func (pb *ProvisionBucket) Value(name string) (ent.Value, error) {
	return pb.selectValues.Get(name)
}

// QueryCurrency queries the "currency" edge of the ProvisionBucket entity.
func (pb *ProvisionBucket) QueryCurrency() *FiatCurrencyQuery {
	return NewProvisionBucketClient(pb.config).QueryCurrency(pb)
}

// QueryLockPaymentOrders queries the "lock_payment_orders" edge of the ProvisionBucket entity.
func (pb *ProvisionBucket) QueryLockPaymentOrders() *LockPaymentOrderQuery {
	return NewProvisionBucketClient(pb.config).QueryLockPaymentOrders(pb)
}

// QueryProviderProfiles queries the "provider_profiles" edge of the ProvisionBucket entity.
func (pb *ProvisionBucket) QueryProviderProfiles() *ProviderProfileQuery {
	return NewProvisionBucketClient(pb.config).QueryProviderProfiles(pb)
}

// Update returns a builder for updating this ProvisionBucket.
// Note that you need to call ProvisionBucket.Unwrap() before calling this method if this ProvisionBucket
// was returned from a transaction, and the transaction was committed or rolled back.
func (pb *ProvisionBucket) Update() *ProvisionBucketUpdateOne {
	return NewProvisionBucketClient(pb.config).UpdateOne(pb)
}

// Unwrap unwraps the ProvisionBucket entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pb *ProvisionBucket) Unwrap() *ProvisionBucket {
	_tx, ok := pb.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProvisionBucket is not a transactional entity")
	}
	pb.config.driver = _tx.drv
	return pb
}

// String implements the fmt.Stringer.
func (pb *ProvisionBucket) String() string {
	var builder strings.Builder
	builder.WriteString("ProvisionBucket(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pb.ID))
	builder.WriteString("min_amount=")
	builder.WriteString(fmt.Sprintf("%v", pb.MinAmount))
	builder.WriteString(", ")
	builder.WriteString("max_amount=")
	builder.WriteString(fmt.Sprintf("%v", pb.MaxAmount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pb.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ProvisionBuckets is a parsable slice of ProvisionBucket.
type ProvisionBuckets []*ProvisionBucket
