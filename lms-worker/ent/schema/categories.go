package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Categories struct{ ent.Schema }

func (Categories) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (Categories) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Courses", Courses.Type),
	}
}
