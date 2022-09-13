package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/chenningg/ulid-test/auth"
	"github.com/chenningg/ulid-test/ent/schema/mixin"
)

// AuthRole holds the schema definition for the AuthRole entity.
type AuthRole struct {
	ent.Schema
}

// Adds id, created_at and updated_at fields.
func (AuthRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the AuthRole.
func (AuthRole) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").
			GoType(auth.AuthRole(1)),
		field.String("description").
			Optional().
			Nillable().
			NotEmpty(),
	}
}

// Edges of the AuthRole.
func (AuthRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("accounts", Account.Type).
			Ref("auth_roles"),
	}
}
