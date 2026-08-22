package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Questions struct{ ent.Schema }

func (Questions) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quiz_id"),
		field.Text("question_text"),
		field.String("option_a"),
		field.String("option_b"),
		field.String("option_c"),
		field.String("option_d"),
		field.Enum("correct_option").Values("A", "B", "C", "D"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Questions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Quizzes", Quizzes.Type).Ref("Questions").Field("quiz_id").Unique().Required(),
	}
}
