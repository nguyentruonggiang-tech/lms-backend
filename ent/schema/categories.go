package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Categories struct{ ent.Schema }

func (Categories) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("slug").Unique(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Categories) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Courses", Courses.Type),
	}
}
