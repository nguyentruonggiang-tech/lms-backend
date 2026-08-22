package schema

import (
	"time"

	softdelete "lms-backend/ent/soft-delete"

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
	}
}

func (Categories) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Courses", Courses.Type),
	}
}

func (Categories) Mixin() []ent.Mixin {
	return []ent.Mixin{
		softdelete.SoftDeleteMixin{},
	}
}
