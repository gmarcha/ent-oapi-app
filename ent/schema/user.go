package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("token").
			Unique().
			Sensitive(),
		field.String("name").
			Unique().
			StructTag(`binding:"required,lowercase"`),
		field.String("first_name").
			StructTag(`binding:"required,lowercase"`),
		field.String("last_name").
			StructTag(`binding:"required,lowercase"`),
		field.Bool("tutor").
			Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).
			Ref("users").
			Annotations(
				entoas.CreateOperation(
					entoas.OperationPolicy(entoas.PolicyExclude),
				),
			),
	}
}
