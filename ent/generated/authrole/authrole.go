// Code generated by ent, DO NOT EDIT.

package authrole

import (
	"fmt"
	"time"

	"github.com/chenningg/ulid-test/auth"
	"github.com/oklog/ulid/v2"
	ulid "github.com/oklog/ulid/v2"
)

const (
	// Label holds the string label denoting the authrole type in the database.
	Label = "auth_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeAccounts holds the string denoting the accounts edge name in mutations.
	EdgeAccounts = "accounts"
	// Table holds the table name of the authrole in the database.
	Table = "auth_roles"
	// AccountsTable is the table that holds the accounts relation/edge. The primary key declared below.
	AccountsTable = "account_auth_roles"
	// AccountsInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountsInverseTable = "accounts"
)

// Columns holds all SQL columns for authrole fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldRole,
	FieldDescription,
}

var (
	// AccountsPrimaryKey and AccountsColumn2 are the table columns denoting the
	// primary key for the accounts relation (M2M).
	AccountsPrimaryKey = []string{"account_id", "auth_role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ULID
)

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r auth.AuthRole) error {
	switch r.String() {
	case "Demo", "Free", "Plus", "Pro", "Enterprise", "Support", "Admin", "SuperAdmin":
		return nil
	default:
		return fmt.Errorf("authrole: invalid enum value for role field: %q", r)
	}
}
