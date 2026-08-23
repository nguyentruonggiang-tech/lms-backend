package schema

import (
	"time"

	softdelete "lms-api/ent/soft-delete"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Users struct{ ent.Schema }

func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("password").Sensitive(),
		field.String("full_name").Optional().Nillable(),
		field.String("avatar").Optional().Nillable(),
		field.Enum("role").Values("student", "admin").Default("student"),
		field.Enum("status").Values("active", "blocked").Default("active"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Enrollments", Enrollments.Type),
		edge.To("LessonProgresses", LessonProgresses.Type),
		edge.To("QuizAttempts", QuizAttempts.Type),
		edge.To("Certificates", Certificates.Type),
		edge.To("Notifications", Notifications.Type),
	}
}

func (Users) Mixin() []ent.Mixin {
	return []ent.Mixin{
		softdelete.SoftDeleteMixin{},
	}
}
