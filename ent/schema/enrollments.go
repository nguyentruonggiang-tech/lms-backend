package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Enrollments struct{ ent.Schema }

func (Enrollments) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("course_id"),
		field.Enum("status").Values("active", "completed", "cancelled").Default("active"),
		field.Float("progress_percent").Default(0),
		field.Time("enrolled_at").Default(time.Now),
		field.Time("completed_at").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Enrollments) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Users", Users.Type).Ref("Enrollments").Field("user_id").Unique().Required(),
		edge.From("Courses", Courses.Type).Ref("Enrollments").Field("course_id").Unique().Required(),
	}
}
