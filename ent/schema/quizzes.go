package schema

import (
	"time"

	softdelete "lms-backend/ent/soft-delete"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Quizzes struct{ ent.Schema }

func (Quizzes) Fields() []ent.Field {
	return []ent.Field{
		field.Int("course_id"),
		field.Int("lesson_id"),
		field.String("title"),
		field.Int("passing_score").Default(70),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Quizzes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Courses", Courses.Type).Ref("Quizzes").Field("course_id").Unique().Required(),
		edge.From("Lessons", Lessons.Type).Ref("Quizzes").Field("lesson_id").Unique().Required(),
		edge.To("Questions", Questions.Type),
		edge.To("QuizAttempts", QuizAttempts.Type),
	}
}

func (Quizzes) Mixin() []ent.Mixin {
	return []ent.Mixin{
		softdelete.SoftDeleteMixin{},
	}
}
