package schema

import (
	"time"

	softdelete "lms-backend/ent/soft-delete"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Sections struct{ ent.Schema }

func (Sections) Fields() []ent.Field {
	return []ent.Field{
		field.Int("course_id"),
		field.String("title"),
		field.Int("sort_order").Default(0),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Sections) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Courses", Courses.Type).Ref("Sections").Field("course_id").Unique().Required(),
		edge.To("Lessons", Lessons.Type),
	}
}

func (Sections) Mixin() []ent.Mixin {
	return []ent.Mixin{
		softdelete.SoftDeleteMixin{},
	}
}
