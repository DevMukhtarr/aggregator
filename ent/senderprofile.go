// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/paycrest/aggregator/ent/apikey"
	"github.com/paycrest/aggregator/ent/senderprofile"
	"github.com/paycrest/aggregator/ent/user"
)

// SenderProfile is the model entity for the SenderProfile schema.
type SenderProfile struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// WebhookURL holds the value of the "webhook_url" field.
	WebhookURL string `json:"webhook_url,omitempty"`
	// DomainWhitelist holds the value of the "domain_whitelist" field.
	DomainWhitelist []string `json:"domain_whitelist,omitempty"`
	// ProviderID holds the value of the "provider_id" field.
	ProviderID string `json:"provider_id,omitempty"`
	// IsPartner holds the value of the "is_partner" field.
	IsPartner bool `json:"is_partner,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SenderProfileQuery when eager-loading is set.
	Edges               SenderProfileEdges `json:"edges"`
	user_sender_profile *uuid.UUID
	selectValues        sql.SelectValues
}

// SenderProfileEdges holds the relations/edges for other nodes in the graph.
type SenderProfileEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// APIKey holds the value of the api_key edge.
	APIKey *APIKey `json:"api_key,omitempty"`
	// PaymentOrders holds the value of the payment_orders edge.
	PaymentOrders []*PaymentOrder `json:"payment_orders,omitempty"`
	// OrderTokens holds the value of the order_tokens edge.
	OrderTokens []*SenderOrderToken `json:"order_tokens,omitempty"`
	// LinkedAddress holds the value of the linked_address edge.
	LinkedAddress []*LinkedAddress `json:"linked_address,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SenderProfileEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// APIKeyOrErr returns the APIKey value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SenderProfileEdges) APIKeyOrErr() (*APIKey, error) {
	if e.APIKey != nil {
		return e.APIKey, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: apikey.Label}
	}
	return nil, &NotLoadedError{edge: "api_key"}
}

// PaymentOrdersOrErr returns the PaymentOrders value or an error if the edge
// was not loaded in eager-loading.
func (e SenderProfileEdges) PaymentOrdersOrErr() ([]*PaymentOrder, error) {
	if e.loadedTypes[2] {
		return e.PaymentOrders, nil
	}
	return nil, &NotLoadedError{edge: "payment_orders"}
}

// OrderTokensOrErr returns the OrderTokens value or an error if the edge
// was not loaded in eager-loading.
func (e SenderProfileEdges) OrderTokensOrErr() ([]*SenderOrderToken, error) {
	if e.loadedTypes[3] {
		return e.OrderTokens, nil
	}
	return nil, &NotLoadedError{edge: "order_tokens"}
}

// LinkedAddressOrErr returns the LinkedAddress value or an error if the edge
// was not loaded in eager-loading.
func (e SenderProfileEdges) LinkedAddressOrErr() ([]*LinkedAddress, error) {
	if e.loadedTypes[4] {
		return e.LinkedAddress, nil
	}
	return nil, &NotLoadedError{edge: "linked_address"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SenderProfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case senderprofile.FieldDomainWhitelist:
			values[i] = new([]byte)
		case senderprofile.FieldIsPartner, senderprofile.FieldIsActive:
			values[i] = new(sql.NullBool)
		case senderprofile.FieldWebhookURL, senderprofile.FieldProviderID:
			values[i] = new(sql.NullString)
		case senderprofile.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case senderprofile.FieldID:
			values[i] = new(uuid.UUID)
		case senderprofile.ForeignKeys[0]: // user_sender_profile
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SenderProfile fields.
func (sp *SenderProfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case senderprofile.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sp.ID = *value
			}
		case senderprofile.FieldWebhookURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field webhook_url", values[i])
			} else if value.Valid {
				sp.WebhookURL = value.String
			}
		case senderprofile.FieldDomainWhitelist:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field domain_whitelist", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &sp.DomainWhitelist); err != nil {
					return fmt.Errorf("unmarshal field domain_whitelist: %w", err)
				}
			}
		case senderprofile.FieldProviderID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_id", values[i])
			} else if value.Valid {
				sp.ProviderID = value.String
			}
		case senderprofile.FieldIsPartner:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_partner", values[i])
			} else if value.Valid {
				sp.IsPartner = value.Bool
			}
		case senderprofile.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				sp.IsActive = value.Bool
			}
		case senderprofile.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sp.UpdatedAt = value.Time
			}
		case senderprofile.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_sender_profile", values[i])
			} else if value.Valid {
				sp.user_sender_profile = new(uuid.UUID)
				*sp.user_sender_profile = *value.S.(*uuid.UUID)
			}
		default:
			sp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SenderProfile.
// This includes values selected through modifiers, order, etc.
func (sp *SenderProfile) Value(name string) (ent.Value, error) {
	return sp.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the SenderProfile entity.
func (sp *SenderProfile) QueryUser() *UserQuery {
	return NewSenderProfileClient(sp.config).QueryUser(sp)
}

// QueryAPIKey queries the "api_key" edge of the SenderProfile entity.
func (sp *SenderProfile) QueryAPIKey() *APIKeyQuery {
	return NewSenderProfileClient(sp.config).QueryAPIKey(sp)
}

// QueryPaymentOrders queries the "payment_orders" edge of the SenderProfile entity.
func (sp *SenderProfile) QueryPaymentOrders() *PaymentOrderQuery {
	return NewSenderProfileClient(sp.config).QueryPaymentOrders(sp)
}

// QueryOrderTokens queries the "order_tokens" edge of the SenderProfile entity.
func (sp *SenderProfile) QueryOrderTokens() *SenderOrderTokenQuery {
	return NewSenderProfileClient(sp.config).QueryOrderTokens(sp)
}

// QueryLinkedAddress queries the "linked_address" edge of the SenderProfile entity.
func (sp *SenderProfile) QueryLinkedAddress() *LinkedAddressQuery {
	return NewSenderProfileClient(sp.config).QueryLinkedAddress(sp)
}

// Update returns a builder for updating this SenderProfile.
// Note that you need to call SenderProfile.Unwrap() before calling this method if this SenderProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *SenderProfile) Update() *SenderProfileUpdateOne {
	return NewSenderProfileClient(sp.config).UpdateOne(sp)
}

// Unwrap unwraps the SenderProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *SenderProfile) Unwrap() *SenderProfile {
	_tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: SenderProfile is not a transactional entity")
	}
	sp.config.driver = _tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *SenderProfile) String() string {
	var builder strings.Builder
	builder.WriteString("SenderProfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sp.ID))
	builder.WriteString("webhook_url=")
	builder.WriteString(sp.WebhookURL)
	builder.WriteString(", ")
	builder.WriteString("domain_whitelist=")
	builder.WriteString(fmt.Sprintf("%v", sp.DomainWhitelist))
	builder.WriteString(", ")
	builder.WriteString("provider_id=")
	builder.WriteString(sp.ProviderID)
	builder.WriteString(", ")
	builder.WriteString("is_partner=")
	builder.WriteString(fmt.Sprintf("%v", sp.IsPartner))
	builder.WriteString(", ")
	builder.WriteString("is_active=")
	builder.WriteString(fmt.Sprintf("%v", sp.IsActive))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sp.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SenderProfiles is a parsable slice of SenderProfile.
type SenderProfiles []*SenderProfile
