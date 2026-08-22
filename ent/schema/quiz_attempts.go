package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type QuizAttempts struct{ ent.Schema }

func (QuizAttempts) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("quiz_id"),
		field.Float("score").Default(0),
		field.Int("total_questions").Default(0),
		field.Int("correct_answers").Default(0),
		field.Bool("is_passed").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (QuizAttempts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Users", Users.Type).Ref("QuizAttempts").Field("user_id").Unique().Required(),
		edge.From("Quizzes", Quizzes.Type).Ref("QuizAttempts").Field("quiz_id").Unique().Required(),
	}
}
