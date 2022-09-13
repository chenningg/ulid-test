package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid/v2"
)

type IdMixin struct {
	mixin.Schema
}

func (IdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", ulid.ULID{}).
			Default(ulid.Make),
	}
}
