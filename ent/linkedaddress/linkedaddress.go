// Code generated by ent, DO NOT EDIT.

package linkedaddress

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the linkedaddress type in the database.
	Label = "linked_address"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldInstitution holds the string denoting the institution field in the database.
	FieldInstitution = "institution"
	// FieldAccountIdentifier holds the string denoting the account_identifier field in the database.
	FieldAccountIdentifier = "account_identifier"
	// FieldAccountName holds the string denoting the account_name field in the database.
	FieldAccountName = "account_name"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldOwnerAddress holds the string denoting the owner_address field in the database.
	FieldOwnerAddress = "owner_address"
	// FieldLastIndexedBlock holds the string denoting the last_indexed_block field in the database.
	FieldLastIndexedBlock = "last_indexed_block"
	// FieldTxHash holds the string denoting the tx_hash field in the database.
	FieldTxHash = "tx_hash"
	// EdgePaymentOrders holds the string denoting the payment_orders edge name in mutations.
	EdgePaymentOrders = "payment_orders"
	// Table holds the table name of the linkedaddress in the database.
	Table = "linked_addresses"
	// PaymentOrdersTable is the table that holds the payment_orders relation/edge.
	PaymentOrdersTable = "payment_orders"
	// PaymentOrdersInverseTable is the table name for the PaymentOrder entity.
	// It exists in this package in order to avoid circular dependency with the "paymentorder" package.
	PaymentOrdersInverseTable = "payment_orders"
	// PaymentOrdersColumn is the table column denoting the payment_orders relation/edge.
	PaymentOrdersColumn = "linked_address_payment_orders"
)

// Columns holds all SQL columns for linkedaddress fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAddress,
	FieldSalt,
	FieldInstitution,
	FieldAccountIdentifier,
	FieldAccountName,
	FieldMetadata,
	FieldOwnerAddress,
	FieldLastIndexedBlock,
	FieldTxHash,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "linked_addresses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"sender_profile_linked_address",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// TxHashValidator is a validator for the "tx_hash" field. It is called by the builders before save.
	TxHashValidator func(string) error
)

// OrderOption defines the ordering options for the LinkedAddress queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByInstitution orders the results by the institution field.
func ByInstitution(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstitution, opts...).ToFunc()
}

// ByAccountIdentifier orders the results by the account_identifier field.
func ByAccountIdentifier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountIdentifier, opts...).ToFunc()
}

// ByAccountName orders the results by the account_name field.
func ByAccountName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountName, opts...).ToFunc()
}

// ByOwnerAddress orders the results by the owner_address field.
func ByOwnerAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerAddress, opts...).ToFunc()
}

// ByLastIndexedBlock orders the results by the last_indexed_block field.
func ByLastIndexedBlock(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastIndexedBlock, opts...).ToFunc()
}

// ByTxHash orders the results by the tx_hash field.
func ByTxHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxHash, opts...).ToFunc()
}

// ByPaymentOrdersCount orders the results by payment_orders count.
func ByPaymentOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPaymentOrdersStep(), opts...)
	}
}

// ByPaymentOrders orders the results by payment_orders terms.
func ByPaymentOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPaymentOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPaymentOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PaymentOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PaymentOrdersTable, PaymentOrdersColumn),
	)
}
