package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LessonProgresses struct{ ent.Schema }

func (LessonProgresses) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("course_id"),
		field.Int("lesson_id"),
		field.Bool("is_completed").Default(false),
		field.Time("completed_at").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (LessonProgresses) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Users", Users.Type).Ref("LessonProgresses").Field("user_id").Unique().Required(),
		edge.From("Courses", Courses.Type).Ref("LessonProgresses").Field("course_id").Unique().Required(),
		edge.From("Lessons", Lessons.Type).Ref("LessonProgresses").Field("lesson_id").Unique().Required(),
	}
}
