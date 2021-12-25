package schema

import (
	"time"

	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable(),
		field.String("name").
			StructTag(`binding:"required,lowercase"`),
		field.String("description").
			Default(""),
		field.Time("date_creation").
			Default(time.Now).
			Immutable(),
		field.Time("date_update").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("date_start").
			StructTag(`binding:"required,validateDateStart"`),
		field.Time("date_end").
			StructTag(`binding:"required,gtfield=DateStart"`),
		field.Int("nb_tutors").
			StructTag(`binding:"required,gt=0"`),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Annotations(
				entoas.CreateOperation(
					entoas.OperationPolicy(entoas.PolicyExclude),
				),
			),
	}
}

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "date_start", "date_end").
			Unique(),
	}
}
